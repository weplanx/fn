package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"time"
)

type TaskOption struct {
	File      *excelize.File
	StreamMap *streamMap
}

type TaskMap struct {
	hashMap     map[string]*TaskOption
	termination map[string]chan int
}

func NewTaskMap() *TaskMap {
	c := new(TaskMap)
	c.hashMap = make(map[string]*TaskOption)
	c.termination = make(map[string]chan int)
	return c
}

func (c *TaskMap) Put(taskId string, option *TaskOption) {
	c.hashMap[taskId] = option
	c.termination[taskId] = make(chan int)
	go func() {
		timer := time.NewTimer(time.Minute * 30)
		defer timer.Stop()
		select {
		case <-timer.C:
			if task, found := c.Get(taskId); found {
				task.StreamMap.Clear()
				c.Remove(taskId)
			}
			break
		case <-c.termination[taskId]:
			break
		}
	}()
}

func (c *TaskMap) Get(taskId string) (option *TaskOption, found bool) {
	found = c.hashMap[taskId] != nil
	if found {
		option = c.hashMap[taskId]
	}
	return
}

func (c *TaskMap) Termination(taskId string) {
	c.termination[taskId] <- 1
}

func (c *TaskMap) Remove(taskId string) {
	delete(c.hashMap, taskId)
	return
}
