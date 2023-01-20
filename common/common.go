package common

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Inject struct {
	Values *Values
	Mongo  *mongo.Client
	Db     *mongo.Database
}

type Values struct {
	// 监听地址
	Address string `env:"address" envDefault:":9000"`
	// 数据库
	Database Database `envPrefix:"DATABASE_"`
	// 存储
	Storage Storage `envPrefix:"STORAGE_"`
}

type Database struct {
	Uri    string `env:"URI"`
	DbName string `env:"DBNAME"`
}

type Storage struct {
	Type string `env:"TYPE"`
	Url  string `env:"URL"`
	Id   string `env:"ID"`
	Key  string `env:"KEY"`
}
