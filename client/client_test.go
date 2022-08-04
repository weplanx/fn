package client_test

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/stretchr/testify/assert"
	"github.com/weplanx/openapi/client"
	"os"
	"testing"
)

var local *client.OpenAPI

func TestMain(m *testing.M) {
	var err error
	var e struct {
		Url    string `env:"URL"`
		Key    string `env:"KEY"`
		Secret string `env:"SECRET"`
	}
	if err := env.Parse(&e); err != nil {
		panic(err)
	}
	if local, err = client.New(
		e.Url,
		client.SetApiGateway(e.Key, e.Secret),
	); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestOpenAPI_Ping(t *testing.T) {
	data, err := local.Ping(context.TODO())
	assert.Nil(t, err)
	t.Log(data)
}

func TestOpenAPI_Ip(t *testing.T) {
	data, err := local.Ip(context.TODO(), "119.41.207.227")
	assert.Nil(t, err)
	t.Log(data)
}
