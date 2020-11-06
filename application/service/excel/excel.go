package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Excel struct {
}

type CreateOption struct {
}

func (c *Excel) Create(option CreateOption) (file *excelize.File, err error) {
	file = excelize.NewFile()
	return
}
