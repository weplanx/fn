package main

import (
	"amqp-session"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	client, err := amqpext.NewSession("amqp://guest:guest@dell")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = client.NewChannel("default")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = client.NewConsume(amqpext.ConsumeOption{
		ChannelID: "default",
		Queue:     "test",
		Consumer:  "consumer",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Callback: func(d amqp.Delivery) {
			logrus.Info("Received a message:", string(d.Body))
		},
	})
	if err != nil {
		logrus.Fatalln(err)
	}
	select {}
}
