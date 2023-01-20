package excel

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type Controller struct {
	Service *Service
}

type CreateDto struct {
	File   string  `json:"file"`
	Sheets []Sheet `json:"sheets"`
}

type Sheet struct {
	Name string          `json:"name"`
	Data [][]interface{} `json:"data"`
}

func (x *Controller) Create(ctx context.Context, c *app.RequestContext) {
	var dto CreateDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.Service.Create(ctx, &dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, utils.H{
		"file": dto.File,
	})
}
