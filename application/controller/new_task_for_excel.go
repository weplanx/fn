package controller

import (
	"github.com/gin-gonic/gin"
)

type NewTaskForExcelBody struct {
	SheetsName []string `json:"sheets_name"`
}

func (c *controller) NewTaskForExcel(ctx *gin.Context) interface{} {
	var body NewTaskForExcelBody
	var err error
	if err = ctx.ShouldBindJSON(&body); err != nil {
		return c.error(err)
	}
	var taskId string
	if taskId, err = c.dep.Excel.NewTask(body.SheetsName); err != nil {
		return c.error(err)
	}
	return c.result(gin.H{
		"task_id": taskId,
	})
}
