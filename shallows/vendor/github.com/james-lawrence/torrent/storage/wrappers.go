package storage

import (
	"fmt"
	"io"
	"os"

	"github.com/anacrolix/missinggo/v2"

	"github.com/james-lawrence/torrent/metainfo"
)

type Client struct {
	ci ClientImpl
}

func NewClient(cl ClientImpl) *Client {
	return &Client{cl}
}

func (cl Client) OpenTorrent(info *metainfo.Info, infoHash metainfo.Hash) (Torrent, error) {
	t, err := cl.ci.OpenTorrent(info, infoHash)
	return Torrent{t}, err
}

func NewTorrent(ti TorrentImpl) Torrent {
	return Torrent{ti: ti}
}

type Torrent struct {
	ti TorrentImpl
}

func (t Torrent) Piece(p metainfo.Piece) Piece {
	return Piece{PieceImpl: t.ti.Piece(p), mip: p}
}

func (t Torrent) Close() error {
	if t.ti == nil {
		return fmt.Errorf("attempted to close a nil torrent storage implementation")
	}
	return t.ti.Close()
}

type Piece struct {
	PieceImpl
	mip metainfo.Piece
}

func (p Piece) WriteAt(b []byte, off int64) (n int, err error) {
	// Callers should not be writing to completed pieces, but it's too
	// expensive to be checking this on every single write using uncached
	// completions.

	// c := p.Completion()
	// if c.Ok && c.Complete {
	// 	err = errors.New("piece already completed")
	// 	return
	// }
	if off+int64(len(b)) > p.mip.Length() {
		panic("write overflows piece")
	}
	b = missinggo.LimitLen(b, p.mip.Length()-off)
	return p.PieceImpl.WriteAt(b, off)
}

func (p Piece) ReadAt(b []byte, off int64) (n int, err error) {
	if off < 0 {
		err = os.ErrInvalid
		return
	}
	if off >= p.mip.Length() {
		err = io.EOF
		return
	}
	b = missinggo.LimitLen(b, p.mip.Length()-off)
	if len(b) == 0 {
		return
	}
	n, err = p.PieceImpl.ReadAt(b, off)
	if n > len(b) {
		panic(n)
	}
	off += int64(n)
	if err == io.EOF && off < p.mip.Length() {
		err = io.ErrUnexpectedEOF
	}
	if err == nil && off >= p.mip.Length() {
		err = io.EOF
	}
	if n == 0 && err == nil {
		err = io.ErrUnexpectedEOF
	}
	if off < p.mip.Length() && err != nil {
		p.MarkNotComplete()
	}
	return
}
