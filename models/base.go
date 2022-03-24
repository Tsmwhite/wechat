package model

import (
	"gorm.io/gorm"
	"wechat/core/database"
)

const IsNoDel = 0
const IsDel = 1

const StatusNormal = 0

var DB *gorm.DB

func init() {
	DB = database.GetDB()
}

type Condition struct {
	Table    string
	Fields   []string
	Where    map[string]interface{}
	SqlStr   string
	Order    string
	Group    string
	Having   string
	Limit    int
	Offset   int
	Distinct []string
	Joins    []string
}

func FindAll(option *Condition, dest interface{}) {
	option.Parse(DB).Find(dest)
}

func (option *Condition) Parse(db *gorm.DB) *gorm.DB {
	if option.Table != "" {
		db = db.Table(option.Table)
	}
	if option.Fields != nil {
		db = db.Select(option.Fields)
	}
	if option.Where != nil {
		db = db.Where(option.Where)
	}
	if option.SqlStr != "" {
		db = db.Where(option.SqlStr)
	}
	if option.Order != "" {
		db = db.Order(option.Order)
	}
	if option.Group != "" {
		db = db.Group(option.Group)
	}
	if option.Having != "" {
		db = db.Having(option.Having)
	}
	if option.Limit > 0 {
		db = db.Limit(option.Limit)
	}
	if option.Offset > 0 {
		db = db.Offset(option.Offset)
	}
	if option.Distinct != nil {
		db = db.Distinct(option.Distinct)
	}
	if option.Joins != nil {
		for _, join := range option.Joins {
			db = db.Joins(join)
		}
	}
	return db
}
