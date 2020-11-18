package qrcode

import (
	"bytes"
	"encoding/base64"
	"func-api/application/service/qrcode"
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
)

func (c *Controller) Testing(ctx *gin.Context) interface{} {
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
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return b64
}
