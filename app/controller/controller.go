package controller

import (
	"funcext/app/common"
	pb "funcext/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	dep *common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.dep = dep
	return c
}
