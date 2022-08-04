package index

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"time"
)

type Controller struct {
	IndexService *Service
}

func (x *Controller) Index(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, utils.H{
		"msg":  "hi",
		"ip":   c.ClientIP(),
		"time": time.Now(),
	})
}

type IpDto struct {
	Value string `query:"value,required"`
}

func (x *Controller) GetIp(ctx context.Context, c *app.RequestContext) {
	var dto IpDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	data, err := x.IndexService.GetIp(ctx, dto.Value)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}
