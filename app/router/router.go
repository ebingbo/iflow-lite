package router

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	DocsRouterInit(r)
	MetricRouterInit(r)
	UserRouterInit(r)
	ProcessRouterInit(r)
	ExecutionRouterInit(r)
	AssignmentRouterInit(r)
	NodeRouterInit(r)
	TransitionRouterInit(r)
	TaskRouterInit(r)
	LogRouterInit(r)
}
