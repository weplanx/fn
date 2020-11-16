package main

import (
	"func-api/application"
	"func-api/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeStorage,
			bootstrap.InitializeExcel,
			bootstrap.InitializeQRCode,
			bootstrap.HttpServer,
		),
		fx.Invoke(application.Application),
	).Run()
}
