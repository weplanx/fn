package main

import (
	"net/http"
	"tencent-workflow-excel/bootstrap"
)

func main() {
	api, err := bootstrap.NewAPI()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/event-invoke", api.Invoke)
	http.ListenAndServe(api.V.Address, nil)
}
