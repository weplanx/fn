package storage

import "funcext/application/service/storage/drive"

type Storage struct {
	Drive interface{}
	drive.API
}

type Option struct {
	Drive  string                 `yaml:"drive"`
	Option map[string]interface{} `yaml:"option"`
}

func (c *Storage) Put(filename string, body []byte) (err error) {
	return c.Drive.(drive.API).Put(filename, body)
}
