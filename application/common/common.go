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
			result := method(ctx)
			switch val := result.(type) {
			case bool:
				if val {
					ctx.Status(200)
				} else {
					ctx.Status(500)
				}
				break
			case error:
				ctx.JSON(400, val.Error())
				break
			default:
				ctx.JSON(200, val)
			}
		} else {
			ctx.Status(404)
		}
	}
}
