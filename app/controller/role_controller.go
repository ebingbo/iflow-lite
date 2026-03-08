package controller

import (
	"iflow-lite/core/http"
	"iflow-lite/service"
	"iflow-lite/type/input"

	"github.com/gin-gonic/gin"
)

func RoleList(ctx *gin.Context) {
	var in input.RoleListInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := service.DefaultRoleService.RoleList(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}
