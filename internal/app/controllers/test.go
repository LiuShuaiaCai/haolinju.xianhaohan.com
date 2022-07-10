package controllers

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"haolinju.xianhaohan.com/internal/pkg/response"
)

func Users(ctx *gin.Context) {
	res, err := svc.Users(ctx)
	if err != nil {
		log.Error(ctx, err, nil)
		response.JSON(ctx, res, response.StatusInternalServerError)
		return
	}
	response.JSON(ctx, res, response.StatusOK)
}
