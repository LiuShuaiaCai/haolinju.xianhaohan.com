package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/models"
	"haolinju.xianhaohan.com/internal/app/requests"
	"haolinju.xianhaohan.com/internal/pkg/response"
	"strconv"
)

func BannerSave(ctx *gin.Context) {
	req := new(models.Banner)
	err := ctx.Bind(req)
	if err != nil {
		return
	}

	err = Svc.BannerSave(ctx, req)

	response.JSON(ctx, nil, err)
}

func BannerList(ctx *gin.Context) {
	req := new(requests.BannerListReq)
	err := ctx.Bind(req)
	if err != nil {
		return
	}
	res, err := Svc.BannerList(ctx, req)

	response.JSON(ctx, res, err)
}

func BannerDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	bannerId, err := strconv.Atoi(id)
	if err != nil {
		response.JSON(ctx, nil, errors.New("参数错误"))
	}

	err = Svc.BannerDelete(bannerId)
	response.JSON(ctx, nil, err)
}
