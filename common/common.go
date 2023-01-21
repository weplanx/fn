package common

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Inject struct {
	Values    *Values
	Mongo     *mongo.Client
	Db        *mongo.Database
	Nats      *nats.Conn
	JetStream nats.JetStreamContext
}

type Values struct {
	Address  string `env:"address" envDefault:":9000"`
	Database `envPrefix:"DATABASE_"`
	Nats     `envPrefix:"NATS_"`
	Storage  `envPrefix:"STORAGE_"`
}

type Database struct {
	Mongo string `env:"MONGO,required"`
	Name  string `env:"NAME,required"`
	//Redis string `env:"REDIS,required"`
}

type Nats struct {
	Hosts []string `env:"HOSTS,required" envSeparator:","`
	Nkey  string   `env:"NKEY,required"`
}

type Storage struct {
	Type string `env:"TYPE"`
	Url  string `env:"URL"`
	Id   string `env:"ID"`
	Key  string `env:"KEY"`
}
