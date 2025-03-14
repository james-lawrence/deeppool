package main

import (
	"context"
	"log"

	"eg/compute/fractal"
	"eg/compute/shallows"

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
			eg.Module(ctx, deb, fractal.Generate),
			eg.Module(ctx, deb, shallows.Generate),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}
}
