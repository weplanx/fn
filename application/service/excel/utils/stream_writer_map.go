package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type streamWriterMap struct {
	hashMap map[string]*excelize.StreamWriter
}

func NewStreamWriterMap() *streamWriterMap {
	c := new(streamWriterMap)
	c.hashMap = make(map[string]*excelize.StreamWriter)
	return c
}

func (c *streamWriterMap) Put(sheet string, streamWriter *excelize.StreamWriter) {
	c.hashMap[sheet] = streamWriter
}

func (c *streamWriterMap) Get(sheet string) (streamWriter *excelize.StreamWriter, found bool) {
	found = c.hashMap[sheet] != nil
	if found {
		streamWriter = c.hashMap[sheet]
	}
	return
}

func (c *streamWriterMap) Flush() (err error) {
	for _, streamWriter := range c.hashMap {
		err = streamWriter.Flush()
		if err != nil {
			return
		}
	}
	return
}

func (c *streamWriterMap) Remove(sheet string) (err error) {
	delete(c.hashMap, sheet)
	return
}

func (c *streamWriterMap) Clear() (err error) {
	for sheet, _ := range c.hashMap {
		delete(c.hashMap, sheet)
	}
	return
}
