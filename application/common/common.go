package common

import (
	"funcext/application/service/excel"
	"funcext/application/service/storage"
	"funcext/config"
	"go.uber.org/fx"
)

type Dependency struct {
	fx.In

	Config  *config.Config
	Storage *storage.Storage
	Excel   *excel.Excel
}
