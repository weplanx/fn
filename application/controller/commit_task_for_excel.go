package controller

import (
	"bytes"
	"context"
	pb "funcext/router"
	"github.com/google/uuid"
)

func (c *controller) CommitTaskForExcel(_ context.Context, param *pb.Task) (result *pb.ExportURL, err error) {
	var buf *bytes.Buffer
	buf, err = c.dep.Excel.Commit(param.TaskId)
	if err != nil {
		return
	}
	filename := uuid.New().String() + ".xlsx"
	if err = c.dep.Storage.Put(filename, buf.Bytes()); err != nil {
		return
	}
	result = &pb.ExportURL{Url: filename}
	return
}
