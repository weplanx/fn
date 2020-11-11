package controller

import (
	"func-api/application/service/excel"
	"github.com/gin-gonic/gin"
)

func (c *controller) AddRowToExcel(ctx *gin.Context) interface{} {
	var body excel.ChunkData
	var err error
	if err = ctx.ShouldBindJSON(&body); err != nil {
		return c.error(err)
	}
	if err = c.dep.Excel.Append(body); err != nil {
		return c.error(err)
	}
	return c.ok()
}
