//go:build !genieql.ignore
// +build !genieql.ignore

package tracking

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/james-lawrence/deeppool/internal/x/sqlx"
)

// DO NOT EDIT: This File was auto generated by the following command:
// genieql auto -o genieql.gen.go
// invoked by go generate @ tracking/10_generate.genieql.go line 3

// Metadata generated by genieql
//
//easyjson:json
type Metadata struct {
	Bytes            uint64
	CreatedAt        time.Time
	Description      string
	ID               string
	Infohash         string
	PiecesDownloaded uint64
	PiecesPending    uint64
	UpdatedAt        time.Time
}

// MetadataScanner scanner interface.
type MetadataScanner interface {
	Scan(i *Metadata) error
	Next() bool
	Close() error
	Err() error
}

type errMetadataScanner struct {
	e error
}

func (t errMetadataScanner) Scan(i *Metadata) error {
	return t.e
}

func (t errMetadataScanner) Next() bool {
	return false
}

func (t errMetadataScanner) Err() error {
	return t.e
}

func (t errMetadataScanner) Close() error {
	return nil
}

// MetadataScannerStaticColumns generated by genieql
const MetadataScannerStaticColumns = `torrents_metadata."bytes",torrents_metadata."created_at",torrents_metadata."description",torrents_metadata."id",torrents_metadata."infohash",torrents_metadata."pieces_downloaded",torrents_metadata."pieces_pending",torrents_metadata."updated_at"`

// NewMetadataScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewMetadataScannerStatic(rows *sql.Rows, err error) MetadataScanner {
	if err != nil {
		return errMetadataScanner{e: err}
	}

	return metadataScannerStatic{
		Rows: rows,
	}
}

// metadataScannerStatic generated by genieql
type metadataScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t metadataScannerStatic) Scan(i *Metadata) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullTime
		c2 sql.NullString
		c3 sql.NullString
		c4 sql.NullString
		c5 sql.NullInt64
		c6 sql.NullInt64
		c7 sql.NullTime
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7); err != nil {
		return err
	}

	if c0.Valid {
		tmp := uint64(c0.Int64)
		i.Bytes = tmp
	}

	if c1.Valid {
		tmp := c1.Time
		i.CreatedAt = tmp
	}

	if c2.Valid {
		tmp := string(c2.String)
		i.Description = tmp
	}

	if c3.Valid {
		if uid, err := uuid.FromBytes([]byte(c3.String)); err != nil {
			return err
		} else {
			i.ID = uid.String()
		}
	}

	if c4.Valid {
		tmp := string(c4.String)
		i.Infohash = tmp
	}

	if c5.Valid {
		tmp := uint64(c5.Int64)
		i.PiecesDownloaded = tmp
	}

	if c6.Valid {
		tmp := uint64(c6.Int64)
		i.PiecesPending = tmp
	}

	if c7.Valid {
		tmp := c7.Time
		i.UpdatedAt = tmp
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t metadataScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t metadataScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t metadataScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewMetadataScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewMetadataScannerStaticRow(row *sql.Row) MetadataScannerStaticRow {
	return MetadataScannerStaticRow{
		row: row,
	}
}

// MetadataScannerStaticRow generated by genieql
type MetadataScannerStaticRow struct {
	err error
	row *sql.Row
}

// Scan generated by genieql
func (t MetadataScannerStaticRow) Scan(i *Metadata) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullTime
		c2 sql.NullString
		c3 sql.NullString
		c4 sql.NullString
		c5 sql.NullInt64
		c6 sql.NullInt64
		c7 sql.NullTime
	)

	if t.err != nil {
		return t.err
	}

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7); err != nil {
		return err
	}

	if c0.Valid {
		tmp := uint64(c0.Int64)
		i.Bytes = tmp
	}

	if c1.Valid {
		tmp := c1.Time
		i.CreatedAt = tmp
	}

	if c2.Valid {
		tmp := string(c2.String)
		i.Description = tmp
	}

	if c3.Valid {
		if uid, err := uuid.FromBytes([]byte(c3.String)); err != nil {
			return err
		} else {
			i.ID = uid.String()
		}
	}

	if c4.Valid {
		tmp := string(c4.String)
		i.Infohash = tmp
	}

	if c5.Valid {
		tmp := uint64(c5.Int64)
		i.PiecesDownloaded = tmp
	}

	if c6.Valid {
		tmp := uint64(c6.Int64)
		i.PiecesPending = tmp
	}

	if c7.Valid {
		tmp := c7.Time
		i.UpdatedAt = tmp
	}

	return nil
}

// Err set an error to return by scan
func (t MetadataScannerStaticRow) Err(err error) MetadataScannerStaticRow {
	t.err = err
	return t
}

// NewMetadataScannerDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewMetadataScannerDynamic(rows *sql.Rows, err error) MetadataScanner {
	if err != nil {
		return errMetadataScanner{e: err}
	}

	return metadataScannerDynamic{
		Rows: rows,
	}
}

// metadataScannerDynamic generated by genieql
type metadataScannerDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t metadataScannerDynamic) Scan(i *Metadata) error {
	const (
		cn0 = "bytes"
		cn1 = "created_at"
		cn2 = "description"
		cn3 = "id"
		cn4 = "infohash"
		cn5 = "pieces_downloaded"
		cn6 = "pieces_pending"
		cn7 = "updated_at"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
		c1      sql.NullTime
		c2      sql.NullString
		c3      sql.NullString
		c4      sql.NullString
		c5      sql.NullInt64
		c6      sql.NullInt64
		c7      sql.NullTime
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case cn0:
			dst = append(dst, &c0)
		case cn1:
			dst = append(dst, &c1)
		case cn2:
			dst = append(dst, &c2)
		case cn3:
			dst = append(dst, &c3)
		case cn4:
			dst = append(dst, &c4)
		case cn5:
			dst = append(dst, &c5)
		case cn6:
			dst = append(dst, &c6)
		case cn7:
			dst = append(dst, &c7)
		default:
			dst = append(dst, &ignored)
		}
	}

	if err := t.Rows.Scan(dst...); err != nil {
		return err
	}

	for _, column := range columns {
		switch column {
		case cn0:
			if c0.Valid {
				tmp := uint64(c0.Int64)
				i.Bytes = tmp
			}

		case cn1:
			if c1.Valid {
				tmp := c1.Time
				i.CreatedAt = tmp
			}

		case cn2:
			if c2.Valid {
				tmp := string(c2.String)
				i.Description = tmp
			}

		case cn3:
			if c3.Valid {
				if uid, err := uuid.FromBytes([]byte(c3.String)); err != nil {
					return err
				} else {
					i.ID = uid.String()
				}
			}

		case cn4:
			if c4.Valid {
				tmp := string(c4.String)
				i.Infohash = tmp
			}

		case cn5:
			if c5.Valid {
				tmp := uint64(c5.Int64)
				i.PiecesDownloaded = tmp
			}

		case cn6:
			if c6.Valid {
				tmp := uint64(c6.Int64)
				i.PiecesPending = tmp
			}

		case cn7:
			if c7.Valid {
				tmp := c7.Time
				i.UpdatedAt = tmp
			}

		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t metadataScannerDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t metadataScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t metadataScannerDynamic) Next() bool {
	return t.Rows.Next()
}

// MetadataInsertWithDefaultsStaticColumns generated by genieql
const MetadataInsertWithDefaultsStaticColumns = `$1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT`

// MetadataInsertWithDefaultsExplode generated by genieql
func MetadataInsertWithDefaultsExplode(a *Metadata) ([]interface{}, error) {
	var (
		c0 sql.NullInt64  // bytes
		c1 sql.NullString // description
		c2 sql.NullString // infohash
		c3 sql.NullInt64  // pieces_downloaded
		c4 sql.NullInt64  // pieces_pending
	)

	c0.Valid = true
	c0.Int64 = int64(a.Bytes)

	c1.Valid = true
	c1.String = a.Description

	c2.Valid = true
	c2.String = a.Infohash

	c3.Valid = true
	c3.Int64 = int64(a.PiecesDownloaded)

	c4.Valid = true
	c4.Int64 = int64(a.PiecesPending)

	return []interface{}{c0, c1, c2, c3, c4}, nil
}

// MetadataInsertWithDefaults generated by genieql
func MetadataInsertWithDefaults(ctx context.Context, q sqlx.Queryer, a Metadata) MetadataScannerStaticRow {
	const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
	var (
		c0 sql.NullInt64  // bytes
		c1 sql.NullString // description
		c2 sql.NullString // infohash
		c3 sql.NullInt64  // pieces_downloaded
		c4 sql.NullInt64
	)
	c0.Valid = true
	c0.Int64 = int64(a.Bytes)
	c1.Valid = true
	c1.String = a.Description
	c2.Valid = true
	c2.String = a.Infohash
	c3.Valid = true
	c3.Int64 = int64(a.PiecesDownloaded)
	c4.Valid = true
	c4.Int64 = int64(a.PiecesPending) // pieces_pending
	return NewMetadataScannerStaticRow(q.QueryRowContext(ctx, query, c0, c1, c2, c3, c4))
}

// MetadataBatchInsertWithDefaults generated by genieql
func NewMetadataBatchInsertWithDefaults(ctx context.Context, q sqlx.Queryer, p ...Metadata) MetadataScanner {
	return &metadataBatchInsertWithDefaults{ctx: ctx, q: q, remaining: p}
}

type metadataBatchInsertWithDefaults struct {
	ctx       context.Context
	q         sqlx.Queryer
	remaining []Metadata
	scanner   MetadataScanner
}

func (t *metadataBatchInsertWithDefaults) Scan(p *Metadata) error {
	return t.scanner.Scan(p)
}

func (t *metadataBatchInsertWithDefaults) Err() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Err()
}

func (t *metadataBatchInsertWithDefaults) Close() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Close()
}

func (t *metadataBatchInsertWithDefaults) Next() bool {
	var advanced bool
	if t.scanner != nil && t.scanner.Next() {
		return true
	}
	if len(t.remaining) > 0 && t.Close() == nil {
		t.scanner, t.remaining, advanced = t.advance(t.remaining...)
		return advanced && t.scanner.Next()
	}
	return false
}

func (t *metadataBatchInsertWithDefaults) advance(p ...Metadata) (MetadataScanner, []Metadata, bool) {
	transform := func(p Metadata) (c0 sql.NullInt64, c1 sql.NullString, c2 sql.NullString, c3 sql.NullInt64, c4 sql.NullInt64, err error) {
		c0.Valid = true
		c0.Int64 = int64(p.Bytes)
		c1.Valid = true
		c1.String = p.Description
		c2.Valid = true
		c2.String = p.Infohash
		c3.Valid = true
		c3.Int64 = int64(p.PiecesDownloaded)
		c4.Valid = true
		c4.Int64 = int64(p.PiecesPending)
		return c0, c1, c2, c3, c4, nil
	}
	switch len(p) {
	case 0:
		return nil, []Metadata(nil), false
	case 1:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4)), p[1:], true
	case 2:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4)), p[2:], true
	case 3:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4)), p[3:], true
	case 4:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4)), p[4:], true
	case 5:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT),($21,DEFAULT,$22,DEFAULT,$23,$24,$25,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			r4c0 sql.NullInt64
			r4c1 sql.NullString
			r4c2 sql.NullString
			r4c3 sql.NullInt64
			r4c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r4c0, r4c1, r4c2, r4c3, r4c4, err = transform(p[4]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4, r4c0, r4c1, r4c2, r4c3, r4c4)), p[5:], true
	case 6:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT),($21,DEFAULT,$22,DEFAULT,$23,$24,$25,DEFAULT),($26,DEFAULT,$27,DEFAULT,$28,$29,$30,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			r4c0 sql.NullInt64
			r4c1 sql.NullString
			r4c2 sql.NullString
			r4c3 sql.NullInt64
			r4c4 sql.NullInt64
			r5c0 sql.NullInt64
			r5c1 sql.NullString
			r5c2 sql.NullString
			r5c3 sql.NullInt64
			r5c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r4c0, r4c1, r4c2, r4c3, r4c4, err = transform(p[4]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r5c0, r5c1, r5c2, r5c3, r5c4, err = transform(p[5]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4, r4c0, r4c1, r4c2, r4c3, r4c4, r5c0, r5c1, r5c2, r5c3, r5c4)), p[6:], true
	case 7:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT),($21,DEFAULT,$22,DEFAULT,$23,$24,$25,DEFAULT),($26,DEFAULT,$27,DEFAULT,$28,$29,$30,DEFAULT),($31,DEFAULT,$32,DEFAULT,$33,$34,$35,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			r4c0 sql.NullInt64
			r4c1 sql.NullString
			r4c2 sql.NullString
			r4c3 sql.NullInt64
			r4c4 sql.NullInt64
			r5c0 sql.NullInt64
			r5c1 sql.NullString
			r5c2 sql.NullString
			r5c3 sql.NullInt64
			r5c4 sql.NullInt64
			r6c0 sql.NullInt64
			r6c1 sql.NullString
			r6c2 sql.NullString
			r6c3 sql.NullInt64
			r6c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r4c0, r4c1, r4c2, r4c3, r4c4, err = transform(p[4]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r5c0, r5c1, r5c2, r5c3, r5c4, err = transform(p[5]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r6c0, r6c1, r6c2, r6c3, r6c4, err = transform(p[6]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4, r4c0, r4c1, r4c2, r4c3, r4c4, r5c0, r5c1, r5c2, r5c3, r5c4, r6c0, r6c1, r6c2, r6c3, r6c4)), p[7:], true
	case 8:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT),($21,DEFAULT,$22,DEFAULT,$23,$24,$25,DEFAULT),($26,DEFAULT,$27,DEFAULT,$28,$29,$30,DEFAULT),($31,DEFAULT,$32,DEFAULT,$33,$34,$35,DEFAULT),($36,DEFAULT,$37,DEFAULT,$38,$39,$40,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			r4c0 sql.NullInt64
			r4c1 sql.NullString
			r4c2 sql.NullString
			r4c3 sql.NullInt64
			r4c4 sql.NullInt64
			r5c0 sql.NullInt64
			r5c1 sql.NullString
			r5c2 sql.NullString
			r5c3 sql.NullInt64
			r5c4 sql.NullInt64
			r6c0 sql.NullInt64
			r6c1 sql.NullString
			r6c2 sql.NullString
			r6c3 sql.NullInt64
			r6c4 sql.NullInt64
			r7c0 sql.NullInt64
			r7c1 sql.NullString
			r7c2 sql.NullString
			r7c3 sql.NullInt64
			r7c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r4c0, r4c1, r4c2, r4c3, r4c4, err = transform(p[4]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r5c0, r5c1, r5c2, r5c3, r5c4, err = transform(p[5]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r6c0, r6c1, r6c2, r6c3, r6c4, err = transform(p[6]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r7c0, r7c1, r7c2, r7c3, r7c4, err = transform(p[7]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4, r4c0, r4c1, r4c2, r4c3, r4c4, r5c0, r5c1, r5c2, r5c3, r5c4, r6c0, r6c1, r6c2, r6c3, r6c4, r7c0, r7c1, r7c2, r7c3, r7c4)), p[8:], true
	case 9:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT),($21,DEFAULT,$22,DEFAULT,$23,$24,$25,DEFAULT),($26,DEFAULT,$27,DEFAULT,$28,$29,$30,DEFAULT),($31,DEFAULT,$32,DEFAULT,$33,$34,$35,DEFAULT),($36,DEFAULT,$37,DEFAULT,$38,$39,$40,DEFAULT),($41,DEFAULT,$42,DEFAULT,$43,$44,$45,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			r4c0 sql.NullInt64
			r4c1 sql.NullString
			r4c2 sql.NullString
			r4c3 sql.NullInt64
			r4c4 sql.NullInt64
			r5c0 sql.NullInt64
			r5c1 sql.NullString
			r5c2 sql.NullString
			r5c3 sql.NullInt64
			r5c4 sql.NullInt64
			r6c0 sql.NullInt64
			r6c1 sql.NullString
			r6c2 sql.NullString
			r6c3 sql.NullInt64
			r6c4 sql.NullInt64
			r7c0 sql.NullInt64
			r7c1 sql.NullString
			r7c2 sql.NullString
			r7c3 sql.NullInt64
			r7c4 sql.NullInt64
			r8c0 sql.NullInt64
			r8c1 sql.NullString
			r8c2 sql.NullString
			r8c3 sql.NullInt64
			r8c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r4c0, r4c1, r4c2, r4c3, r4c4, err = transform(p[4]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r5c0, r5c1, r5c2, r5c3, r5c4, err = transform(p[5]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r6c0, r6c1, r6c2, r6c3, r6c4, err = transform(p[6]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r7c0, r7c1, r7c2, r7c3, r7c4, err = transform(p[7]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r8c0, r8c1, r8c2, r8c3, r8c4, err = transform(p[8]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4, r4c0, r4c1, r4c2, r4c3, r4c4, r5c0, r5c1, r5c2, r5c3, r5c4, r6c0, r6c1, r6c2, r6c3, r6c4, r7c0, r7c1, r7c2, r7c3, r7c4, r8c0, r8c1, r8c2, r8c3, r8c4)), p[9:], true
	default:
		const query = `INSERT INTO "torrents_metadata" ("bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at") VALUES ($1,DEFAULT,$2,DEFAULT,$3,$4,$5,DEFAULT),($6,DEFAULT,$7,DEFAULT,$8,$9,$10,DEFAULT),($11,DEFAULT,$12,DEFAULT,$13,$14,$15,DEFAULT),($16,DEFAULT,$17,DEFAULT,$18,$19,$20,DEFAULT),($21,DEFAULT,$22,DEFAULT,$23,$24,$25,DEFAULT),($26,DEFAULT,$27,DEFAULT,$28,$29,$30,DEFAULT),($31,DEFAULT,$32,DEFAULT,$33,$34,$35,DEFAULT),($36,DEFAULT,$37,DEFAULT,$38,$39,$40,DEFAULT),($41,DEFAULT,$42,DEFAULT,$43,$44,$45,DEFAULT),($46,DEFAULT,$47,DEFAULT,$48,$49,$50,DEFAULT) RETURNING "bytes","created_at","description","id","infohash","pieces_downloaded","pieces_pending","updated_at"`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullString
			r0c2 sql.NullString
			r0c3 sql.NullInt64
			r0c4 sql.NullInt64
			r1c0 sql.NullInt64
			r1c1 sql.NullString
			r1c2 sql.NullString
			r1c3 sql.NullInt64
			r1c4 sql.NullInt64
			r2c0 sql.NullInt64
			r2c1 sql.NullString
			r2c2 sql.NullString
			r2c3 sql.NullInt64
			r2c4 sql.NullInt64
			r3c0 sql.NullInt64
			r3c1 sql.NullString
			r3c2 sql.NullString
			r3c3 sql.NullInt64
			r3c4 sql.NullInt64
			r4c0 sql.NullInt64
			r4c1 sql.NullString
			r4c2 sql.NullString
			r4c3 sql.NullInt64
			r4c4 sql.NullInt64
			r5c0 sql.NullInt64
			r5c1 sql.NullString
			r5c2 sql.NullString
			r5c3 sql.NullInt64
			r5c4 sql.NullInt64
			r6c0 sql.NullInt64
			r6c1 sql.NullString
			r6c2 sql.NullString
			r6c3 sql.NullInt64
			r6c4 sql.NullInt64
			r7c0 sql.NullInt64
			r7c1 sql.NullString
			r7c2 sql.NullString
			r7c3 sql.NullInt64
			r7c4 sql.NullInt64
			r8c0 sql.NullInt64
			r8c1 sql.NullString
			r8c2 sql.NullString
			r8c3 sql.NullInt64
			r8c4 sql.NullInt64
			r9c0 sql.NullInt64
			r9c1 sql.NullString
			r9c2 sql.NullString
			r9c3 sql.NullInt64
			r9c4 sql.NullInt64
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, err = transform(p[0]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, err = transform(p[1]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r2c0, r2c1, r2c2, r2c3, r2c4, err = transform(p[2]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r3c0, r3c1, r3c2, r3c3, r3c4, err = transform(p[3]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r4c0, r4c1, r4c2, r4c3, r4c4, err = transform(p[4]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r5c0, r5c1, r5c2, r5c3, r5c4, err = transform(p[5]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r6c0, r6c1, r6c2, r6c3, r6c4, err = transform(p[6]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r7c0, r7c1, r7c2, r7c3, r7c4, err = transform(p[7]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r8c0, r8c1, r8c2, r8c3, r8c4, err = transform(p[8]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		if r9c0, r9c1, r9c2, r9c3, r9c4, err = transform(p[9]); err != nil {
			return NewMetadataScannerStatic(nil, err), []Metadata(nil), false
		}
		return NewMetadataScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r1c0, r1c1, r1c2, r1c3, r1c4, r2c0, r2c1, r2c2, r2c3, r2c4, r3c0, r3c1, r3c2, r3c3, r3c4, r4c0, r4c1, r4c2, r4c3, r4c4, r5c0, r5c1, r5c2, r5c3, r5c4, r6c0, r6c1, r6c2, r6c3, r6c4, r7c0, r7c1, r7c2, r7c3, r7c4, r8c0, r8c1, r8c2, r8c3, r8c4, r9c0, r9c1, r9c2, r9c3, r9c4)), []Metadata(nil), false
	}
}
