package controller

import (
	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

func ProcessGet(ctx *gin.Context) {
	var in input.ProcessGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func ProcessGetByID(ctx *gin.Context) {
	var in input.ProcessGetByIDInput
	if err := ctx.ShouldBindUri(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessGetByID(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func ProcessDisable(ctx *gin.Context) {
	var in input.ProcessDisableInput
	if err := ctx.ShouldBindUri(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessDisable(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func ProcessEnable(ctx *gin.Context) {
	var in input.ProcessEnableInput
	if err := ctx.ShouldBindUri(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessEnable(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func ProcessDelete(ctx *gin.Context) {
	var in input.ProcessDeleteInput
	if err := ctx.ShouldBindUri(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessDelete(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func ProcessTake(ctx *gin.Context) {
	var in input.ProcessTakeInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessTake(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
func ProcessAdd(ctx *gin.Context) {
	var in input.ProcessAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
func ProcessUpdate(ctx *gin.Context) {
	var in input.ProcessUpdateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessUpdate(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
func ProcessQuery(ctx *gin.Context) {
	var in input.ProcessQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultProcessService.ProcessQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
