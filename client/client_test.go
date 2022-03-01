package client

import (
	"context"
	"github.com/caarlos0/env/v6"
	"os"
	"testing"
)

type XOption struct {
	URL    string `env:"X_URL"`
	Key    string `env:"X_KEY"`
	SECRET string `env:"X_SECRET"`
}

var x *OpenAPI

func TestMain(m *testing.M) {
	var opt XOption
	if err := env.Parse(&opt); err != nil {
		panic(err)
	}
	x = New(opt.URL, SetCertification(opt.Key, opt.SECRET))
	os.Exit(m.Run())
}

func TestOpenAPI_Ping(t *testing.T) {
	result, err := x.Ping(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
