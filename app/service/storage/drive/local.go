package drive

import "github.com/spf13/afero"

type Local struct {
	fs *afero.Fs
	API
}

func InitializeLocal(path string) *Local {
	c := new(Local)
	AppFs := afero.NewBasePathFs(afero.NewOsFs(), path)
	c.fs = &AppFs
	return c
}

func (c *Local) Put(filename string, body []byte) (err error) {
	var file afero.File
	file, err = (*c.fs).Create(filename)
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
