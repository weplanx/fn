package client

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Session struct {
	conn            *amqp.Connection
	channel         map[string]*amqp.Channel
	notifyConnClose chan *amqp.Error
	notifyChanClose map[string]chan *amqp.Error
}

func NewSession(url string) (session *Session, err error) {
	session = new(Session)
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	session.conn = conn
	session.notifyConnClose = make(chan *amqp.Error)
	conn.NotifyClose(session.notifyConnClose)
	session.channel = make(map[string]*amqp.Channel)
	session.notifyChanClose = make(map[string]chan *amqp.Error)
	return
}

func (c *Session) Wait() {
	select {
	case <-c.notifyConnClose:
		logrus.Info("The current connection has been interrupted")
	}
}

func (c *Session) NewChannel(identity string) (err error) {
	channel, err := c.conn.Channel()
	if err != nil {
		return
	}
	c.channel[identity] = channel
	c.notifyChanClose[identity] = make(chan *amqp.Error)
	channel.NotifyClose(c.notifyChanClose[identity])
	return
}
