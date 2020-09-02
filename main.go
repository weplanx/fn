package main

import (
	"amqp-session/client"
	"github.com/sirupsen/logrus"
)

func main() {
	session, err := client.NewSession("amqp://guest:guest@localhost")
	if err != nil {
		logrus.Fatalln(err)
	}
	defer session.Wait()
	err = session.NewChannel("test")
	if err != nil {
		logrus.Fatalln(err)
	}
}
