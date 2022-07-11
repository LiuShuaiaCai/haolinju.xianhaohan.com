package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/models"
)

func (s *Service) Users(ctx *gin.Context) (users []models.User, err error) {
	return nil, errors.New("参数错误")
	return s.model.Users(ctx)
}
