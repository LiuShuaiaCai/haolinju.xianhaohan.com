package models

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/pkg/log"
)

type User struct {
	Id     int
	Openid string
	Phone  string
	Name   string
}

func (m *Model) Users(ctx *gin.Context) (users []User, err error) {
	if err := m.db.Find(&users).Error; err != nil {
		log.Warn(ctx, "查询账号数据失败", log.Fields{
			"error": err,
		})
		return nil, err
	}

	return
}
