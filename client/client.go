package client

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Option struct {
	Key    string
	Secret string
}

type OpenAPI struct {
	Client *resty.Client
	*Option
}

type OptionFunc func(x *OpenAPI)

// SetCertification 设置认证
func SetCertification(key string, secret string) OptionFunc {
	return func(x *OpenAPI) {
		x.Key = key
		x.Secret = secret
	}
}

// New 新建客户端
func New(url string, options ...OptionFunc) *OpenAPI {
	x := new(OpenAPI)
	x.Option = new(Option)
	for _, v := range options {
		v(x)
	}
	x.Client = resty.New().SetBaseURL(url)
	x.Client.JSONMarshal = jsoniter.Marshal
	x.Client.JSONUnmarshal = jsoniter.Unmarshal
	return x
}

// R 创建请求
func (x *OpenAPI) R(method string, path string) *resty.Request {
	req := x.Client.R().
		SetHeader("accept", "application/json").
		SetHeader("source", "apigw test").
		SetHeader("x-date", time.Now().UTC().Format(http.TimeFormat))
	req.Method = method
	req.URL = path
	return req
}

// SetAuthorization 设置应用认证
func (x *OpenAPI) SetAuthorization(req *resty.Request) {
	var headers []string
	var headersKVString strings.Builder
	for k, _ := range req.Header {
		if k == "Accept" {
			continue
		}
		headers = append(headers, strings.ToLower(k))
	}
	sort.Strings(headers)
	for _, v := range headers {
		headersKVString.WriteString(fmt.Sprintf("%s: %s\n", v, req.Header.Get(v)))
	}
	accept := "application/json"
	contextMd5 := ""
	if req.Body != nil {
		hashMd5 := md5.New()
		hashMd5.Write(req.Body.([]byte))
		contextMd5 = hex.EncodeToString(hashMd5.Sum(nil))
	}
	pathAndParameters := req.URL
	if len(req.QueryParam) != 0 {
		pathAndParameters += fmt.Sprintf(`?%s`, req.QueryParam.Encode())
	}
	signToString := fmt.Sprintf("%s%s\n%s\n\n%s\n%s",
		headersKVString.String(), req.Method, accept, contextMd5, pathAndParameters,
	)
	hmacSha256 := hmac.New(sha256.New, []byte(x.Secret))
	hmacSha256.Write([]byte(signToString))
	signature := base64.StdEncoding.EncodeToString(hmacSha256.Sum(nil))
	Authorization := fmt.Sprintf(
		`hmac id="%s", algorithm="hmac-sha256", headers="%s", signature="%s"`,
		x.Key, strings.Join(headers, " "), signature,
	)
	req = req.SetHeader("Authorization", Authorization)
}

// Ping 测试
func (x *OpenAPI) Ping(ctx context.Context) (result map[string]interface{}, err error) {
	req := x.R(resty.MethodGet, "/")
	x.SetAuthorization(req)
	if _, err = req.SetContext(ctx).
		SetResult(&result).
		Send(); err != nil {
		return
	}
	return
}

func (x *OpenAPI) Ip(ctx context.Context, ip string) (result map[string]interface{}, err error) {
	req := x.R(resty.MethodGet, "/ip").
		SetQueryParam("value", ip)
	x.SetAuthorization(req)
	if _, err = req.SetContext(ctx).
		SetResult(&result).
		Send(); err != nil {
		return
	}
	return
}
