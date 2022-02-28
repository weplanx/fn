//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"openapi/app"
	"openapi/common"
)

func App(value *common.Values) (*gin.Engine, error) {
	wire.Build(
		wire.Struct(new(common.Inject), "*"),
		app.Provides,
	)
	return &gin.Engine{}, nil
}
