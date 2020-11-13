package controller

import (
	"func-api/application/common"
)

type Controller struct {
	*common.Dependency
}

func New(dep *common.Dependency) *Controller {
	c := new(Controller)
	c.Dependency = dep
	return c
}
