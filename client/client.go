package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/weplanx/openapi/api/excel"
	"github.com/weplanx/openapi/model"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	*client.Client
	*Option
}

type Option struct {
	Url    string
	Key    string
	Secret string
	Cos
}

type Cos struct {
	Url       string `env:"URL"`
	SecretID  string `env:"SECRETID"`
	SecretKey string `env:"SECRETKEY"`
}

type OptionFunc func(x *Client)

// SetApiGateway 设置网关认证
// https://cloud.tencent.com/document/product/628/55088
func SetApiGateway(key string, secret string) OptionFunc {
	return func(x *Client) {
		x.Key = key
		x.Secret = secret
	}
}

// SetCos 设置对象存储
// https://cloud.tencent.com/document/product/436
func SetCos(url string, id string, key string) OptionFunc {
	return func(x *Client) {
		x.Cos.Url = url
		x.Cos.SecretID = id
		x.Cos.SecretKey = key
	}
}

type M = map[string]interface{}

// New 创建
func New(url string, options ...OptionFunc) (x *Client, err error) {
	x = new(Client)

	if x.Client, err = client.NewClient(
		client.WithDialer(standard.NewDialer()),
		client.WithResponseBodyStream(true),
		client.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}),
	); err != nil {
		return
	}

	x.Option = &Option{Url: url}
	for _, v := range options {
		v(x)
	}

	return
}

// R 创建请求
func (x *Client) R(method string, path string) *OpenAPI {
	return &OpenAPI{
		Client: x,
		Method: method,
		Path:   path,
		Header: map[string]string{
			"accept": "application/json",
			"source": "apigw test",
			"x-date": time.Now().UTC().Format(http.TimeFormat),
		},
	}
}

type OpenAPI struct {
	Client *Client
	Method string
	Path   string
	Header map[string]string
	Query  url.Values
	Body   []byte
}

func (x *OpenAPI) SetHeaders(v map[string]string) *OpenAPI {
	for k, v := range v {
		x.Header[k] = v
	}
	return x
}

func (x *OpenAPI) SetQuery(v url.Values) *OpenAPI {
	x.Query = v
	return x
}

func (x *OpenAPI) SetData(v interface{}) *OpenAPI {
	x.Body, _ = sonic.Marshal(v)
	return x
}

func (x *OpenAPI) SetAuthorization() string {
	var headers []string
	var headersKVString strings.Builder
	for k := range x.Header {
		if k == "Accept" {
			continue
		}
		headers = append(headers, strings.ToLower(k))
	}
	sort.Strings(headers)
	for _, v := range headers {
		headersKVString.WriteString(fmt.Sprintf("%s: %s\n", v, x.Header[v]))
	}
	accept := "application/json"
	contextMd5 := ""
	if x.Method == "POST" {
		hashMd5 := md5.New()
		hashMd5.Write(x.Body)
		md5Str := hex.EncodeToString(hashMd5.Sum(nil))
		contextMd5 = base64.StdEncoding.EncodeToString([]byte(md5Str))
	}
	pathAndParameters := x.Path
	if x.Query != nil {
		pathAndParameters += fmt.Sprintf(`?%s`, x.Query.Encode())
	}
	signToString := fmt.Sprintf("%s%s\n%s\n\n%s\n%s",
		headersKVString.String(), x.Method, accept, contextMd5, pathAndParameters,
	)
	fmt.Println(signToString)
	hmacSha256 := hmac.New(sha256.New, []byte(x.Client.Secret))
	hmacSha256.Write([]byte(signToString))
	signature := base64.StdEncoding.EncodeToString(hmacSha256.Sum(nil))
	return fmt.Sprintf(
		`hmac id="%s", algorithm="hmac-sha256", headers="%s", signature="%s"`,
		x.Client.Key, strings.Join(headers, " "), signature,
	)
}

func (x *OpenAPI) Send(ctx context.Context) (resp *protocol.Response, err error) {
	req := new(protocol.Request)
	req.Header.SetMethod(x.Method)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI(fmt.Sprintf(`%s%s?%s`, x.Client.Url, x.Path, x.Query.Encode()))
	req.SetHeaders(x.Header)
	//if x.Client.Key != "" && x.Client.Secret != "" {
	//	req.SetHeader("Authorization", x.SetAuthorization())
	//}
	req.SetBody(x.Body)
	resp = new(protocol.Response)
	if err = x.Client.Do(ctx, req, resp); err != nil {
		return
	}
	return
}

// Ping 测试
func (x *Client) Ping(ctx context.Context) (result M, err error) {
	var resp *protocol.Response
	if resp, err = x.R("GET", "/").Send(ctx); err != nil {
		return
	}
	result = make(M)
	if err = sonic.Unmarshal(resp.Body(), &result); err != nil {
		return
	}
	return
}

// GetIp 获取 Ip
func (x *Client) GetIp(ctx context.Context, ip string) (data M, err error) {
	query := make(url.Values)
	query.Set("value", ip)
	var resp *protocol.Response
	if resp, err = x.R("GET", "/ip").
		SetQuery(query).
		Send(ctx); err != nil {
		return
	}
	data = make(M)
	if err = sonic.Unmarshal(resp.Body(), &data); err != nil {
		return
	}
	return
}

func (x *Client) GetCountries(ctx context.Context, fields []string) (data []model.Country, err error) {
	query := make(url.Values)
	query.Set("fields", strings.Join(fields, ","))
	var resp *protocol.Response
	if resp, err = x.R("GET", "/geo/countries").
		SetQuery(query).
		Send(ctx); err != nil {
		return
	}
	data = make([]model.Country, 0)
	if err = sonic.Unmarshal(resp.Body(), &data); err != nil {
		return
	}
	return
}

func (x *Client) GetStates(ctx context.Context, country string, fields []string) (data []model.State, err error) {
	query := make(url.Values)
	query.Set("country", country)
	query.Set("fields", strings.Join(fields, ","))
	var resp *protocol.Response
	if resp, err = x.R("GET", "/geo/states").
		SetQuery(query).
		Send(ctx); err != nil {
		return
	}
	data = make([]model.State, 0)
	if err = sonic.Unmarshal(resp.Body(), &data); err != nil {
		return
	}
	return
}

func (x *Client) GetCities(ctx context.Context, country string, state string, fields []string) (data []model.City, err error) {
	query := make(url.Values)
	query.Set("country", country)
	query.Set("state", state)
	query.Set("fields", strings.Join(fields, ","))
	var resp *protocol.Response
	if resp, err = x.R("GET", "/geo/cities").
		SetQuery(query).
		Send(ctx); err != nil {
		return
	}
	data = make([]model.City, 0)
	if err = sonic.Unmarshal(resp.Body(), &data); err != nil {
		return
	}
	return
}

func (x *Client) CreateExcel(ctx context.Context, dto excel.CreateDto) (data utils.H, err error) {
	var resp *protocol.Response
	if resp, err = x.R("POST", "/excel").
		SetData(dto).
		Send(ctx); err != nil {
		return
	}
	if err = sonic.Unmarshal(resp.Body(), &data); err != nil {
		return
	}
	return
}

type Sheets map[string][][]interface{}

func (x *Client) Excel(ctx context.Context, name string, sheets Sheets) (err error) {
	var u *url.URL
	u, err = url.Parse(x.Cos.Url)
	baseURL := &cos.BaseURL{BucketURL: u}
	cosClient := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  x.Cos.SecretID,
			SecretKey: x.Cos.SecretKey,
		},
	})

	var b []byte
	if b, err = msgpack.Marshal(sheets); err != nil {
		return
	}
	if _, err = cosClient.Object.Put(ctx, name, bytes.NewBuffer(b), nil); err != nil {
		return
	}

	return
}
