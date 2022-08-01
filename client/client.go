package client

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/url"
)

type OpenAPI struct {
	Client *client.Client
	Url    string
}

type M = map[string]interface{}

func New(url string) (oapi *OpenAPI, err error) {
	oapi = new(OpenAPI)
	if oapi.Client, err = client.NewClient(client.WithResponseBodyStream(true)); err != nil {
		return
	}
	oapi.Url = url
	return
}

func (x *OpenAPI) R(method string, path string) (req *protocol.Request) {
	req = new(protocol.Request)
	req.SetMethod(method)
	req.SetRequestURI(x.Url + path)
	return
}

func (x *OpenAPI) Do(ctx context.Context, req *protocol.Request) (resp *protocol.Response, err error) {
	// TODO: 授权类型判断
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
