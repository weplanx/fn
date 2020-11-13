package controller

import (
	"func-api/application/common/typ"
	"github.com/gin-gonic/gin"
)

func (c *Controller) AddRowToExcel(ctx *gin.Context) interface{} {
	var body typ.ChunkData
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	if err = c.Excel.Append(body); err != nil {
		return err
	}
	return true
}
