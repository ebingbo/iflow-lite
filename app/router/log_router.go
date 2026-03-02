package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func LogRouterInit(e *gin.Engine) {
	group := e.Group("/api/log")
	group.GET("/get", controller.LogGet)
	group.POST("/add", controller.LogAdd)
	group.POST("/query", controller.LogQuery)
}
