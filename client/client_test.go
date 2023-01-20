package client_test

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/stretchr/testify/assert"
	"github.com/weplanx/openapi/api/excel"
	"github.com/weplanx/openapi/client"
	"os"
	"testing"
)

var x *client.Client

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
	if x, err = client.New(e.Url, client.SetApiGateway(e.Key, e.Secret)); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestClient_Ping(t *testing.T) {
	data, err := x.Ping(context.TODO())
	assert.Nil(t, err)
	t.Log(data)
}

func TestClient_GetIp(t *testing.T) {
	data, err := x.GetIp(context.TODO(), "119.41.207.227")
	assert.Nil(t, err)
	t.Log(data)
}

func TestClient_GetCountries(t *testing.T) {
	data, err := x.GetCountries(context.TODO(), []string{"iso3"})
	assert.Nil(t, err)
	t.Log(data)
}

func TestClient_GetStates(t *testing.T) {
	data, err := x.GetStates(context.TODO(), "CN", []string{"type"})
	assert.Nil(t, err)
	t.Log(data)
}

func TestClient_GetCities(t *testing.T) {
	data, err := x.GetCities(context.TODO(), "CN", "AH", []string{"latitude"})
	assert.Nil(t, err)
	t.Log(data)
}

func TestClient_CreateExcel(t *testing.T) {
	r, err := x.CreateExcel(context.TODO(), excel.CreateDto{
		Sheets: []excel.Sheet{
			{
				Name: "测试1",
				Data: [][]interface{}{
					{"Name", "Age"},
					{"AC", 15},
					{"HJ", 16},
					{"UI", 17},
				},
			},
		},
	})
	assert.Nil(t, err)
	t.Log(r)
}
