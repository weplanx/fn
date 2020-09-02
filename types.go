package amqpext

import "github.com/streadway/amqp"

type ConsumeOption struct {
	ID        string
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	Callback  func(d amqp.Delivery)
}
