package main

import (
	"funcext/application"
	"funcext/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeStorage,
			bootstrap.InitializeExcel,
		),
		fx.Invoke(
			application.Application,
		),
	).Run()
}
