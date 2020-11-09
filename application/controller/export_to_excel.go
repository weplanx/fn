package controller

import (
	"bytes"
	"context"
	pb "funcext/router"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
	"sync"
)

func (c *controller) ExportToExcel(_ context.Context, param *pb.ExportToExcelParameter) (
	response *pb.ExportToExcelResponse, err error,
) {
	file := excelize.NewFile()
	var wg sync.WaitGroup
	wg.Add(len(param.Sheets))
	for _, sheet := range param.Sheets {
		go func(sheet *pb.Sheet) {
			defer wg.Done()
			var err error
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
	filename := uuid.New().String() + ".xlsx"
	var buf *bytes.Buffer
	if buf, err = file.WriteToBuffer(); err != nil {
		return
	}
	if err = c.dep.Storage.Put(filename, buf.Bytes()); err != nil {
		return
	}
	response = &pb.ExportToExcelResponse{
		Url: filename,
	}
	return
}
