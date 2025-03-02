package release

import (
	"context"
	"eg/compute/tarball"

	"github.com/egdaemon/eg/runtime/wasi/eg"
	"github.com/egdaemon/eg/runtime/wasi/egenv"
	"github.com/egdaemon/eg/runtime/wasi/shell"
	"github.com/egdaemon/eg/runtime/x/wasi/egbug"
)

// prints the directory tree of the system cache
func SystemCache(dir string) eg.OpFn {
	return func(ctx context.Context, op eg.Op) error {
		privileged := shell.Runtime().Privileged().Lenient(true).Directory("/")
		return shell.Run(
			ctx,
			privileged.Newf("tree -a %s", dir),
		)
	}
}

func Tarball(ctx context.Context, op eg.Op) error {
	archive := tarball.GitPattern("retrovibed")
	return eg.Perform(
		ctx,
		egbug.Log("CHECKPOINT 0", archive),
		tarball.Pack(archive),
		egbug.Log("CHECKPOINT 1"),
		SystemCache(egenv.CacheDirectory(".eg", "tarball")),
		shell.Op(
			shell.Newf("ls -lha %s", tarball.Archive(archive)),
		),
		// egbug.FileTree,
		tarball.SHA256Op(archive),
		egbug.Log("CHECKPOINT 2"),
	)
}

func Release(ctx context.Context, op eg.Op) error {
	return eg.Perform(
		ctx,
		tarball.Github(
			tarball.Archive(tarball.GitPattern("retrovibed")),
			egenv.CacheDirectory("flatpak.client.yml"),
			egenv.CacheDirectory("flatpak.daemon.yml"),
		),
	)
}
