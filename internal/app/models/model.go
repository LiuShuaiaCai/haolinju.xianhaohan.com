package models

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	"haolinju.xianhaohan.com/internal/conf"
	"haolinju.xianhaohan.com/internal/pkg/component"
)

var Provider = wire.NewSet(New)

type Model struct {
	db    *gorm.DB
	redis *redis.Client
}

// New new a Dao and return.
func New(conf *conf.AppConfig) (d *Model, cf func(), err error) {
	model := &Model{}

	// 链接数据库
	model.db, err = component.NewMysql(conf.Db.Mysql)

	// 链接redis

	// ...

	return model, model.Close, err
}

func (m *Model) Close() {
	sqlDB, _ := m.db.DB()
	_ = sqlDB.Close()
	//_ = d.redis.Close()
}
