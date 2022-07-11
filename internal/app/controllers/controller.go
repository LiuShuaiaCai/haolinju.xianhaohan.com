package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	middleware "haolinju.xianhaohan.com/internal/app/middlewares"
	"haolinju.xianhaohan.com/internal/app/services"
)

var (
	Provider = wire.NewSet(New)
	Svc      *services.Service
)

type Controller struct {
	service *services.Service
}

// New new a bm server.
func New(s *services.Service) (engine *gin.Engine, err error) {
	Svc = s
	// 初始化引擎
	engine = gin.Default()

	// 加载中间件
	engine.Use(middleware.Trace())
	return
}
