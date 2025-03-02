package fractal

import (
	"context"
	"eg/compute/tarball"
	"os"

	"github.com/egdaemon/eg/runtime/wasi/eg"
	"github.com/egdaemon/eg/runtime/wasi/egenv"
	"github.com/egdaemon/eg/runtime/wasi/shell"
	"github.com/egdaemon/eg/runtime/x/wasi/egflatpak"
	"github.com/egdaemon/eg/runtime/x/wasi/egfs"
)

func tarballpattern() string {
	return tarball.GitPattern("retrovibed")
}

func flutterRuntime() shell.Command {
	return shell.Runtime().Directory(egenv.WorkingDirectory("fractal")).Environ("PUB_CACHE", egenv.CacheDirectory(".eg", "dart"))
}

func Build(ctx context.Context, _ eg.Op) error {
	runtime := flutterRuntime()
	return shell.Run(
		ctx,
		runtime.New("flutter create --platforms=linux ."),
		runtime.Newf("flutter build bundle"),
		runtime.Newf("flutter build linux"),
	)
}

func Tests(ctx context.Context, _ eg.Op) error {
	runtime := flutterRuntime()
	return shell.Run(
		ctx,
		runtime.New("flutter test"),
	)
}

func Linting(ctx context.Context, _ eg.Op) error {
	runtime := flutterRuntime()
	return shell.Run(
		ctx,
		runtime.New("flutter analyze"),
	)
}

func Generate(ctx context.Context, _ eg.Op) error {
	return shell.Run(
		ctx,
		shell.New("PATH=\"${PATH}:${HOME}/.pub-cache/bin\" protoc --dart_out=grpc:fractal/lib/media -I.proto .proto/media.proto"),
		shell.New("PATH=\"${PATH}:${HOME}/.pub-cache/bin\" protoc --dart_out=grpc:fractal/lib/rss -I.proto .proto/rss.proto"),
	)
}

func Install(ctx context.Context, op eg.Op) error {
	runtime := shell.Runtime()
	dstdir := tarball.Path(tarballpattern())
	builddir := egenv.WorkingDirectory("fractal", "build", egfs.FindFirst(os.DirFS(egenv.WorkingDirectory("fractal", "build")), "bundle"))

	return shell.Run(
		ctx,
		runtime.Newf("mkdir -p %s", dstdir),
		runtime.Newf("cp -R %s/* %s", builddir, dstdir),
	)
}

func FlatpakDir(ctx context.Context, op eg.Op) error {
	runtime := shell.Runtime()
	builddir := egenv.WorkingDirectory("fractal", "build", egfs.FindFirst(os.DirFS(egenv.WorkingDirectory("fractal", "build")), "bundle"))

	b := egflatpak.New(
		"space.retrovibe.Daemon", "fractal",
		egflatpak.Option().SDK("org.gnome.Sdk", "47").Runtime("org.gnome.Platform", "47").
			Modules(egflatpak.ModuleCopy(builddir)).
			AllowWayland().
			AllowDRI().
			AllowNetwork().
			AllowDownload().
			AllowMusic().
			AllowVideos()...)

	if err := egflatpak.Build(ctx, runtime, b); err != nil {
		return err
	}

	return nil
}

func FlatpakManifestDaemon(ctx context.Context, o eg.Op) error {
	pattern := tarballpattern()

	b := egflatpak.New(
		"space.retrovibe.Daemon", "shallows",
		egflatpak.Option().SDK("org.gnome.Sdk", "47").Runtime("org.gnome.Platform", "47").
			Modules(
				egflatpak.NewModule("retrovibed", "simple", egflatpak.ModuleOptions().Commands(
					"ls -lha .",
					"mv lib* /usr/lib",
					"mv duckdb* /usr/lib",
					"pwd",
					"ls -lha /app/lib",
					"ls -lha .",
					"cp -r . /app/bin",
				).Sources(
					egflatpak.SourceTarball(tarball.GithubDownloadURL(pattern), tarball.SHA256(pattern)),
					egflatpak.SourceTarball(
						"https://github.com/duckdb/duckdb/releases/download/v1.1.3/libduckdb-linux-amd64.zip",
						"81199bf01b6d49941a38f426cad60e73c1ccd43f1f769a65ed8097d53fc7e40b",
					),
					// flatpak.SourceGit("https://github.com/duckdb/duckdb.git", "v1.1.3"),
				)...),
				// egflatpak.ModuleTarball(tarball.GithubDownloadURL(pattern), tarball.SHA256(pattern)),
			).
			AllowWayland().
			AllowDRI().
			AllowNetwork().
			AllowDownload().
			AllowMusic().
			AllowVideos()...)

	return egflatpak.ManifestOp(egenv.CacheDirectory("flatpak.daemon.yml"), b)(ctx, o)
}

func FlatpakManifestClient(ctx context.Context, o eg.Op) error {
	pattern := tarballpattern()
	// https://github.com/duckdb/duckdb/releases/download/v1.1.3/libduckdb-linux-amd64.zip
	// /app/lib
	b := egflatpak.New(
		"space.retrovibe.Client", "fractal",
		egflatpak.Option().SDK("org.gnome.Sdk", "47").Runtime("org.gnome.Platform", "47").
			Modules(
				egflatpak.ModuleTarball(tarball.GithubDownloadURL(pattern), tarball.SHA256(pattern)),
			).
			AllowWayland().
			AllowDRI().
			AllowNetwork().
			AllowDownload().
			AllowMusic().
			AllowVideos()...)

	return egflatpak.ManifestOp(egenv.CacheDirectory("flatpak.client.yml"), b)(ctx, o)
}
