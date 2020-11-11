package common

import (
	"func-api/application/service/excel"
	"func-api/application/service/storage"
	"func-api/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Dependency struct {
	fx.In

	Config  *config.Config
	Storage *storage.Service
	Excel   *excel.Service
}

func Handle(handlersFn interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if method, ok := handlersFn.(func(ctx *gin.Context) interface{}); ok {
			ctx.JSON(200, method(ctx))
		}
	}
}
