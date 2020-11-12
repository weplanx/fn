package controller

import (
	"bytes"
	"func-api/application/common/typ"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sync"
)

type SimpleExcelBody struct {
	Sheets []typ.Sheet `json:"sheets"`
}

func (c *controller) SimpleExcel(ctx *gin.Context) interface{} {
	var body SimpleExcelBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	file := excelize.NewFile()
	var wg sync.WaitGroup
	wg.Add(len(body.Sheets))
	for _, sheet := range body.Sheets {
		go func(sheet typ.Sheet) {
			defer wg.Done()
			var streamWriter *excelize.StreamWriter
			if streamWriter, err = file.NewStreamWriter(sheet.Name); err != nil {
				return
			}
			for _, row := range sheet.Rows {
				if err = streamWriter.SetRow(row.Axis, []interface{}{
					excelize.Cell{Value: row.Value},
				}); err != nil {
					return
				}
			}
			if err = streamWriter.Flush(); err != nil {
				return
			}
		}(sheet)
	}
	wg.Wait()
	var buf *bytes.Buffer
	if buf, err = file.WriteToBuffer(); err != nil {
		return err
	}
	filename := uuid.New().String() + ".xlsx"
	if err = c.dep.Storage.Put(filename, buf.Bytes()); err != nil {
		return err
	}
	return gin.H{
		"url": filename,
	}
}
