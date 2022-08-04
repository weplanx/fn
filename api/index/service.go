package index

import (
	"context"
	"github.com/weplanx/openapi/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

type Service struct {
	*common.Inject
}

func (x *Service) GetIp(ctx context.Context, value string) (data map[string]interface{}, err error) {
	ip := uint64(0)
	for k, v := range strings.Split(value, ".") {
		n, _ := strconv.ParseUint(v, 10, 64)
		ip |= n << ((3 - uint64(k)) << 3)
	}

	option := options.FindOne().
		SetProjection(bson.M{"_id": 0, "range": 0})
	if err = x.Db.Collection("ip").FindOne(ctx,
		bson.M{
			"range": bson.M{"$gt": ip, "$lte": ip},
		},
		option,
	).Decode(&data); err != nil {
		return
	}

	return
}
