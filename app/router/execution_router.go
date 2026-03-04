package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func ExecutionRouterInit(e *gin.Engine) {
	group := e.Group("/api/execution")
	group.POST("/add", controller.ExecutionAdd)
	group.GET("/get", controller.ExecutionGet)
	group.POST("/query", controller.ExecutionQuery)
	group.POST("/start", controller.ExecutionStart)
}
