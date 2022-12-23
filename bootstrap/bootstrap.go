package bootstrap

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/weplanx/openapi/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// LoadValues 加载配置
func LoadValues() (values *common.Values, err error) {
	values = new(common.Values)
	if err = env.Parse(values); err != nil {
		return
	}
	return
}

// UseMongoDB 初始化 MongoDB
// 配置文档 https://www.mongodb.com/docs/drivers/go/current/
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
func UseMongoDB(values *common.Values) (*mongo.Client, error) {
	return mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(values.Database.Uri),
	)
}

// UseDatabase 初始化数据库
// 配置文档 https://www.mongodb.com/docs/drivers/go/current/
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
func UseDatabase(client *mongo.Client, values *common.Values) (db *mongo.Database) {
	option := options.Database().
		SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	return client.Database(values.Database.DbName, option)
}

// UseHertz 使用 Hertz
// 配置文档 https://www.cloudwego.io/zh/docs/hertz/reference/config
func UseHertz(values *common.Values) (h *server.Hertz, err error) {
	h = server.Default(
		server.WithHostPorts(values.Address),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
		server.WithExitWaitTime(0),
	)
	return
}
