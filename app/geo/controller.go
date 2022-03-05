package geo

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Controller struct {
	Service *Service
}

type CountryQuery struct {
	Fields string `form:"fields"`
}

func (x *Controller) Countries(c *gin.Context) interface{} {
	var query CountryQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		return err
	}
	fields := make([]string, 0)
	if query.Fields != "" {
		fields = strings.Split(query.Fields, ",")
	}
	data, err := x.Service.FindCountries(c.Request.Context(), fields)
	if err != nil {
		return err
	}
	return data
}

type StatesQuery struct {
	Country string `form:"country" binding:"required"`
	Fields  string `form:"fields"`
}

func (x *Controller) States(c *gin.Context) interface{} {
	var query StatesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		return err
	}
	fields := make([]string, 0)
	if query.Fields != "" {
		fields = strings.Split(query.Fields, ",")
	}
	data, err := x.Service.FindStates(c.Request.Context(), query.Country, fields)
	if err != nil {
		return err
	}
	return data
}

type CitiesQuery struct {
	Country string `form:"country" binding:"required"`
	State   string `form:"state" binding:"required"`
	Fields  string `form:"fields"`
}

func (x *Controller) Cities(c *gin.Context) interface{} {
	var query CitiesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		return err
	}
	fields := make([]string, 0)
	if query.Fields != "" {
		fields = strings.Split(query.Fields, ",")
	}
	data, err := x.Service.FindCities(c.Request.Context(),
		query.Country,
		query.State,
		fields,
	)
	if err != nil {
		return err
	}
	return data
}
