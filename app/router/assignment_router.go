package router

import (
	"iflow-lite/controller"

	"github.com/gin-gonic/gin"
)

func AssignmentRouterInit(e *gin.Engine) {
	group := e.Group("api/assignment")
	group.GET("/get", controller.AssignmentGet)
	group.POST("/query", controller.AssignmentQuery)
	group.POST("/add", controller.AssignmentAdd)
}
