package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/weplanx/go/route"
	"github.com/weplanx/openapi/app/index"
	"github.com/weplanx/openapi/common"
)

var Provides = wire.NewSet(
	index.Provides,
	New,
)

func New(
	values *common.Values,
	index *index.Controller,
) *gin.Engine {
	r := globalMiddleware(gin.New(), values)
	r.GET("/", route.Use(index.Index))
	return r
}
