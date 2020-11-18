package qrcode

import (
	"func-api/application/model"
	"func-api/application/service/qrcode"
	"github.com/gin-gonic/gin"
	"log"
)

type _PreBuiltBody struct {
	Options []qrcode.Option `json:"options"`
}

func (c *Controller) PreBuilt(ctx *gin.Context) interface{} {
	var body _PreBuiltBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	hashMap := make(map[string]*qrcode.Option)
	keys := make([]string, len(body.Options))
	for index, option := range body.Options {
		hashMap[option.Content] = &option
		keys[index] = option.Content
	}
	var results []map[string]interface{}
	c.Db.Model(&model.Object{}).Where("`key` in ?", keys).Select([]string{"`key`"}).Find(&results)
	existsMap := make(map[string]*qrcode.Option)
	for _, data := range results {
		key := data["key"].(string)
		if hashMap[key] != nil {
			existsMap[key] = hashMap[key]
			delete(hashMap, key)
		}
	}
	log.Println(existsMap)
	log.Println(hashMap)
	return true
}
