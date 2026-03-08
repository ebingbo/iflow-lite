package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func RoleRouterInit(e *gin.Engine) {
	group := e.Group("/api/role")
	group.POST("/list", controller.RoleList)
}
