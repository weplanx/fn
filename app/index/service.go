package index

import (
	"context"
	"github.com/weplanx/openapi/common"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
)

type Service struct {
	*common.Inject
}

func (x *Service) FindIp(ctx context.Context, value string) (data map[string]interface{}, err error) {
	ip := uint64(0)
	for k, v := range strings.Split(value, ".") {
		n, _ := strconv.ParseUint(v, 10, 64)
		ip |= n << ((3 - uint64(k)) << 3)
	}
	if err = x.Db.Collection("ip").FindOne(ctx,
		bson.M{
			"start": bson.M{"$lte": ip},
			"end":   bson.M{"$gt": ip},
		},
	).Decode(&data); err != nil {
		return
	}
	delete(data, "_id")
	delete(data, "start")
	delete(data, "end")
	return
}
