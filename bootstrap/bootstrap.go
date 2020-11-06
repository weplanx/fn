package bootstrap

import (
	"errors"
	"funcext/app/service/storage"
	"funcext/app/service/storage/drive"
	"funcext/app/types"
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
func LoadConfiguration() (cfg *types.Config, err error) {
	if _, err = os.Stat("./config/config.yml"); os.IsNotExist(err) {
		err = LoadConfigurationNotExists
		return
	}
	var buf []byte
	buf, err = ioutil.ReadFile("./config/config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return
	}
	return
}

func LoadStorage(cfg *types.Config) (stg *storage.Storage, err error) {
	option := cfg.Storage
	if reflect.DeepEqual(option, types.StorageOption{}) {
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
