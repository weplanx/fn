package excel

import (
	"bytes"
	"context"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/weplanx/openapi/common"
	"github.com/xuri/excelize/v2"
	"net/http"
	"net/url"
)

type Service struct {
	*common.Inject
}

func (x *Service) Create(ctx context.Context, dto *CreateDto) (err error) {

	file := excelize.NewFile()
	defer file.Close()
	for _, sheet := range dto.Sheets {
		var streamWriter *excelize.StreamWriter
		if streamWriter, err = file.NewStreamWriter(sheet.Name); err != nil {
			return
		}
		data := sheet.Data
		for rowID := 1; rowID <= len(data); rowID++ {
			colN := len(data[0])
			row := make([]interface{}, colN)
			for colID := 0; colID < colN; colID++ {
				row[colID] = data[rowID-1][colID]
			}
			cell, _ := excelize.CoordinatesToCellName(1, rowID)
			if err = streamWriter.SetRow(cell, row); err != nil {
				return
			}
		}
		if err = streamWriter.Flush(); err != nil {
			return
		}
	}
	var buf *bytes.Buffer
	if buf, err = file.WriteToBuffer(); err != nil {
		return
	}
	if dto.File == "" {
		if dto.File, err = gonanoid.New(); err != nil {
			return
		}
		dto.File = fmt.Sprintf(`%s.xlsx`, dto.File)
	}
	if err = x.Upload(ctx, buf, dto.File); err != nil {
		return
	}
	return
}

func (x *Service) Upload(ctx context.Context, buf *bytes.Buffer, name string) (err error) {
	option := x.Values.Storage
	var u *url.URL
	u, err = url.Parse(option.Url)
	b := &cos.BaseURL{BucketURL: u}
	switch x.Values.Storage.Type {
	case `cos`:
		client := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  option.Id,
				SecretKey: option.Key,
			},
		})
		if _, err = client.Object.Put(ctx, name, buf, nil); err != nil {
			return
		}
		break
	}
	return
}
