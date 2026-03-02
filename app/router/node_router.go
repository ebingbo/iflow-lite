package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func NodeRouterInit(e *gin.Engine) {
	group := e.Group("/api/node")
	group.GET("/get", controller.NodeGet)
	group.POST("/add", controller.NodeAdd)
	group.POST("/query", controller.NodeQuery)
	group.POST("/list", controller.NodeList)
}
