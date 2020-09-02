package main

import (
	"amqp-session/client"
	"github.com/sirupsen/logrus"
)

func main() {
	_, err := client.NewSession("amqp://guest:guest@localhost")
	if err != nil {
		logrus.Fatalln(err)
	}
	select {}
}
