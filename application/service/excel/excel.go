package excel

import (
	"bytes"
	"errors"
	"funcext/application/service/excel/utils"
	pb "funcext/router"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
	"time"
)

var (
	AppendError = errors.New("an exception occurred in the row writing of Excel")
	CommitError = errors.New("an exception occurred commit of Excel")
)

type Excel struct {
	Task *utils.TaskMap
}

func (c *Excel) NewTask(sheets []string) (taskId string, err error) {
	taskId = uuid.New().String()
	file := excelize.NewFile()
	streamWriterMap := utils.NewStreamWriterMap()
	for _, sheet := range sheets {
		file.NewSheet(sheet)
		var streamWriter *excelize.StreamWriter
		streamWriter, err = file.NewStreamWriter(sheet)
		if err != nil {
			return
		}
		streamWriterMap.Put(sheet, streamWriter)
	}
	go func() {
		timer := time.NewTimer(time.Minute * 30)
		defer timer.Stop()
		select {
		case <-timer.C:
			c.Task.Remove(taskId)
			break
		}
	}()
	c.Task.Put(taskId, &utils.TaskOption{
		File:            file,
		StreamWriterMap: streamWriterMap,
	})
	return
}

func (c *Excel) Append(data *pb.StreamRow) (err error) {
	var task *utils.TaskOption
	var found bool
	if task, found = c.Task.Get(data.TaskId); !found {
		return AppendError
	}
	var streamWriter *excelize.StreamWriter
	if streamWriter, found = task.StreamWriterMap.Get(data.Sheet); !found {
		return AppendError
	}
	if err = streamWriter.SetRow(data.Axis, []interface{}{
		excelize.Cell{Value: data.Value},
	}); err != nil {
		return
	}
	return
}

func (c *Excel) Commit(taskId string) (buf *bytes.Buffer, err error) {
	task, found := c.Task.Get(taskId)
	if !found {
		err = CommitError
		return
	}
	if err = task.StreamWriterMap.Flush(); err != nil {
		return
	}
	if buf, err = task.File.WriteToBuffer(); err != nil {
		return
	}
	err = c.Task.Remove(taskId)
	if err != nil {
		return
	}
	return
}
