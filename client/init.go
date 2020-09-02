package client

import (
	"amqp-session/types"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Session struct {
	url             string
	conn            *amqp.Connection
	connected       bool
	channel         map[string]*amqp.Channel
	notifyConnClose chan *amqp.Error
	notifyChanClose map[string]chan *amqp.Error
	consumeOptions  map[string]*types.ConsumeOption
}

func NewSession(url string) (session *Session, err error) {
	session = new(Session)
	session.url = url
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	session.conn = conn
	session.connected = true
	session.notifyConnClose = make(chan *amqp.Error)
	conn.NotifyClose(session.notifyConnClose)
	go session.listen()
	session.channel = make(map[string]*amqp.Channel)
	session.notifyChanClose = make(map[string]chan *amqp.Error)
	session.consumeOptions = make(map[string]*types.ConsumeOption)
	return
}

func (c *Session) listen() {
	select {
	case <-c.notifyConnClose:
		logrus.Error("AMQP connection has been disconnected")
		c.reconnected()
	}
}

func (c *Session) reconnected() {
	c.connected = false
	count := 0
	for {
		time.Sleep(time.Second * 5)
		count++
		logrus.Info("Trying to reconnect:", count)
		conn, err := amqp.Dial(c.url)
		if err != nil {
			logrus.Error(err)
			continue
		}
		c.conn = conn
		c.connected = true
		c.notifyConnClose = make(chan *amqp.Error)
		conn.NotifyClose(c.notifyConnClose)
		go c.listen()
		for ID := range c.channel {
			err = c.NewChannel(ID)
			if err != nil {
				continue
			}
		}
		for _, option := range c.consumeOptions {
			err = c.NewConsume(*option)
			if err != nil {
				continue
			}
		}
		logrus.Info("Attempt to reconnect successfully")
		break
	}
}

func (c *Session) NewChannel(ID string) (err error) {
	channel, err := c.conn.Channel()
	if err != nil {
		return
	}
	c.channel[ID] = channel
	c.notifyChanClose[ID] = make(chan *amqp.Error)
	channel.NotifyClose(c.notifyChanClose[ID])
	return
}

func (c *Session) NewConsume(option types.ConsumeOption) (err error) {
	c.consumeOptions[option.Consumer] = &option
	msgs, err := c.channel[option.ChannelID].Consume(
		option.Queue,
		option.Consumer,
		option.AutoAck,
		option.Exclusive,
		option.NoLocal,
		option.NoWait,
		option.Args,
	)
	go func() {
		for d := range msgs {
			option.Callback(d)
		}
	}()
	return
}
