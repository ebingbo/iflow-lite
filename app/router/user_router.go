package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(e *gin.Engine) {
	group := e.Group("/api/user")
	group.GET("/get", controller.UserGet)
	group.POST("/add", controller.UserAdd)
	group.POST("/login", controller.UserLogin)
}
