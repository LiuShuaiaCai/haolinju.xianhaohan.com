package component

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"haolinju.xianhaohan.com/internal/conf"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"time"
)

type MysqlConfig struct {
	Client conf.Db_Mysql `toml:"Client"`
}

func NewMysql(m *conf.Db_Mysql_Client) (db *gorm.DB, err error) {
	log.Info(nil, "[DB] new db client", log.Fields{
		"conf": m,
	})

	// 开启DB
	if db, err = gorm.Open("mysql", m.Dsn); err != nil {
		log.Error(nil, "[DB] new db client error", log.Fields{
			"error": err,
		})
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	if err = db.DB().PingContext(ctx); err != nil {
		log.Error(nil, "[DB] new db client error", log.Fields{
			"error": err,
		})
		return nil, err
	}

	/**
	 * SetMaxOpenConns设置与数据库的最大打开连接数
	 * 如果MaxIdleConns大于0，而新的MaxOpenConns小于MaxIdlecons，则MaxIdlecons将减少以匹配新的MaxOpenConns限制
	 * 如果n<=0，则对打开的连接数没有限制
	 * 默认值为0（无限制）
	 */
	db.DB().SetMaxOpenConns(int(m.Max))

	/**
	 * SetMaxIdleConns设置空闲连接池中的最大连接数
	 * 如果MaxOpenConns大于0但小于新的MaxIdleConns，则新的MaxIdleConns将减少以匹配MaxOpenConns限制
	 * 如果n<=0，则不保留空闲连接
	 * /默认的最大空闲连接数当前为2。这可能会在未来的版本中发生变化
	 */
	db.DB().SetMaxIdleConns(int(m.Idle))

	/**
	 * SetConnMaxLifetime设置可重复使用连接的最长时间
	 * 在重新使用之前，过期的连接可能会被延迟关闭
	 * 如果d<=0，则连接不会因连接的使用年限而关闭
	 */
	db.DB().SetConnMaxLifetime(time.Duration(m.LifeTime) * time.Second)

	// LogMode设置日志模式，对于详细日志为'true'，对于无日志为'false'，默认情况下，只打印错误日志
	db.LogMode(m.LogMode)

	return
}
