package conf

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

type Setter interface {
	Set([]byte) error
}

func Watch(s Setter) (*fsnotify.Watcher, error) {
	// 检查文件是否存在
	err := checkPathExists()
	if err != nil {
		return nil, err
	}

	// 读取配置文件
	text, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	err = s.Set(text)
	if err != nil {
		return nil, err
	}

	// 创建一个监控对象
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// 添加要监控的配置文件
	err = watcher.Add(confPath)
	if err != nil {
		return nil, err
	}

	// 开启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logrus.Info("watch config event", event)

				if event.Op&fsnotify.Write == fsnotify.Write {
					logrus.Info("modified file:", event.Name)
					text, err := ioutil.ReadFile(confPath)
					if err != nil {
						logrus.Error("read config file error:", err)
					}
					err = s.Set(text)
					if err != nil {
						logrus.Error("set config error:", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logrus.Error("config watch error:", err)
			}
		}
	}()

	return watcher, nil
}

// 检查配置文件是否存在
func checkPathExists() (err error) {
	if confPath == "" {
		confPathDir, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			return err
		}
		confPath = fmt.Sprintf("%s/%s", confPathDir, "configs/config.yaml")
	}

	// 文件信息
	s, err := os.Stat(confPath)
	if err != nil {
		return err
	}

	// 判断是否是文件
	if s.IsDir() {
		return errors.New("config is not a file")
	}

	// 判断是否是yaml文件
	fileExt := path.Ext(confPath)
	if fileExt != ".yaml" {
		return errors.New("config is not a yaml file")
	}

	return nil
}
