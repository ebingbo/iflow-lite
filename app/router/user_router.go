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
	group.POST("/query", controller.UserQuery)
	group.POST("/status/update", controller.UserStatusUpdate)
	group.POST("/list", controller.UserList)
	group.POST("/role/list", controller.UserRoleList)
	group.POST("/role/update", controller.UserRoleUpdate)
	group.GET("/profile", controller.UserProfile)
	group.POST("/profile/update", controller.UserProfileUpdate)
	group.POST("/password/update", controller.UserPasswordUpdate)
	group.POST("/password/forgot", controller.UserPasswordForgot)
}
