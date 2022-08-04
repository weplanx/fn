package client

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
	"net/url"
	"time"
)

type OpenAPI struct {
	Client *client.Client
	Url    string
	*Option
}

type Option struct {
	Key    string
	Secret string
}

type OptionFunc func(x *OpenAPI)

// SetApiGateway 设置网关认证
// https://cloud.tencent.com/document/product/628/55088
func SetApiGateway(key string, secret string) OptionFunc {
	return func(x *OpenAPI) {
		x.Key = key
		x.Secret = secret
	}
}

type M = map[string]interface{}

// New 创建客户端
func New(url string, options ...OptionFunc) (oapi *OpenAPI, err error) {
	oapi = new(OpenAPI)

	if oapi.Client, err = client.NewClient(
		client.WithResponseBodyStream(true),
	); err != nil {
		return
	}
	oapi.Url = url
	oapi.Option = new(Option)
	for _, v := range options {
		v(oapi)
	}

	return
}

// R 创建请求
func (x *OpenAPI) R(method string, path string) (req *protocol.Request) {
	req = new(protocol.Request)
	req.SetMethod(method)
	req.SetRequestURI(x.Url + path)
	req.SetHeader("accept", "application/json")
	req.SetHeader("source", "apigw test")
	req.SetHeader("x-date", time.Now().UTC().Format(http.TimeFormat))
	return
}

// Do 发起请求
func (x *OpenAPI) Do(ctx context.Context, req *protocol.Request) (resp *protocol.Response, err error) {
	//var headers []string
	//var headersKVString strings.Builder
	//req.Header.Header()
	//
	//for k, _ := range req.SetHeader() {
	//	if k == "Accept" {
	//		continue
	//	}
	//	headers = append(headers, strings.ToLower(k))
	//}
	//
	//accept := "application/json"
	//contextMd5 := ""
	//if req.Body() != nil {
	//	hashMd5 := md5.New()
	//	hashMd5.Write(req.Body())
	//	contextMd5 = hex.EncodeToString(hashMd5.Sum(nil))
	//}

	resp = new(protocol.Response)
	if err = x.Client.Do(ctx, req, resp); err != nil {
		return
	}
	return
}

// Ping 测试
func (x *OpenAPI) Ping(ctx context.Context) (result M, err error) {
	req := x.R(consts.MethodGet, "")
	var resp = new(protocol.Response)
	if resp, err = x.Do(ctx, req); err != nil {
		return
	}
	result = make(M)
	if err = json.NewDecoder(resp.BodyStream()).Decode(&result); err != nil {
		return
	}
	return
}

// Ip 获取 Ip
func (x *OpenAPI) Ip(ctx context.Context, ip string) (result M, err error) {
	req := x.R(consts.MethodGet, "/ip")
	params := url.Values{}
	params.Set("ip", ip)
	req.SetQueryString(params.Encode())
	var resp = new(protocol.Response)
	if resp, err = x.Do(ctx, req); err != nil {
		return
	}
	result = make(M)
	if err = json.NewDecoder(resp.BodyStream()).Decode(&result); err != nil {
		return
	}
	return
}
