package bootstrap

import (
	"github.com/caarlos0/env/v9"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"tencent-workflow-excel/common"
)

func LoadValues() (values *common.Values, err error) {
	values = new(common.Values)
	if err = env.Parse(values); err != nil {
		return
	}
	return
}

func UseCos(values *common.Values) (client *cos.Client, err error) {
	u, _ := url.Parse(values.Cos.Url)
	client = cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  values.Cos.SecretId,
			SecretKey: values.Cos.SecretKey,
		},
	})
	return
}
