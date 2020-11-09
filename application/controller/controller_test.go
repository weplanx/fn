package controller

import (
	"funcext/bootstrap"
	pb "funcext/router"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
)

var client pb.RouterClient

func TestMain(m *testing.M) {
	os.Chdir("../..")
	cfg, err := bootstrap.LoadConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	grpcConn, err := grpc.Dial(cfg.Listen, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client = pb.NewRouterClient(grpcConn)
	os.Exit(m.Run())
}
