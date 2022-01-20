package model

import (
	"gorm.io/gorm"
	"wechat/lib/database"
)

const IsNoDel = 0
const IsDel = 1

const StatusNormal = 0

var DB *gorm.DB

func init() {
	DB = database.GetDB()
}
