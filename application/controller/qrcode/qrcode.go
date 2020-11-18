package qrcode

import (
	"bytes"
	"func-api/application/model"
	"func-api/application/service/qrcode"
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
)

func (c *Controller) Factory(ctx *gin.Context) interface{} {
	var body qrcode.Option
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var im image.Image
	if im, err = c.QRCode.Factory(body); err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = png.Encode(buf, im); err != nil {
		return err
	}
	var count int64
	c.Db.Model(&model.Object{}).Where("`key` = ?", body.Content).Count(&count)
	if count == 0 {
		if err = c.Db.Create(&model.Object{
			Key:   body.Content,
			Value: buf.Bytes(),
		}).Error; err != nil {
			return err
		}
	} else {
		if err = c.Db.Model(&model.Object{}).
			Where("`key` = ?", body.Content).
			Update("value", buf.Bytes()).
			Error; err != nil {
			return err
		}
	}
	return true
}
