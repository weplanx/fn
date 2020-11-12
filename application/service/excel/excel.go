package excel

import (
	"bytes"
	"errors"
	"func-api/application/common/typ"
	"func-api/application/service/excel/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
	"time"
)

var (
	AppendError = errors.New("an exception occurred in the row writing of Excel")
	CommitError = errors.New("an exception occurred commit of Excel")
)

type Service struct {
	Task *utils.TaskMap
}

func (c *Service) NewTask(sheetsDef []string) (taskId string, err error) {
	taskId = uuid.New().String()
	file := excelize.NewFile()
	streamWriterMap := utils.NewStreamWriterMap()
	for _, sheetName := range sheetsDef {
		file.NewSheet(sheetName)
		var streamWriter *excelize.StreamWriter
		streamWriter, err = file.NewStreamWriter(sheetName)
		if err != nil {
			return
		}
		streamWriterMap.Put(sheetName, streamWriter)
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
		File:      file,
		StreamMap: streamWriterMap,
	})
	return
}

func (c *Service) Append(data typ.ChunkData) (err error) {
	var task *utils.TaskOption
	var found bool
	if task, found = c.Task.Get(data.TaskId); !found {
		return AppendError
	}
	var streamWriter *excelize.StreamWriter
	if streamWriter, found = task.StreamMap.Get(data.SheetName); !found {
		return AppendError
	}
	task.File.Lock()
	for _, row := range data.Rows {
		if err = streamWriter.SetRow(row.Axis, []interface{}{
			excelize.Cell{Value: row.Value},
		}); err != nil {
			return
		}
	}
	task.File.Unlock()
	return
}

func (c *Service) Commit(taskId string) (buf *bytes.Buffer, err error) {
	task, found := c.Task.Get(taskId)
	if !found {
		err = CommitError
		return
	}
	if err = task.StreamMap.Flush(); err != nil {
		return
	}
	if buf, err = task.File.WriteToBuffer(); err != nil {
		return
	}
	c.Task.Remove(taskId)
	return
}
