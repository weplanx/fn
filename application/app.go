package application

import (
	"func-api/application/common"
	"func-api/application/controller/excel"
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

	return
}
