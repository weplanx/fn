package geo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strings"
)

type Controller struct {
	Service *Service
}

type CountriesDto struct {
	Fields string `query:"fields"`
}

func (x *Controller) Countries(ctx context.Context, c *app.RequestContext) {
	var dto CountriesDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}
	fields := make([]string, 0)
	if dto.Fields != "" {
		fields = strings.Split(dto.Fields, ",")
	}

	data, err := x.Service.FindCountries(ctx, fields)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

type StatesDto struct {
	Country string `query:"country,required"`
	Fields  string `query:"fields"`
}

func (x *Controller) States(ctx context.Context, c *app.RequestContext) {
	var dto StatesDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}
	fields := make([]string, 0)
	if dto.Fields != "" {
		fields = strings.Split(dto.Fields, ",")
	}
	data, err := x.Service.FindStates(ctx, dto.Country, fields)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

type CitiesDto struct {
	Country string `query:"country,required"`
	State   string `query:"state,required"`
	Fields  string `query:"fields"`
}

func (x *Controller) Cities(ctx context.Context, c *app.RequestContext) {
	var dto CitiesDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}
	fields := make([]string, 0)
	if dto.Fields != "" {
		fields = strings.Split(dto.Fields, ",")
	}
	data, err := x.Service.FindCities(ctx, dto.Country, dto.State, fields)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}
