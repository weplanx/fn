package main

import (
	"funcext/app"
	"funcext/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.LoadStorage,
		),
		fx.Invoke(
			app.Application,
		),
	).Run()
}
