package qrcode

import (
	"bytes"
	"encoding/base64"
	"func-api/application/service/qrcode"
	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Testing(ctx *gin.Context) interface{} {
	var body qrcode.Option
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var dc *gg.Context
	if dc, err = c.QRCode.Factory(body); err != nil {
		return err
	}
	var buf bytes.Buffer
	if err = dc.EncodePNG(&buf); err != nil {
		return err
	}
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return b64
}
