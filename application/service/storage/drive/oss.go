package drive

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Oss struct {
	client *oss.Client
	bucket *oss.Bucket
	API
}

type OssOption struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
}

func InitializeOss(option OssOption) (c *Oss, err error) {
	c = new(Oss)
	if c.client, err = oss.New(
		option.Endpoint,
		option.AccessKeyId,
		option.AccessKeySecret,
	); err != nil {
		return
	}
	if c.bucket, err = c.client.Bucket(
		option.BucketName,
	); err != nil {
		return
	}
	return
}

func (c *Oss) Put(filename string, body []byte) (err error) {
	return c.bucket.PutObject(filename, bytes.NewReader(body))
}
