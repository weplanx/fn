package main

import (
	"github.com/weplanx/fn/bootstrap"
	"net/http"
)

func main() {
	api, err := bootstrap.NewAPI()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/event-invoke", api.EventInvoke)
	http.ListenAndServe(api.V.Address, nil)
}
