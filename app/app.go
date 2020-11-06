package app

import (
	"funcext/app/common"
	"funcext/app/controller"
	pb "funcext/router"
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func Application(dep common.Dependency) (err error) {
	cfg := dep.Config
	if cfg.Debug != "" {
		go http.ListenAndServe(cfg.Debug, nil)
	}
	var listen net.Listener
	listen, err = net.Listen("tcp", cfg.Listen)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	pb.RegisterRouterServer(
		server,
		controller.New(&dep),
	)
	go server.Serve(listen)
	return
}
