package excel

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"runtime/debug"
)

type CommitTaskForExcelBody struct {
	TaskId string `json:"task_id"`
}

func (c *Controller) CommitTask(ctx *gin.Context) interface{} {
	var body CommitTaskForExcelBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var buf *bytes.Buffer
	if buf, err = c.Excel.Commit(body.TaskId); err != nil {
		return err
	}
	filename := uuid.New().String() + ".xlsx"
	if err = c.Storage.Put(filename, buf.Bytes()); err != nil {
		return err
	}
	buf.Reset()
	debug.FreeOSMemory()
	return gin.H{
		"url": filename,
	}
}
