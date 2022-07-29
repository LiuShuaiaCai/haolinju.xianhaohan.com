package commons

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/internal/app/controllers"
	"haolinju.xianhaohan.com/internal/pkg/response"
)

func UploadConfig(ctx *gin.Context) {
	res, err := controllers.Svc.UploadConfig(ctx)
	fmt.Println(res)
	response.JSON(ctx, res, err)
}
