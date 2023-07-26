package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"strings"
	"tencent-workflow-excel/common"
	"time"
)

type API struct {
	*common.Inject
}

type M map[string]interface{}

type BodyDto struct {
	Key string `msgpack:"key"`
}

func (x *API) Invoke(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`已触发: %s`, time.Now())))
}

type Metadata struct {
	Name  string   `msgpack:"name"`
	Parts []string `msgpack:"parts"`
}

func (x *API) toExcel(ctx context.Context, body io.Reader) (err error) {
	var metadata Metadata
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
