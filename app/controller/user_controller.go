package controller

import (
	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

func UserGet(ctx *gin.Context) {
	var in input.UserGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultUserService.UserGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserProfile(ctx *gin.Context) {
	result, err := service.DefaultUserService.UserProfile(ctx.Request.Context())
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
	result, err := service.DefaultUserService.UserAdd(ctx.Request.Context(), &in)
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
	result, err := service.DefaultUserService.UserLogin(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserList(ctx *gin.Context) {
	var in input.UserListInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultUserService.UserList(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserProfileUpdate(ctx *gin.Context) {
	var in input.UserProfileUpdateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultUserService.UserProfileUpdate(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserPasswordUpdate(ctx *gin.Context) {
	var in input.UserPasswordUpdateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultUserService.UserPasswordUpdate(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserPasswordForgot(ctx *gin.Context) {
	var in input.UserPasswordForgotInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultUserService.UserPasswordForgot(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
