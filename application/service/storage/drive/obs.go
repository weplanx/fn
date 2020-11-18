package drive

import (
	"bytes"
	"func-api/library/obs"
)

type Obs struct {
	client *obs.ObsClient
	bucket string
	API
}

type ObsOption struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
}

func InitializeObs(option ObsOption) (c *Obs, err error) {
	c = new(Obs)
	if c.client, err = obs.New(
		option.AccessKeyId,
		option.AccessKeySecret,
		option.Endpoint,
	); err != nil {
		return
	}
	c.bucket = option.BucketName
	return
}

func (c *Obs) Put(filename string, body []byte) (err error) {
	input := new(obs.PutObjectInput)
	input.Bucket = c.bucket
	input.Key = filename
	input.Body = bytes.NewReader(body)
	if _, err = c.client.PutObject(input); err != nil {
		return
	}
	return
}
