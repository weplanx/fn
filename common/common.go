package common

import "go.mongodb.org/mongo-driver/mongo"

type Inject struct {
	Values *Values
	Mongo  *mongo.Client
	Db     *mongo.Database
}

type Values struct {
	// MongoDB 配置
	Database Database `envPrefix:"DATABASE_"`
}

type Database struct {
	Uri string `env:"URI"`
	Db  string `env:"DBNAME"`
}
