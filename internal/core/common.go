package core

import (
	"haolinju.xianhaohan.com/internal/conf"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"haolinju.xianhaohan.com/internal/pkg/validation"
)

func Init() (c *conf.AppConfig, f func(), err error) {
	// 初始化验证器
	err = validation.ValidatorLocalInit("zh")
	if err != nil {
		return nil, nil, err
	}

	// 初始化配置文件
	c, confClose, err := conf.Init()
	if err != nil {
		return nil, nil, err
	}

	// 初始化log
	log.Init()

	return c, func() {
		confClose()
	}, nil
}
