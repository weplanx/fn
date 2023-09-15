package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/weplanx/fn/common"
	"github.com/xuri/excelize/v2"
	"io"
	"strings"
)

func (x *API) TencentCosExcel(ctx context.Context, dto Dto) (err error) {
	for _, record := range dto.Records {
		prefix := fmt.Sprintf(`/%s/%s/`, record.CosBucket.Appid, record.CosBucket.Name)
		key := strings.Replace(record.Cos.Key, prefix, "", -1)
		var resp *cos.Response
		if resp, err = x.Client.Object.Get(ctx, key, nil); err != nil {
			return
		}
		if err = x.TencentCosExcelParse(ctx, resp.Body); err != nil {
			return
		}
	}
	return
}

func (x *API) TencentCosExcelParse(ctx context.Context, body io.Reader) (err error) {
	var metadata common.ExcelMetadata
	if err = msgpack.NewDecoder(body).Decode(&metadata); err != nil {
		return
	}
	file := excelize.NewFile()
	defer file.Close()
	for _, key := range metadata.Parts {
		var streamWriter *excelize.StreamWriter
		args := strings.Split(key, ".")
		if streamWriter, err = file.NewStreamWriter(args[1]); err != nil {
			return
		}
		var resp *cos.Response
		if resp, err = x.Client.Object.Get(ctx, key, nil); err != nil {
			return
		}
		dec := msgpack.NewDecoder(resp.Body)
		rowID := 1
		for {
			var data []interface{}
			if err = dec.Decode(&data); err != nil {
				if err == io.EOF {
					break
				}
				return
			}
			cell, _ := excelize.CoordinatesToCellName(1, rowID)
			if err = streamWriter.SetRow(cell, data); err != nil {
				return
			}
			rowID++
		}
		if err = streamWriter.Flush(); err != nil {
			return
		}
	}
	var buf *bytes.Buffer
	if buf, err = file.WriteToBuffer(); err != nil {
		return
	}
	key := fmt.Sprintf(`%s.xlsx`, metadata.Name)
	if _, err = x.Client.Object.Put(ctx, key, buf, nil); err != nil {
		return
	}
	return
}
