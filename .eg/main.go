package main

import (
	"context"
	"eg/compute/fractal"
	"eg/compute/shallows"
	"log"

	"github.com/egdaemon/eg/runtime/wasi/eg"
	"github.com/egdaemon/eg/runtime/wasi/egenv"
	"github.com/egdaemon/eg/runtime/wasi/eggit"
)

func main() {
	ctx, done := context.WithTimeout(context.Background(), egenv.TTL())
	defer done()

	deb := eg.Container("fractal.ubuntu.24.10")
	err := eg.Perform(
		ctx,
		eggit.AutoClone,
		eg.Build(deb.BuildFromFile(".eg/Containerfile")),
		eg.Parallel(
			eg.Module(
				ctx,
				deb,
				fractal.Generate,
			),
			eg.Module(
				ctx,
				deb,
				shallows.Generate,
			),
		),
		eg.Parallel(
			eg.Module(ctx, deb, fractal.Build),
			eg.Module(
				ctx,
				deb,
				shallows.Compile(),
			),
		),
		eg.Parallel(
			eg.Module(ctx, deb, fractal.Tests),
			eg.Module(ctx, deb, fractal.Linting),
			eg.Module(ctx, deb, shallows.Test()),
		),
		// eg.Module(ctx, deb.OptionLiteral("--publish", "3000:3000"), www.Build, www.Webserver),
	)

	if err != nil {
		log.Fatalln(err)
	}
}
