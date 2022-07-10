package routes

import "github.com/gin-gonic/gin"

func Api(router *gin.Engine) {
	api := router.Group("api")

	// V1版本
	V1(api)
}
