package main

import "openapi/bootstrap"

func main() {
	values, err := bootstrap.SetValues()
	if err != nil {
		panic(err)
	}
	app, err := App(values)
	if err != nil {
		panic(err)
	}
	app.Run(":9000")
}
