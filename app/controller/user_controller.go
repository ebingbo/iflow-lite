package controller

import (
	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

func UserGet(ctx *gin.Context) {
	var in input.UserGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserAdd(ctx *gin.Context) {
	var in input.UserAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserLogin(ctx *gin.Context) {
	var in input.UserLoginInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserLogin(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
