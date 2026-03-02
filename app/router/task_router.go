package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func TaskRouterInit(e *gin.Engine) {
	group := e.Group("/api/task")
	group.GET("/get", controller.TaskGet)
	group.POST("/add", controller.TaskAdd)
	group.POST("/query", controller.TaskQuery)
}
