package common

import (
	"func-api/application/service/excel"
	"func-api/application/service/storage"
	"func-api/config"
	"go.uber.org/fx"
)

type Dependency struct {
	fx.In

	Config  *config.Config
	Storage *storage.Storage
	Excel   *excel.Excel
}

func Inject(i interface{}) *Dependency {
	return i.(*Dependency)
}
