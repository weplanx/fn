package amqpext

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Session struct {
	url             string
	connected       bool
	conn            *amqp.Connection
	done            chan int
	notifyConnClose chan *amqp.Error
	channel         map[string]*amqp.Channel
	notifyChanClose map[string]chan *amqp.Error
	consumeOptions  map[string]*ConsumeOption
}

func NewSession(url string) (session *Session, err error) {
	session = new(Session)
	session.url = url
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	session.connected = true
	session.conn = conn
	session.done = make(chan int)
	session.notifyConnClose = make(chan *amqp.Error)
	conn.NotifyClose(session.notifyConnClose)
	go session.listenConn()
	session.channel = make(map[string]*amqp.Channel)
	session.notifyChanClose = make(map[string]chan *amqp.Error)
	session.consumeOptions = make(map[string]*ConsumeOption)
	return
}

func (c *Session) listenConn() {
	select {
	case <-c.notifyConnClose:
		logrus.Error("AMQP connection has been disconnected")
		c.reconnected()
	case <-c.done:
		break
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
		go c.listenConn()
		for ID := range c.channel {
			err = c.Channel(ID)
			if err != nil {
				continue
			}
		}
		for _, option := range c.consumeOptions {
			err = c.Consume(*option)
			if err != nil {
				continue
			}
		}
		logrus.Info("Attempt to reconnect successfully")
		break
	}
}

func (c *Session) Channel(ID string) (err error) {
	channel, err := c.conn.Channel()
	if err != nil {
		return
	}
	c.channel[ID] = channel
	c.notifyChanClose[ID] = make(chan *amqp.Error)
	channel.NotifyClose(c.notifyChanClose[ID])
	go c.listenChannel(ID)
	return
}

func (c *Session) listenChannel(ID string) {
	select {
	case <-c.notifyChanClose[ID]:
		logrus.Error("Channel connection is disconnected:", ID)
		c.refreshChannel(ID)
	case <-c.done:
		break
	}
}

func (c *Session) refreshChannel(ID string) {
	for {
		err := c.Channel(ID)
		if err != nil {
			continue
		}
		for _, option := range c.consumeOptions {
			if option.ID == ID {
				err = c.Consume(*option)
				if err != nil {
					continue
				}
			}
		}
		logrus.Info("Channel refresh successfully")
		break
	}
}

func (c *Session) CloseChannel(ID string) error {
	return c.channel[ID].Close()
}

func (c *Session) Consume(option ConsumeOption) (err error) {
	c.consumeOptions[option.Consumer] = &option
	msgs, err := c.channel[option.ID].Consume(
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

func (c *Session) Close() (err error) {
	for _, channel := range c.channel {
		err = channel.Close()
	}
	c.done <- 1
	err = c.conn.Close()
	return
}
