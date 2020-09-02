package main

import (
	"amqp-ext"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	session, err := amqpext.NewSession("amqp://guest:guest@dell")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = session.Channel("default")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = session.Consume(amqpext.ConsumeOption{
		ID:        "default",
		Queue:     "test",
		Consumer:  "consumer-1",
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
	go func() {
		time.Sleep(time.Second * 10)
		session.CloseChannel("default")
	}()
	select {}
}
