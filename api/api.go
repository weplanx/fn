package api

import (
	"context"
	"github.com/bytedance/go-tagexpr/v2/binding"
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/bytedance/sonic/decoder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/google/wire"
	"github.com/weplanx/openapi/api/geo"
	"github.com/weplanx/openapi/api/index"
	"github.com/weplanx/openapi/api/qrcode"
	"github.com/weplanx/openapi/common"
	"net/http"
	"time"
)

var Provides = wire.NewSet(
	index.Provides,
	geo.Provides,
	qrcode.Provides,
)

type API struct {
	*common.Inject

	Hertz *server.Hertz

	IndexController  *index.Controller
	IndexService     *index.Service
	GeoController    *geo.Controller
	GeoService       *geo.Service
	QrcodeController *qrcode.Controller
	QrcodeService    *qrcode.Service
}

func (x *API) Routes() (h *server.Hertz, err error) {
	h = x.Hertz
	h.Use(x.AccessLog())
	h.Use(x.ErrHandler())

	h.GET("/", x.IndexController.Index)
	h.GET("/ip", x.IndexController.GetIp)

	_geo := h.Group("/geo")
	{
		_geo.GET("/countries", x.GeoController.Countries)
		_geo.GET("/states", x.GeoController.States)
		_geo.GET("/cities", x.GeoController.Cities)
	}

	return
}

// AccessLog 日志
func (x *API) AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Microseconds
		hlog.CtxTracef(c, "status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host())
	}
}

// ErrHandler 错误处理
func (x *API) ErrHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
		err := c.Errors.Last()
		if err == nil {
			return
		}

		if err.IsType(errors.ErrorTypePublic) {
			statusCode := http.StatusBadRequest
			result := utils.H{"message": err.Error()}
			if meta, ok := err.Meta.(map[string]interface{}); ok {
				if meta["statusCode"] != nil {
					statusCode = meta["statusCode"].(int)
				}
				if meta["code"] != nil {
					result["code"] = meta["code"]
				}
			}
			c.JSON(statusCode, result)
			return
		}

		switch anyType := err.Err.(type) {
		case decoder.SyntaxError:
			c.JSON(http.StatusBadRequest, utils.H{
				"message": anyType.Description(),
			})
			break
		case *binding.Error:
			c.JSON(http.StatusBadRequest, utils.H{
				"message": anyType.Error(),
			})
			break
		case *validator.Error:
			c.JSON(http.StatusBadRequest, utils.H{
				"message": anyType.Error(),
			})
			break
		default:
			logger.Error(err)
			c.Status(http.StatusInternalServerError)
		}
	}
}
