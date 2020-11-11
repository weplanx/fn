package application

import (
	"func-api/application/common"
	"func-api/application/controller"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func Application(router *gin.Engine, dep common.Dependency) (err error) {
	control := controller.New(&dep)
	router.POST("/simple_excel", common.Handle(control.SimpleExcel))
	router.POST("/new_task_for_excel", common.Handle(control.NewTaskForExcel))
	router.POST("/add_row_to_excel", common.Handle(control.AddRowToExcel))
	router.POST("/commit_task_for_excel", common.Handle(control.CommitTaskForExcel))
	return
}
