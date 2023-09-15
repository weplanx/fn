//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/weplanx/fn/api"
	"github.com/weplanx/fn/common"
)

func NewAPI() (*api.API, error) {
	wire.Build(
		wire.Struct(new(api.API), "*"),
		wire.Struct(new(common.Inject), "*"),
		LoadStaticValues,
		UseCos,
	)
	return &api.API{}, nil
}
