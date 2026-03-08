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
	group.POST("/query/claimable", controller.TaskClaimableQuery)
	group.POST("/candidate/list", controller.TaskCandidateList)
	group.POST("/claim", controller.TaskClaim)
	group.POST("/complete", controller.TaskComplete)
	group.POST("/skip", controller.TaskSkip)
	group.POST("/delegate", controller.TaskDelegate)
}
