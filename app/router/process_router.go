package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func ProcessRouterInit(e *gin.Engine) {
	group := e.Group("/api/process")
	group.GET("/get", controller.ProcessGet)
	group.GET("/take", controller.ProcessTake)
	group.POST("/add", controller.ProcessAdd)
	group.POST("/update", controller.ProcessUpdate)
	group.POST("/query", controller.ProcessQuery)
	group.POST("/delete/:id", controller.ProcessDelete)
	group.POST("/disable/:id", controller.ProcessDisable)
	group.POST("/enable/:id", controller.ProcessEnable)
}
