package services

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/models"
	"haolinju.xianhaohan.com/internal/app/requests"
)

func (s *Service) BannerSave(ctx *gin.Context, banner *models.Banner) (err error) {
	return s.model.SaveBanner(ctx, banner)
}

func (s *Service) BannerList(ctx *gin.Context, req *requests.BannerListReq) (resp []requests.BannerListResp, err error) {
	return s.model.BannerList(ctx, req.VillageId)
}

func (s *Service) BannerDelete(bannerId int) (err error) {
	return s.model.BannerDelete(bannerId)
}
