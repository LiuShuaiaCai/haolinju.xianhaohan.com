package models

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/requests"
	"haolinju.xianhaohan.com/internal/pkg/log"
)

type Banner struct {
	BaseModel
	Name      string `gorm:"column:name" json:"name"`
	ImgUrl    string `gorm:"column:img_url" json:"img_url"`
	JumpUrl   string `gorm:"column:jump_url" json:"jump_url"`
	Sort      int16  `gorm:"column:sort" json:"sort"`
	VillageId int16  `gorm:"column:village_id" json:"village_id"`
}

func (Banner) TableName() string {
	return "hh_banners"
}

// 保存banner
func (m *Model) SaveBanner(ctx *gin.Context, banner *Banner) (err error) {
	if err = m.db.Create(&banner).Error; err != nil {
		log.Error(ctx, "Banner保存失败", log.Fields{
			"banner": banner,
			"error":  err.Error(),
		})
	}

	return
}

// 查询Banner列表
func (m *Model) BannerList(ctx *gin.Context, villageId int16) (resp []requests.BannerListResp, err error) {
	if err = m.db.Model(&Banner{}).Where("village_id in (?)", []int16{0, villageId}).Order("sort asc").Find(&resp).Error; err != nil {
		log.Error(ctx, "Banner查询失败", log.Fields{
			"village_id": villageId,
			"error":      err.Error(),
		})
	}

	return
}

// 删除Banner
func (m *Model) BannerDelete(bannerId int) (err error) {
	err = m.db.Where("id", bannerId).Delete(&Banner{}).Error

	return err
}
