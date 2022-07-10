package models

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"haolinju.xianhaohan.com/internal/conf"
	"haolinju.xianhaohan.com/internal/pkg/component"
)

var Provider = wire.NewSet(New)

type Model struct {
	hlj   *gorm.DB
	redis *redis.Client
}

// New new a Dao and return.
func New(conf *conf.AppConfig) (d *Model, cf func(), err error) {
	model := &Model{}

	// 链接数据库
	model.hlj, err = component.NewMysql(conf.Db.Mysql.Haolinju)

	// 链接redis

	// ...

	return model, model.Close, err
}

func (d *Model) Close() {
	_ = d.hlj.Close()
	//_ = d.redis.Close()
}
