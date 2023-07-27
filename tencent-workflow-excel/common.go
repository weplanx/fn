package main

import (
	"github.com/caarlos0/env/v9"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

type Inject struct {
	V      *Values
	Client *cos.Client
}

type Values struct {
	Event string `env:"SCF_CUSTOM_CONTAINER_EVENT"`
	Cos   struct {
		Url       string `env:"URL"`
		SecretId  string `env:"SECRETID"`
		SecretKey string `env:"SECRETKEY"`
	} `envPrefix:"COS_"`
}

func Load(x *Inject) (err error) {
	x.V = new(Values)
	if err = env.Parse(x.V); err != nil {
		return
	}

	u, _ := url.Parse(x.V.Cos.Url)
	x.Client = cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  x.V.Cos.SecretId,
			SecretKey: x.V.Cos.SecretKey,
		},
	})
	return
}
