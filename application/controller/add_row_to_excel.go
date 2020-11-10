package controller

import (
	pb "funcext/router"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
)

func (c *controller) AddRowToExcel(stream pb.Router_AddRowToExcelServer) (err error) {
	for {
		var data *pb.StreamRow
		data, err = stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&empty.Empty{})
		}
		if err != nil {
			return
		}
		err = c.dep.Excel.Append(data)
		if err != nil {
			return
		}
	}
}
