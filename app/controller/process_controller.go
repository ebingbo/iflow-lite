package controller

import (
	"fmt"

	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

var processService = service.NewProcessService()

func ProcessGet(ctx *gin.Context) {
	var in input.ProcessGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := processService.ProcessGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	http.JsonResponse(ctx, result)
}

func ProcessAdd(ctx *gin.Context) {
	var in input.ProcessAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := processService.ProcessAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
