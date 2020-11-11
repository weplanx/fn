package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommitTaskForExcelBody struct {
	TaskId string `json:"task_id"`
}

func (c *controller) CommitTaskForExcel(ctx *gin.Context) interface{} {
	var body CommitTaskForExcelBody
	var err error
	if err = ctx.ShouldBindJSON(&body); err != nil {
		return c.error(err)
	}
	var buf *bytes.Buffer
	if buf, err = c.dep.Excel.Commit(body.TaskId); err != nil {
		return c.error(err)
	}
	filename := uuid.New().String() + ".xlsx"
	if err = c.dep.Storage.Put(filename, buf.Bytes()); err != nil {
		return c.error(err)
	}
	return c.result(gin.H{
		"url": filename,
	})
}
