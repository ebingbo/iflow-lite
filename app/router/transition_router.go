package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func TransitionRouterInit(e *gin.Engine) {
	group := e.Group("/api/transition")
	group.GET("/get", controller.TransitionGet)
	group.POST("/add", controller.TransitionAdd)
	group.POST("/list", controller.TransitionList)
	group.POST("/delete/:id", controller.TransitionDelete)
	group.POST("/update", controller.TransitionUpdate)
}
