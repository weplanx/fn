package qrcode

import (
	"errors"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/skip2/go-qrcode"
)

var (
	FontNotExists = errors.New("the current font does not exist")
)

type Service struct {
	Fonts map[string]*truetype.Font
}

type Option struct {
	Content string `json:"content"`
	Size    int    `json:"size"`
	Text    Text   `json:"text"`
}

type Text struct {
	Value  string  `json:"value"`
	Type   string  `json:"type"`
	Size   float64 `json:"size"`
	Margin float64 `json:"margin"`
}

func (c *Service) Factory(option Option) (dc *gg.Context, err error) {
	var qr *qrcode.QRCode
	if qr, err = qrcode.New(option.Content, qrcode.Medium); err != nil {
		return
	}
	dc = gg.NewContextForImage(qr.Image(option.Size))
	if option.Text != (Text{}) {
		text := option.Text
		if c.Fonts[text.Type] == nil {
			err = FontNotExists
			return
		}
		dc.SetFontFace(truetype.NewFace(
			c.Fonts[text.Type],
			&truetype.Options{
				Size: text.Size,
			},
		))
		dc.SetRGB(0, 0, 0)
		if text.Value == "" {
			text.Value = option.Content
		}
		dc.DrawStringAnchored(
			text.Value,
			float64(option.Size)/2,
			float64(option.Size)-text.Margin,
			0.5,
			0.5,
		)
	}
	return
}
