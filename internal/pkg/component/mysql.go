package component

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"haolinju.xianhaohan.com/internal/conf"
	"haolinju.xianhaohan.com/internal/pkg/log"
	"time"
)

type MysqlConfig struct {
	Client conf.Db_Mysql `toml:"Client"`
}

func NewMysql(m *conf.Db_Mysql) (db *gorm.DB, err error) {
	log.Info(nil, "[DB] new db client", log.Fields{
		"conf": m,
	})

	//设置表名的前缀(只会修改默认的表名规则)
	//db.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return m.TablePrefix + defaultTableName
	//}
	// 开启DB
	if db, err = gorm.Open(mysql.Open(m.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "hh_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
		},
	}); err != nil {
		log.Error(nil, "[DB] new db client error", log.Fields{
			"error": err,
		})
		return nil, err
	}

	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Error(nil, "[DB] new db pool error", log.Fields{
			"error": err,
		})
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	if err = sqlDB.PingContext(ctx); err != nil {
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
	sqlDB.SetMaxOpenConns(int(m.Max))

	/**
	 * SetMaxIdleConns设置空闲连接池中的最大连接数
	 * 如果MaxOpenConns大于0但小于新的MaxIdleConns，则新的MaxIdleConns将减少以匹配MaxOpenConns限制
	 * 如果n<=0，则不保留空闲连接
	 * /默认的最大空闲连接数当前为2。这可能会在未来的版本中发生变化
	 */
	sqlDB.SetMaxIdleConns(int(m.Idle))

	/**
	 * SetConnMaxLifetime设置可重复使用连接的最长时间
	 * 在重新使用之前，过期的连接可能会被延迟关闭
	 * 如果d<=0，则连接不会因连接的使用年限而关闭
	 */
	sqlDB.SetConnMaxLifetime(time.Duration(m.LifeTime) * time.Second)

	// LogMode设置日志模式，对于详细日志为'true'，对于无日志为'false'，默认情况下，只打印错误日志
	if m.LogMode == true {
		db = db.Debug()
	}

	return
}
