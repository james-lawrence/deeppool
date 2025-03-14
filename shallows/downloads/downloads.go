package downloads

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/james-lawrence/deeppool/internal/x/errorsx"
	"github.com/james-lawrence/deeppool/internal/x/langx"
	"github.com/james-lawrence/deeppool/internal/x/slicesx"
	"github.com/james-lawrence/deeppool/internal/x/sqlx"
	"github.com/james-lawrence/deeppool/internal/x/userx"
	"github.com/james-lawrence/deeppool/tracking"
	"github.com/james-lawrence/torrent"
	"github.com/james-lawrence/torrent/storage"
)

type downloader interface {
	Start(t torrent.Metadata) (dl torrent.Torrent, added bool, err error)
}

func NewDirectoryWatcher(ctx context.Context, q sqlx.Queryer, dl downloader, s storage.ClientImpl) (d Directory, err error) {
	var (
		w *fsnotify.Watcher
	)

	if w, err = fsnotify.NewWatcher(); err != nil {
		return d, err
	}

	return Directory{
		d: dl,
		w: w,
		c: userx.DefaultCacheDirectory(userx.DefaultRelRoot()),
		s: s,
		q: q,
	}.background(ctx), nil
}

type Directory struct {
	d downloader
	q sqlx.Queryer
	w *fsnotify.Watcher
	c string
	s storage.ClientImpl
}

func (t Directory) Add(path string) (err error) {
	defer func() {
		if err == nil {
			log.Println("watching", path)
		}
	}()

	if err = errorsx.Wrapf(t.w.Add(path), "unable to watch: %s", path); err != nil {
		return err
	}

	err = fs.WalkDir(os.DirFS(path), ".", func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".torrent") {
			return nil
		}

		go t.download(context.Background(), filepath.Join(path, name))

		return nil
	})

	return errorsx.Wrap(err, "unable to find existing torrents")
}

// background download
func (t Directory) download(ctx context.Context, path string) {
	meta, err := torrent.NewFromMetaInfoFile(path, torrent.OptionStorage(t.s))
	if err != nil {
		log.Println("unable to process", path, "ignoring", err)
		return
	}

	if info, err := meta.Metainfo().UnmarshalInfo(); err == nil && !info.Private {
		meta = meta.Merge(torrent.OptionTrackers(tracking.PublicTrackers()))
	}

	var (
		md tracking.Metadata
	)

	dl, _, err := t.d.Start(meta)
	if err != nil {
		log.Println(errorsx.Wrap(err, "unable to start torrent"))
		return
	}

	log.Println("wait for torrent info", meta.InfoHash)
	select {
	case <-dl.GotInfo():
	case <-ctx.Done():
		log.Println("failed to retrieve torrent information, manually restart will be required")
		return
	}

	if err = tracking.MetadataInsertWithDefaults(
		ctx,
		t.q,
		tracking.NewMetadata(langx.Autoptr(dl.Metadata().InfoHash),
			tracking.MetadataOptionFromInfo(dl.Info()),
			tracking.MetadataOptionTrackers(slicesx.Flatten(meta.Trackers...)...),
		),
	).Scan(&md); err != nil {
		log.Println(errorsx.Wrap(err, "unable to insert metadata"))
		return
	}

	if err = tracking.MetadataDownloadByID(ctx, t.q, md.ID).Scan(&md); err != nil {
		log.Println(errorsx.Wrap(err, "unable to mark metadata as downloading"))
		return
	}

	pctx, done := context.WithCancel(ctx)
	defer done()

	if err := dl.Tune(torrent.TuneTrackers(slicesx.Flatten(meta.Trackers...))); err != nil {
		log.Println(errorsx.Wrap(err, "unable to tune torrent"))
		return
	} else {
		log.Println("tuned trackers", meta.Trackers)
	}

	errorsx.Log(tracking.Download(pctx, t.q, &md, dl))
}

func (t Directory) background(ctx context.Context) Directory {
	go func() {
		for {
			select {
			case evt := <-t.w.Events:
				switch evt.Op {
				case fsnotify.Create:
				case fsnotify.Chmod:
				case fsnotify.Write:
					continue // explicitly ignored to reduce noise.
				default:
					log.Println("change ignored", evt.Op)
					continue
				}

				go t.download(ctx, evt.Name)
			case err := <-t.w.Errors:
				log.Println("watch error", err)
			case <-ctx.Done():
				log.Println("context completed", ctx.Err())
				return
			}
		}
	}()

	return t
}
