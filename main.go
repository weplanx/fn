package main

import (
	"github.com/weplanx/openapi/bootstrap"
)

func main() {
	api, err := bootstrap.NewAPI()
	if err != nil {
		panic(err)
	}

	h, err := api.Routes()
	if err != nil {
		panic(err)
	}

	h.Spin()
}
