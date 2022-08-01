package client_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weplanx/openapi/client"
	"os"
	"testing"
)

type Option struct {
	URL    string `env:"X_URL"`
	Key    string `env:"X_KEY"`
	SECRET string `env:"X_SECRET"`
}

var local *client.OpenAPI

func TestMain(m *testing.M) {
	var err error
	//var opt Option
	//if err := env.Parse(&opt); err != nil {
	//	panic(err)
	//}
	if local, err = client.New("http://localhost:9000"); err != nil {
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
