package qrcode

import (
	"errors"
	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
	"image"
)

var (
	FontNotExists = errors.New("the current font does not exist")
)

type Service struct {
	Fonts map[string]string
}

type Option struct {
	Content string `json:"content"`
	Size    int    `json:"size"`
	Font    Font   `json:"font"`
}

type Font struct {
	Text   string  `json:"text"`
	Type   string  `json:"type"`
	Size   float64 `json:"size"`
	Margin float64 `json:"margin"`
}

func (c *Service) Factory(option Option) (im image.Image, err error) {
	var qr *qrcode.QRCode
	if qr, err = qrcode.New(option.Content, qrcode.Medium); err != nil {
		return
	}
	dc := gg.NewContextForImage(qr.Image(option.Size))
	if option.Font != (Font{}) {
		font := option.Font
		if typ, ok := c.Fonts[font.Type]; ok {
			if err = dc.LoadFontFace(typ, font.Size); err != nil {
				return
			}
		} else {
			err = FontNotExists
			return
		}
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(
			font.Text,
			float64(option.Size)/2,
			float64(option.Size)-font.Margin,
			0.5,
			0.5,
		)
	}
	im = dc.Image()
	return
}
