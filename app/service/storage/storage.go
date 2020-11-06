package storage

import "funcext/app/service/storage/drive"

type Storage struct {
	Drive interface{}
	drive.API
}

func (c *Storage) Put(filename string, body []byte) (err error) {
	return c.Drive.(drive.API).Put(filename, body)
}
