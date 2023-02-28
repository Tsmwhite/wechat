package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"wechat/config"
	"wechat/env"
)

var _DB_ *gorm.DB

func initDB() *gorm.DB {
	dsn := config.DBEnv.Dsn
	fmt.Println("dsn", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256, // string 类型字段的默认长度
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	sqlDB, err := db.DB()
	if err != nil {
		//TODO
		panic("db error")
	}
	if env.Debug {
		db = db.Debug()
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.DBEnv.MaxIdleConns)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.DBEnv.MaxOpenConns)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.DBEnv.ConnMaxLifetime) * time.Minute)
	return db
}

func GetDB() *gorm.DB {
	if _DB_ == nil {
		_DB_ = initDB()
	}
	return _DB_
}

func NewDB() *gorm.DB {
	return initDB()
}
