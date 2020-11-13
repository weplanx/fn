package controller

import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

type SqlExcelBody struct {
	SheetName string                  `json:"sheet_name"`
	Mapping   map[string]FieldMapping `json:"mapping"`
	Raw       string                  `json:"raw"`
}

type FieldMapping struct {
	Name string `json:"name"`
	Col  string `json:"col"`
}

func (c *Controller) SqlExcel(ctx *gin.Context) interface{} {
	var body SqlExcelBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var results []map[string]interface{}
	c.Db.Raw(body.Raw).Scan(&results)
	file := excelize.NewFile()
	var streamWriter *excelize.StreamWriter
	if streamWriter, err = file.NewStreamWriter(body.SheetName); err != nil {
		return err
	}
	for _, mapping := range body.Mapping {
		if err = streamWriter.SetRow(mapping.Col+"1", []interface{}{
			excelize.Cell{Value: mapping.Name},
		}); err != nil {
			return err
		}
	}
	for i, data := range results {
		index := strconv.Itoa(i + 2)
		for key, value := range data {
			streamWriter.SetRow(body.Mapping[key].Col+index, []interface{}{
				excelize.Cell{Value: value},
			})
		}
	}
	if err = streamWriter.Flush(); err != nil {
		return err
	}
	var buf *bytes.Buffer
	if buf, err = file.WriteToBuffer(); err != nil {
		return err
	}
	filename := uuid.New().String() + ".xlsx"
	if err = c.Storage.Put(filename, buf.Bytes()); err != nil {
		return err
	}
	return gin.H{
		"url": filename,
	}
}
