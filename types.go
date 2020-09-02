package amqp_session

import "github.com/streadway/amqp"

type ConsumeOption struct {
	ChannelID string
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	Callback  func(d amqp.Delivery)
}
