package app

import (
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/weplanx/openapi/common"
	"go.uber.org/zap"
	"os"
	"time"
)

func globalMiddleware(r *gin.Engine, values *common.Values) *gin.Engine {

	if os.Getenv("GIN_MODE") == "release" {
		logger, _ := zap.NewProduction()
		r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		r.Use(requestid.New())
	} else {
		r.Use(gin.Logger())
	}
	r.SetTrustedProxies(values.TrustedProxies)
	r.Use(gin.CustomRecovery(catchError))
	return r
}

func catchError(c *gin.Context, err interface{}) {
	c.AbortWithStatusJSON(500, gin.H{
		"message": err,
	})
}
