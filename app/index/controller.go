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

type IpQuery struct {
	Value string `form:"value" binding:"required,ip4_addr"`
}

func (x *Controller) Ip(c *gin.Context) interface{} {
	var query IpQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		return err
	}
	data, err := x.Service.FindIp(c.Request.Context(), query.Value)
	if err != nil {
		return err
	}
	return data
}
