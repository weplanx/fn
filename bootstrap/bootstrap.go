package bootstrap

import (
	"github.com/caarlos0/env/v6"
	"github.com/google/wire"
	"github.com/weplanx/openapi/common"
)

var Provides = wire.NewSet()

// SetValues 初始化配置
func SetValues() (values *common.Values, err error) {
	values = new(common.Values)
	if err = env.Parse(values); err != nil {
		return
	}
	return
}
