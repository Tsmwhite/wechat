package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var _DB_ *gorm.DB

func init() {
	var err error
	dsn := "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
	_DB_, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256, // string 类型字段的默认长度
		DisableDatetimePrecision: true,
	}), &gorm.Config{})
	sqlDB, err := _DB_.DB()
	if err != nil {
		//TODO
		panic("db error")
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return _DB_
}
