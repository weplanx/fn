package qrcode

import (
	"bytes"
	"context"
	"func-api/application/model"
	"func-api/application/service/qrcode"
	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
)

type _PreBuiltBody struct {
	Options []qrcode.Option `json:"options"`
	Update  bool            `json:"update"`
}

func (c *Controller) PreBuilt(ctx *gin.Context) interface{} {
	var body _PreBuiltBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	hashMap := make(map[string]qrcode.Option)
	keys := make([]string, len(body.Options))
	for index, option := range body.Options {
		hashMap[option.Content] = option
		keys[index] = option.Content
	}
	var results []map[string]interface{}
	c.Db.Model(&model.Object{}).Where("`key` in ?", keys).Select([]string{"`key`"}).Find(&results)
	existsMap := make(map[string]qrcode.Option)
	for _, data := range results {
		key := data["key"].(string)
		if hashMap[key] != (qrcode.Option{}) {
			if body.Update {
				existsMap[key] = hashMap[key]
			}
			delete(hashMap, key)
		}
	}
	tx := c.Db.WithContext(context.Background())
	for _, option := range existsMap {
		go func(option qrcode.Option) {
			var dc *gg.Context
			if dc, err = c.QRCode.Factory(option); err != nil {
				return
			}
			var buf bytes.Buffer
			if err = dc.EncodePNG(&buf); err != nil {
				return
			}
			tx.Model(&model.Object{}).Where("`key` = ?", option.Content).Update("value", buf.Bytes())
		}(option)
	}
	data := make([]model.Object, len(hashMap))
	index := 0
	for _, option := range hashMap {
		var dc *gg.Context
		if dc, err = c.QRCode.Factory(option); err != nil {
			return err
		}
		var buf bytes.Buffer
		if err = dc.EncodePNG(&buf); err != nil {
			return err
		}
		data[index] = model.Object{
			Key:   option.Content,
			Value: buf.Bytes(),
		}
		index++
	}
	if len(data) != 0 {
		go c.Db.Create(data)
	}
	return true
}
