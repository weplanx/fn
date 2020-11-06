package bootstrap

import (
	"errors"
	"funcext/application/service/excel"
	"funcext/application/service/storage"
	"funcext/application/service/storage/drive"
	"funcext/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

var (
	LoadConfigurationNotExists = errors.New("the configuration file does not exist")
	LoadStorageNotExists       = errors.New("please configure for storage settings")
)

// Load application configuration
// reference config.example.yml
func LoadConfiguration() (cfg *config.Config, err error) {
	if _, err = os.Stat("./config.yml"); os.IsNotExist(err) {
		err = LoadConfigurationNotExists
		return
	}
	var buf []byte
	buf, err = ioutil.ReadFile("./config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return
	}
	return
}

func InitializeStorage(cfg *config.Config) (stg *storage.Storage, err error) {
	option := cfg.Storage
	if reflect.DeepEqual(option, storage.Option{}) {
		err = LoadStorageNotExists
		return
	}
	stg = new(storage.Storage)
	switch option.Drive {
	case "local":
		stg.Drive = drive.InitializeLocal(option.Option["path"].(string))
		break
	}
	return
}

func InitializeExcel() (ex *excel.Excel, err error) {
	ex = new(excel.Excel)
	return
}
