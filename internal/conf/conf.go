package conf

import (
	"github.com/google/wire"
	"gopkg.in/yaml.v3"
	"haolinju.xianhaohan.com/internal/pkg/conf"
)

var (
	Conf     *AppConfig
	Provider = wire.NewSet(Init)
)

func (c *AppConfig) Set(text []byte) error {
	var appConfig AppConfig
	if err := yaml.Unmarshal(text, &appConfig); err != nil {
		return err
	}
	Conf = &appConfig
	return nil
}

func Init() (c *AppConfig, fc func(), err error) {
	watcher, err := conf.Watch(Conf)
	if err != nil {
		return nil, nil, err
	}

	return Conf, func() {
		_ = watcher.Close()
	}, err
}
