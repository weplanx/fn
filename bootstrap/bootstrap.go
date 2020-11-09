package bootstrap

import (
	"errors"
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
	case "oss":
		if stg.Drive, err = drive.InitializeOss(drive.OssOption{
			Endpoint:        option.Option["endpoint"].(string),
			AccessKeyId:     option.Option["access_key_id"].(string),
			AccessKeySecret: option.Option["access_key_secret"].(string),
			BucketName:      option.Option["bucket_name"].(string),
		}); err != nil {
			return
		}
		break
	}
	return
}
