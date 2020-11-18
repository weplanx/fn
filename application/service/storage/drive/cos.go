package drive

import (
	"bytes"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

type Cos struct {
	client *cos.Client
	API
}

type CosOption struct {
	Region     string `yaml:"region"`
	SecretId   string `yaml:"secret_id"`
	SecretKey  string `yaml:"secret_key"`
	BucketName string `yaml:"bucket_name"`
}

func InitializeCos(option CosOption) (c *Cos, err error) {
	c = new(Cos)
	var u *url.URL
	if u, err = url.Parse(
		"https://" + option.BucketName + ".cos." + option.Region + ".myqcloud.com",
	); err != nil {
		return
	}
	c.client = cos.NewClient(
		&cos.BaseURL{BucketURL: u},
		&http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  option.SecretId,
				SecretKey: option.SecretKey,
			},
		},
	)
	return
}

func (c *Cos) Put(filename string, body []byte) (err error) {
	if _, err = c.client.Object.Put(
		context.Background(),
		filename,
		bytes.NewReader(body),
		&cos.ObjectPutOptions{},
	); err != nil {
		return
	}
	return
}
