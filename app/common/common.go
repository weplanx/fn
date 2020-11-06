package common

import (
	"funcext/app/service/storage"
	"funcext/app/types"
	"go.uber.org/fx"
)

type Dependency struct {
	fx.In

	Config  *types.Config
	Storage *storage.Storage
}
