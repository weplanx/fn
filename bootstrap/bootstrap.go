package bootstrap

import (
	"context"
	"errors"
	"func-api/application/service/excel"
	"func-api/application/service/excel/utils"
	"func-api/application/service/qrcode"
	"func-api/application/service/storage"
	"func-api/application/service/storage/drive"
	"func-api/config"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype/truetype"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
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
	var bs []byte
	bs, err = ioutil.ReadFile("./config/config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return
	}
	return
}

// Initialize database configuration
// If it is another database, replace the driver here
// gorm.Open(mysql.Open(option.Dsn),...)
// reference https://gorm.io/docs/connecting_to_the_database.html
func InitializeDatabase(cfg *config.Config) (db *gorm.DB, err error) {
	option := cfg.Database
	db, err = gorm.Open(mysql.Open(option.Dsn), &gorm.Config{
		Logger: nil,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   option.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	if option.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(option.MaxIdleConns)
	}
	if option.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(option.MaxOpenConns)
	}
	if option.ConnMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(option.ConnMaxLifetime))
	}
	return
}

// Initialize local storage or object storage
// reference config.example.yml
// oss https://help.aliyun.com/document_detail/32144.html?spm=5176.87240.400427.53.55df4614cxcDia
// obs https://support.huaweicloud.com/sdk-go-devg-obs/obs_23_0101.html
// cos https://cloud.tencent.com/document/product/436/31215
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
	case "obs":
		if stg.Drive, err = drive.InitializeObs(drive.ObsOption{
			Endpoint:        option.Option["endpoint"].(string),
			AccessKeyId:     option.Option["access_key_id"].(string),
			AccessKeySecret: option.Option["access_key_secret"].(string),
			BucketName:      option.Option["bucket_name"].(string),
		}); err != nil {
			return
		}
		break
	case "cos":
		if stg.Drive, err = drive.InitializeCos(drive.CosOption{
			Region:     option.Option["region"].(string),
			SecretId:   option.Option["secret_id"].(string),
			SecretKey:  option.Option["secret_key"].(string),
			BucketName: option.Option["bucket_name"].(string),
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

// Initialize QRCode function logic
func InitializeQRCode(cfg *config.Config) (qr *qrcode.Service, err error) {
	qr = new(qrcode.Service)
	qr.Fonts = make(map[string]*truetype.Font)
	for key, val := range cfg.Fonts {
		var bs []byte
		if bs, err = ioutil.ReadFile(val); err != nil {
			return
		}
		if qr.Fonts[key], err = truetype.Parse(bs); err != nil {
			return
		}
	}
	return
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
