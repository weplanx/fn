package drive

import (
	"github.com/spf13/afero"
	"path"
)

type Local struct {
	fs afero.Fs
	API
}

func InitializeLocal(path string) *Local {
	c := new(Local)
	AppFs := afero.NewBasePathFs(afero.NewOsFs(), path)
	c.fs = AppFs
	return c
}

func (c *Local) Put(filename string, body []byte) (err error) {
	var file afero.File
	dir, _ := path.Split(filename)
	if err = c.fs.MkdirAll(dir, 0755); err != nil {
		return
	}
	file, err = c.fs.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		return
	}
	return
}
