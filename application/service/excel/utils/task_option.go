package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type TaskOption struct {
	File            *excelize.File
	StreamWriterMap *streamWriterMap
}
