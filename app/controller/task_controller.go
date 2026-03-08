package controller

import (
	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

func TaskGet(ctx *gin.Context) {
	var in input.TaskGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultTaskService.TaskGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func TaskAdd(ctx *gin.Context) {
	var in input.TaskAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultTaskService.TaskAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func TaskQuery(ctx *gin.Context) {
	var in input.TaskQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultTaskService.TaskQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func TaskClaimableQuery(ctx *gin.Context) {
	var in input.TaskClaimableQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultTaskService.TaskClaimableQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func TaskCandidateList(ctx *gin.Context) {
	var in input.TaskCandidateListInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultTaskService.TaskCandidateList(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func TaskComplete(ctx *gin.Context) {
	var in input.TaskCompleteInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := service.DefaultTaskService.TaskComplete(ctx.Request.Context(), &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}

func TaskClaim(ctx *gin.Context) {
	var in input.TaskClaimInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := service.DefaultTaskService.TaskClaim(ctx.Request.Context(), &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}

func TaskSkip(ctx *gin.Context) {
	var in input.TaskSkipInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := service.DefaultTaskService.TaskSkip(ctx.Request.Context(), &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}

func TaskDelegate(ctx *gin.Context) {
	var in input.TaskDelegateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := service.DefaultTaskService.TaskDelegate(ctx.Request.Context(), &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}
