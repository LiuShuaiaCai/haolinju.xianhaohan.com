package routes

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/controllers"
	"haolinju.xianhaohan.com/internal/app/controllers/commons"
)

func V1(router *gin.RouterGroup) {
	// 设置API V1路由组
	v1 := router.Group("v1")
	v1.GET("/test", controllers.Users)

	// 公共接口
	comment := v1.Group("common")
	{
		// 获取文件上传的配置信息
		comment.GET("upload/config", commons.UploadConfig)
	}

	// Banner接口
	banner := v1.Group("banner")
	{
		banner.POST("save", controllers.BannerSave)
		banner.GET("list", controllers.BannerList)
		banner.POST("delete/:id", controllers.BannerDelete)
	}

}
