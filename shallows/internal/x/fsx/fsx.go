package fsx

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/james-lawrence/deeppool/internal/x/debugx"
)

func ErrIsNotExist(err error) error {
	if errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func IgnoreIsNotExist(err error) error {
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}

func IgnoreIsExist(err error) error {
	if errors.Is(err, os.ErrExist) {
		return nil
	}

	return err
}

func AutoCached(path string, gen func() ([]byte, error)) (s []byte, err error) {
	if s, err = os.ReadFile(path); err == nil {
		return s, nil
	}

	if s, err = gen(); err != nil {
		return nil, err
	}

	if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}

	if err = os.WriteFile(path, s, 0600); err != nil {
		return nil, err
	}

	return s, err
}

// IsRegularFile returns true IFF a non-directory file exists at the provided path.
func IsRegularFile(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

type Virtual interface {
	// returns the path rooted at the virtual fs from the fragments.
	Path(rel ...string) string
	MkDirAll(path string, perm os.FileMode) error
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
}

func VirtualAsFS(v Virtual) fs.FS {
	return vstoragefs{Virtual: v, pathrewrite: func(s string) string { return s }}
}

func VirtualAsFSWithRewrite(v Virtual, rewrite func(s string) string) fs.FS {
	return vstoragefs{Virtual: v, pathrewrite: rewrite}
}

func DirVirtual(dir string) Virtual {
	return dirvirt{root: dir}
}

type dirvirt struct {
	root string
}

func (t dirvirt) Path(rel ...string) string {
	return filepath.Join(t.root, filepath.Join(rel...))
}

func (t dirvirt) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(filepath.Join(t.root, name), flag, perm)
}

func (t dirvirt) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, filepath.Join(t.root, newpath))
}

func (t dirvirt) MkDirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(filepath.Join(t.root, path), perm)
}

type vstoragefs struct {
	Virtual
	pathrewrite func(s string) string
}

func (t vstoragefs) Open(name string) (fs.File, error) {
	debugx.Println("opening", name, "as", t.pathrewrite(name))
	return t.Virtual.OpenFile(t.pathrewrite(name), os.O_RDONLY, 0600)
}
