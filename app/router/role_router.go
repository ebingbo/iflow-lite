package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func RoleRouterInit(e *gin.Engine) {
	group := e.Group("/api/role")
	group.POST("/list", controller.RoleList)
	group.POST("/add", controller.RoleAdd)
	group.POST("/update", controller.RoleUpdate)
	group.POST("/delete/:id", controller.RoleDelete)
	group.POST("/query", controller.RoleQuery)
}
