package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type streamMap struct {
	hashMap map[string]*excelize.StreamWriter
}

func NewStreamWriterMap() *streamMap {
	c := new(streamMap)
	c.hashMap = make(map[string]*excelize.StreamWriter)
	return c
}

func (c *streamMap) Put(sheetName string, streamWriter *excelize.StreamWriter) {
	c.hashMap[sheetName] = streamWriter
}

func (c *streamMap) Get(sheetName string) (streamWriter *excelize.StreamWriter, found bool) {
	found = c.hashMap[sheetName] != nil
	if found {
		streamWriter = c.hashMap[sheetName]
	}
	return
}

func (c *streamMap) Flush() (err error) {
	for _, streamWriter := range c.hashMap {
		err = streamWriter.Flush()
		if err != nil {
			return
		}
	}
	return
}

func (c *streamMap) Remove(sheetName string) (err error) {
	delete(c.hashMap, sheetName)
	return
}
