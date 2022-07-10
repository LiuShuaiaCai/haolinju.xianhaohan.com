package core

import (
	"context"
	"github.com/gin-gonic/gin"
	routes "haolinju.xianhaohan.com/internal/app/routers"
	"haolinju.xianhaohan.com/internal/app/services"
	"haolinju.xianhaohan.com/internal/conf"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"net/http"
	"time"
)

type App struct {
	Svc    *services.Service
	Server *http.Server
}

// TODO::备用
func NewApp(svc *services.Service, engine *gin.Engine) (app *App, closeFunc func(), err error) {
	// 初始化路由
	routes.Api(engine)

	// 服务配置
	server := &http.Server{
		Addr:           conf.Conf.Server.Http.Addr,
		Handler:        engine,
		ReadTimeout:    time.Duration(conf.Conf.Server.Http.Timeout) * time.Second,
		WriteTimeout:   time.Duration(conf.Conf.Server.Http.Timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(nil, "服务启动失败", log.Fields{"error": err})
			panic(err)
		}
	}()

	app = &App{
		Svc:    svc,
		Server: server,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			log.Error(nil, "Server Shutdown failed", log.Fields{"error": err})
			panic(err)
		}
		cancel()
	}
	return
}
