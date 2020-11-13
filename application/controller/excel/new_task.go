package excel

import (
	"github.com/gin-gonic/gin"
)

type _NewTaskBody struct {
	SheetsDef []string `json:"sheets_name"`
}

func (c *Controller) NewTask(ctx *gin.Context) interface{} {
	var body _NewTaskBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var taskId string
	if taskId, err = c.Excel.NewTask(body.SheetsDef); err != nil {
		return err
	}
	return gin.H{
		"task_id": taskId,
	}
}
