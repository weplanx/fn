package controller

import (
	"func-api/application/common"
)

type controller struct {
	dep *common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.dep = dep
	return c
}
