package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"haolinju.xianhaohan.com/common"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"time"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		requestId := c.GetHeader(common.RequestId)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(common.RequestId, requestId)
		// 请求前

		c.Next()

		// 请求后

		latency := time.Since(t) / 1e6
		log.Info(c, "request info", log.Fields{
			"path":    c.FullPath(),
			"latency": fmt.Sprintf("%v ms", latency),
		})
		//log.Print(latency)
		// 获取发送的 status
		//status := c.Writer.Status()
		//log.Println(status)
	}
}
