package controller

import (
	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

func ExecutionGet(ctx *gin.Context) {
	var in input.ExecutionGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultExecutionService.ExecutionGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
func ExecutionStart(ctx *gin.Context) {
	var in input.ExecutionStartInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultExecutionService.ExecutionStart(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
func ExecutionAdd(ctx *gin.Context) {
	var in input.ExecutionAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultExecutionService.ExecutionAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func ExecutionQuery(ctx *gin.Context) {
	var in input.ExecutionQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultExecutionService.ExecutionQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
