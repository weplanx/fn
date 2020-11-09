package application

import (
	"funcext/application/common"
	"funcext/application/controller"
	pb "funcext/router"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcZap.StreamServerInterceptor(logger),
		),
		grpc.UnaryInterceptor(
			grpcZap.UnaryServerInterceptor(logger),
		),
	)
	pb.RegisterRouterServer(
		server,
		controller.New(&dep),
	)
	go server.Serve(listen)
	return
}
