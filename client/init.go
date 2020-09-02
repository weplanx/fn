package client

import (
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
	return
}

func (c *Session) listen() {
	select {
	case <-c.notifyConnClose:
		logrus.Error("AMQP connection has been disconnected")
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
			logrus.Info("Attempt to reconnect successfully")
			break
		}
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
