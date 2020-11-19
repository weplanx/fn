package application

import (
	"func-api/application/common"
	"func-api/application/controller/excel"
	"func-api/application/controller/qrcode"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func Application(router *gin.Engine, dep common.Dependency) (err error) {
	excelGroup := router.Group("/excel")
	{
		control := excel.New(&dep)
		excelGroup.POST("/simple", common.Handle(control.Simple))
		excelGroup.POST("/new_task", common.Handle(control.NewTask))
		excelGroup.POST("/add_row", common.Handle(control.AddRow))
		excelGroup.POST("/commit_task", common.Handle(control.CommitTask))
	}
	qrGroup := router.Group("/qrcode")
	{
		qr := qrcode.New(&dep)
		qrGroup.POST("/testing", common.Handle(qr.Testing))
		qrGroup.POST("/pre_built", common.Handle(qr.PreBuilt))
		qrGroup.POST("/export", common.Handle(qr.Export))
	}
	return
}
