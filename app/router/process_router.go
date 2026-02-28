package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func ProcessRouterInit(e *gin.Engine) {
	group := e.Group("/api/process")
	group.GET("/get", controller.ProcessGet)
	group.POST("/add", controller.ProcessAdd)
}
