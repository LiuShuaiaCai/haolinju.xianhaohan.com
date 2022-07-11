package routes

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/controllers"
	"haolinju.xianhaohan.com/internal/app/controllers/comments"
)

func V1(router *gin.RouterGroup) {
	// 设置API V1路由组
	v1 := router.Group("v1")
	v1.GET("/test", controllers.Users)

	// 公共接口
	comment := v1.Group("comment")
	{
		// 获取文件上传的配置信息
		comment.GET("upload/config", comments.UploadConfig)
	}

}
