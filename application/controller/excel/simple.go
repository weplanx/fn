package excel

import (
	"bytes"
	"func-api/application/service/excel/typ"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type _SimpleBody struct {
	Sheets []typ.Sheet `json:"sheets"`
}

func (c *Controller) Simple(ctx *gin.Context) interface{} {
	var body _SimpleBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var file *excelize.File
	if file, err = c.Excel.Simple(body.Sheets); err != nil {
		return err
	}
	var buf *bytes.Buffer
	if buf, err = file.WriteToBuffer(); err != nil {
		return err
	}
	date := time.Now().Format("2006-01-02")
	filename := date + "/" + uuid.New().String() + ".xlsx"
	if err = c.Storage.Put(filename, buf.Bytes()); err != nil {
		return err
	}
	return gin.H{
		"url": filename,
	}
}
