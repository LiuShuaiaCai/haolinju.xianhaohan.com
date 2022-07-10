package main

import (
	"haolinju.xianhaohan.com/internal/core"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化APP
	_, closeFunc, err := core.InitApp()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		s := <-c
		log.Info(nil, "get a signal", log.Fields{"signal": s.String()})
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info(nil, "api exit...", nil)
			closeFunc()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
