package storage

import (
	"func-api/application/service/storage/drive"
)

type Service struct {
	Drive interface{}
	drive.API
}

type Option struct {
	Drive  string                 `yaml:"drive"`
	Option map[string]interface{} `yaml:"option"`
}

func (c *Service) Put(filename string, body []byte) (err error) {
	return c.Drive.(drive.API).Put(filename, body)
}
