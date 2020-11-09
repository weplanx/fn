package controller

import (
	"bytes"
	pb "funcext/router"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
	"io"
	"sync"
)

func (c *controller) ExportToExcel(stream pb.Router_ExportToExcelServer) (err error) {
	file := excelize.NewFile()
	filename := uuid.New().String() + ".xlsx"
	for {
		var param *pb.ExportToExcelParameter
		param, err = stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ExportToExcelResponse{
				Url: filename,
			})
		}
		if err != nil {
			return
		}
		var wg sync.WaitGroup
		wg.Add(len(param.Sheets))
		for _, sheet := range param.Sheets {
			go func(sheet *pb.Sheet) {
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
			return
		}
		if err = c.dep.Storage.Put(filename, buf.Bytes()); err != nil {
			return
		}
	}
}
