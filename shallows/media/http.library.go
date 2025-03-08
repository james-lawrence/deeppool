package media

import (
	"crypto/md5"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
	"github.com/james-lawrence/deeppool/internal/env"
	"github.com/james-lawrence/deeppool/internal/x/errorsx"
	"github.com/james-lawrence/deeppool/internal/x/formx"
	"github.com/james-lawrence/deeppool/internal/x/httpx"
	"github.com/james-lawrence/deeppool/internal/x/iox"
	"github.com/james-lawrence/deeppool/internal/x/jwtx"
	"github.com/james-lawrence/deeppool/internal/x/langx"
	"github.com/james-lawrence/deeppool/internal/x/md5x"
	"github.com/james-lawrence/deeppool/internal/x/numericx"
	"github.com/james-lawrence/deeppool/internal/x/sqlx"
	"github.com/james-lawrence/deeppool/internal/x/sqlxx"
	"github.com/james-lawrence/deeppool/library"
	"github.com/james-lawrence/deeppool/tracking"
	"github.com/james-lawrence/torrent/metainfo"
	"github.com/james-lawrence/torrent/storage"
	"github.com/justinas/alice"
)

type HTTPLibraryOption func(*HTTPLibrary)

func HTTPLibraryOptionJWTSecret(j jwtx.JWTSecretSource) HTTPLibraryOption {
	return func(t *HTTPLibrary) {
		t.jwtsecret = j
	}
}

func NewHTTPLibrary(q sqlx.Queryer, c storage.ClientImpl, options ...HTTPLibraryOption) *HTTPLibrary {
	svc := langx.Clone(HTTPLibrary{
		q:         q,
		c:         c,
		jwtsecret: env.JWTSecret,
		decoder:   formx.NewDecoder(),
	}, options...)

	return &svc
}

type HTTPLibrary struct {
	q         sqlx.Queryer
	c         storage.ClientImpl
	jwtsecret jwtx.JWTSecretSource
	decoder   *form.Decoder
}

func (t *HTTPLibrary) Bind(r *mux.Router) {
	r.StrictSlash(false)

	r.Path("/").Methods(http.MethodGet).Handler(alice.New(
		httpx.ContextBufferPool512(),
		httpx.ParseForm,
		// httpauth.AuthenticateWithToken(t.jwtsecret),
		// AuthzTokenHTTP(t.jwtsecret, AuthzPermUsermanagement),
		httpx.Timeout2s(),
	).ThenFunc(t.search))

	r.Path("/").Methods(http.MethodPost).Handler(alice.New(
		httpx.ContextBufferPool512(),
		httpx.ParseForm,
		// httpauth.AuthenticateWithToken(t.jwtsecret),
		// AuthzTokenHTTP(t.jwtsecret, AuthzPermUsermanagement),
		httpx.TimeoutRollingRead(3*time.Second),
	).ThenFunc(t.upload))

	r.Path("/{id}").Methods(http.MethodDelete).Handler(alice.New(
		httpx.ContextBufferPool512(),
		httpx.ParseForm,
		// httpauth.AuthenticateWithToken(t.jwtsecret),
		// AuthzTokenHTTP(t.jwtsecret, AuthzPermUsermanagement),
		httpx.Timeout2s(),
	).ThenFunc(t.delete))
}

func (t *HTTPLibrary) delete(w http.ResponseWriter, r *http.Request) {
	var (
		md library.Metadata
		id = mux.Vars(r)["id"]
	)

	if err := library.MetadataTombstoneByID(r.Context(), t.q, id).Scan(&md); sqlx.ErrNoRows(err) != nil {
		log.Println(errorsx.Wrap(err, "unable to tombstone metadata"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusNotFound))
		return
	} else if err != nil {
		log.Println(errorsx.Wrap(err, "unable to tombstone metadata"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusInternalServerError))
		return
	}

	if err := httpx.WriteJSON(w, httpx.GetBuffer(r), &MediaDeleteResponse{
		Media: langx.Autoptr(
			langx.Clone(
				Media{},
				MediaOptionFromLibraryMetadata(langx.Clone(md, library.MetadataOptionJSONSafeEncode))),
		),
	}); err != nil {
		log.Println(errorsx.Wrap(err, "unable to write response"))
		return
	}
}

func (t *HTTPLibrary) upload(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		f      multipart.File
		fh     *multipart.FileHeader
		copied iox.Copied
		mhash  = md5.New()
	)

	if f, fh, err = r.FormFile("content"); err != nil {
		log.Println(errorsx.Wrap(err, "content parameter required"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusBadRequest))
		return
	}
	defer f.Close()

	tmp, err := os.CreateTemp("", "retrovibed.upload.*")
	if err != nil {
		log.Println(errorsx.Wrap(err, "unable to create temporary file"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusInternalServerError))
		return
	}

	mi, err := metainfo.NewFromReader(io.TeeReader(r.Body, io.MultiWriter(tmp, mhash, &copied)), metainfo.OptionDisplayName(fh.Filename))
	if err != nil {
		log.Println(errorsx.Wrap(err, "unable to read upload"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusInternalServerError))
		return
	}
	_ = mi

	lmd := library.Metadata{
		ID:          md5x.FormatString(mhash),
		Description: fh.Filename,
		Bytes:       uint64(copied),
		Mimetype:    fh.Header.Get("Content-Type"),
	}

	if err = library.MetadataInsertWithDefaults(r.Context(), sqlx.Debug(t.q), lmd).Scan(&lmd); err != nil {
		log.Println(errorsx.Wrap(err, "unable to record library metadata record"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusInternalServerError))
		return
	}

	if err := httpx.WriteJSON(w, httpx.GetBuffer(r), &MediaUploadResponse{
		Media: langx.Autoptr(
			langx.Clone(
				Media{},
				MediaOptionFromLibraryMetadata(lmd),
			),
		),
	}); err != nil {
		log.Println(errorsx.Wrap(err, "unable to write response"))
		return
	}
}

func (t *HTTPLibrary) downloading(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		msg DownloadSearchResponse = DownloadSearchResponse{
			Next: &DownloadSearchRequest{
				Limit: 100,
			},
		}
	)

	if err = t.decoder.Decode(msg.Next, r.Form); err != nil {
		log.Println(errorsx.Wrap(err, "unable to decode request"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusBadRequest))
		return
	}
	msg.Next.Limit = numericx.Min(msg.Next.Limit, 100)

	q := tracking.MetadataSearchBuilder().Where(
		squirrel.And{
			tracking.MetadataQueryInitiated(),
			tracking.MetadataQueryIncomplete(),
			tracking.MetadataQueryNotPaused(),
		},
	).OrderBy("created_at DESC").Offset(msg.Next.Offset * msg.Next.Limit).Limit(msg.Next.Limit)

	err = sqlxx.ScanEach(tracking.MetadataSearch(r.Context(), t.q, q), func(p *tracking.Metadata) error {
		tmp := langx.Clone(Download{}, DownloadOptionFromTorrentMetadata(langx.Clone(*p, tracking.MetadataOptionJSONSafeEncode)))
		msg.Items = append(msg.Items, &tmp)
		return nil
	})

	if err != nil {
		log.Println(errorsx.Wrap(err, "encoding failed"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusInternalServerError))
		return
	}

	if err = httpx.WriteJSON(w, httpx.GetBuffer(r), &msg); err != nil {
		log.Println(errorsx.Wrap(err, "unable to write response"))
		return
	}
}

func (t *HTTPLibrary) search(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		msg MediaSearchResponse = MediaSearchResponse{
			Next: &MediaSearchRequest{
				Limit: 100,
			},
		}
	)

	if err = t.decoder.Decode(msg.Next, r.Form); err != nil {
		log.Println(errorsx.Wrap(err, "unable to decode request"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusBadRequest))
		return
	}
	msg.Next.Limit = numericx.Min(msg.Next.Limit, 100)

	q := library.MetadataSearchBuilder().Where(squirrel.And{
		library.MetadataQueryVisible(),
		library.MetadataQuerySearch(msg.Next.Query, "description"),
	}).OrderBy("created_at DESC").Offset(msg.Next.Offset * msg.Next.Limit).Limit(msg.Next.Limit)

	err = sqlxx.ScanEach(library.MetadataSearch(r.Context(), t.q, q), func(p *library.Metadata) error {
		tmp := langx.Clone(Media{}, MediaOptionFromLibraryMetadata(langx.Clone(*p, library.MetadataOptionJSONSafeEncode)))
		msg.Items = append(msg.Items, &tmp)
		return nil
	})

	if err != nil {
		log.Println(errorsx.Wrap(err, "encoding failed"))
		errorsx.MaybeLog(httpx.WriteEmptyJSON(w, http.StatusInternalServerError))
		return
	}

	if err = httpx.WriteJSON(w, httpx.GetBuffer(r), &msg); err != nil {
		log.Println(errorsx.Wrap(err, "unable to write response"))
		return
	}
}
