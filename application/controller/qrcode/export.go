package qrcode

import (
	"bytes"
	"func-api/application/model"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"runtime/debug"
	"strconv"
	"time"
)

type _ExportBody struct {
	Lists []string `json:"lists"`
}

func (c *Controller) Export(ctx *gin.Context) interface{} {
	var body _ExportBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	file := excelize.NewFile()
	keys := make([]string, len(body.Lists))
	for index, content := range body.Lists {
		keys[index] = content
	}
	var objects []model.Object
	c.Db.Where("`key` in ?", keys).Find(&objects)
	for i, object := range objects {
		cell := "A" + strconv.Itoa(i+1)
		if err = file.SetRowHeight("Sheet1", i+1, 128); err != nil {
			return err
		}
		if err = file.AddPictureFromBytes(
			"Sheet1", cell, "", "", ".png", object.Value,
		); err != nil {
			return err
		}
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
	buf.Reset()
	debug.FreeOSMemory()
	return gin.H{
		"url": filename,
	}
}
