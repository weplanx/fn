package controller

import (
	"func-api/application/common"
	"github.com/gin-gonic/gin"
)

type controller struct {
	dep *common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.dep = dep
	return c
}

func (c *controller) error(err error) interface{} {
	return gin.H{
		"error": 1,
		"msg":   err.Error(),
	}
}

func (c *controller) ok() interface{} {
	return gin.H{
		"error": 0,
		"msg":   "ok",
	}
}

func (c *controller) result(data interface{}) interface{} {
	return gin.H{
		"error": 0,
		"data":  data,
	}
}
