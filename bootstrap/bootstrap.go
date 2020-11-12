package bootstrap

import (
	"context"
	"errors"
	"func-api/application/service/excel"
	"func-api/application/service/excel/utils"
	"func-api/application/service/storage"
	"func-api/application/service/storage/drive"
	"func-api/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
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

// Initialize local storage or object storage
// reference config.example.yml
func InitializeStorage(cfg *config.Config) (stg *storage.Service, err error) {
	option := cfg.Storage
	if reflect.DeepEqual(option, storage.Option{}) {
		err = LoadStorageNotExists
		return
	}
	stg = new(storage.Service)
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

// Initialize Excel function logic
func InitializeExcel() *excel.Service {
	ex := new(excel.Service)
	ex.Task = utils.NewTaskMap()
	return ex
}

// Start http service
// https://gin-gonic.com/docs/examples/custom-http-config/
func HttpServer(lc fx.Lifecycle, cfg *config.Config) (serve *gin.Engine) {
	if cfg.Debug != "" {
		go http.ListenAndServe(cfg.Debug, nil)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	serve = gin.New()
	serve.Use(gin.Logger())
	serve.Use(gin.Recovery())
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go serve.Run(cfg.Listen)
			return nil
		},
	})
	return
}
