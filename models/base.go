package model

import (
	"gorm.io/gorm"
	"wechat/lib/database"
)

var DB *gorm.DB

func init() {
	DB = database.GetDB()
}
