package fn

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/weplanx/fn/common"
	"net/http"
	"net/url"
)

type Fn struct {
	Cos *cos.Client
}

type Option func(x *Fn) error

func SetCos(address string, id string, key string) Option {
	return func(x *Fn) (err error) {
		var u *url.URL
		if u, err = url.Parse(address); err != nil {
			return
		}
		baseURL := &cos.BaseURL{BucketURL: u}
		x.Cos = cos.NewClient(baseURL, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  id,
				SecretKey: key,
			},
		})
		return
	}
}

func New(options ...Option) (f *Fn, err error) {
	f = new(Fn)
	for _, v := range options {
		if err = v(f); err != nil {
			return
		}
	}
	return
}

func (x *Fn) TencentCosExcel(ctx context.Context, name string, sheets common.ExcelSheets) (err error) {
	metadata := common.ExcelMetadata{
		Name:  name,
		Parts: []string{},
	}
	for sheet, data := range sheets {
		w := bytes.NewBuffer(nil)
		enc := msgpack.NewEncoder(w)
		for _, v := range data {
			if err = enc.Encode(v); err != nil {
				return
			}
		}
		key := fmt.Sprintf(`%s.%s.pack`, name, sheet)
		metadata.Parts = append(metadata.Parts, key)
		if _, err = x.Cos.Object.Put(ctx, key, w, nil); err != nil {
			return
		}
	}

	var b []byte
	if b, err = msgpack.Marshal(metadata); err != nil {
		return
	}
	key := fmt.Sprintf(`%s.excel`, name)
	if _, err = x.Cos.Object.Put(ctx, key, bytes.NewBuffer(b), nil); err != nil {
		return
	}
	return
}
