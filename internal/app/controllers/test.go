package controllers

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/pkg/response"
)

func Users(ctx *gin.Context) {
	res, err := Svc.Users(ctx)

	response.JSON(ctx, res, err)
}
