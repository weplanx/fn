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
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var buf *bytes.Buffer
	if buf, err = c.dep.Excel.Commit(body.TaskId); err != nil {
		return err
	}
	filename := uuid.New().String() + ".xlsx"
	if err = c.dep.Storage.Put(filename, buf.Bytes()); err != nil {
		return err
	}
	buf.Reset()
	return gin.H{
		"url": filename,
	}
}
