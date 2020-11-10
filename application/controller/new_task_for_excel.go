package controller

import (
	"context"
	pb "funcext/router"
)

func (c *controller) NewTaskForExcel(_ context.Context, param *pb.NewExcel) (*pb.Task, error) {
	taskId, err := c.dep.Excel.NewTask(param.Sheets)
	if err != nil {
		return nil, err
	}
	return &pb.Task{TaskId: taskId}, nil
}
