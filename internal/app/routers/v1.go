package routes

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/controllers"
)

func V1(router *gin.RouterGroup) {
	// 设置API V1路由组
	v1 := router.Group("v1")

	v1.GET("/test", controllers.Users)

}
