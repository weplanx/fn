package index

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Controller struct {
	Service *Service
}

func (x *Controller) Index(c *gin.Context) interface{} {
	return gin.H{"time": time.Now()}
}
