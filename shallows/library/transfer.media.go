package library

import (
	"context"
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"iter"
	"log"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/james-lawrence/deeppool/internal/x/errorsx"
	"github.com/james-lawrence/deeppool/internal/x/fsx"
)

type Transfered struct {
	Path     string
	Mimetype *mimetype.MIME
	MD5      hash.Hash
	Bytes    uint64
}

// used to import files from a given directory tree into the library.
// it'll walk the tree, create a copy of each file into the media based on the contents md5.
func TransferMedia(ctx context.Context, rootstore fsx.Virtual, subtree string) iter.Seq[Transfered] {
	fsi := os.DirFS(rootstore.Path(subtree))

	return func(yield func(Transfered) bool) {
		err := fs.WalkDir(fsi, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			tmp, err := os.MkdirTemp(rootstore.Path(".tmp"), "transferring.*")
			if err != nil {
				return err
			}

			var (
				tx = Transfered{
					Path: tmp,
					MD5:  md5.New(),
				}
			)

			log.Println("initiated copy", path, "to", tmp)
			defer log.Println("completed copy", path, "to", tmp)

			src, err := os.Open(path)
			if err != nil {
				return err
			}
			defer src.Close()

			dst, err := os.OpenFile(tmp, os.O_CREATE|os.O_RDWR|os.O_EXCL, d.Type().Perm())
			if err != nil {
				return err
			}
			defer dst.Close()

			if n, err := io.Copy(io.MultiWriter(dst, tx.MD5), src); err != nil {
				return err
			} else {
				tx.Bytes = uint64(n)
			}

			if !yield(tx) {
				return fmt.Errorf("failed to yield transferred media")
			}

			return nil
		})
		errorsx.Log(err)
	}
}
