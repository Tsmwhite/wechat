package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"wechat/core/database"
	"wechat/core/log"
)

const IsNoDel = 0
const IsDel = 1

const StatusNormal = 0

func DB() *gorm.DB {
	return database.GetDB()
}

func NewDB() *gorm.DB {
	return database.NewDB()
}


const RecordDel = 1
const RecordNoDel = 0

func NewExpr(expr string, args ...interface{}) clause.Expr {
	return gorm.Expr(expr, args)
}

type Condition struct {
	Table      string
	Fields     []string
	Where      map[string]interface{}
	Expr       clause.Expr
	SqlStr     string
	Order      string
	Group      string
	Having     string
	Limit      int
	Offset     int
	Distinct   []string
	Joins      []string
	IncludeDel bool
}

type BaseModal struct{}

func (m *BaseModal) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *BaseModal) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func First(option *Condition, dest interface{}) {
	if err := option.Parse(DB()).First(dest).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error.Println("Find Error:", err, "\noption:", option)
		}
	}
}

func Find(option *Condition, dest interface{}) {
	if err := option.Parse(DB()).Find(dest).Error; err != nil {
		fmt.Println("Find", err)
		log.Error.Println("FindAll Error:", err, "\noption:", option)
	}
}

func (option *Condition) Parse(db *gorm.DB) *gorm.DB {
	// 查询未删除数据
	if !option.IncludeDel {
		db = db.Where(option.Table+".is_del = ?", RecordNoDel)
	}
	if option.Table != "" {
		db = db.Table(option.Table)
	}
	if option.Fields != nil {
		db = db.Select(option.Fields)
	}
	if option.Where != nil {
		db = db.Where(option.Where)
	}
	if option.Expr.SQL != "" {
		db = db.Where(option.Expr)
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
		if option.Limit == 0 {
			db = db.Limit(option.Offset)
		} else {
			db = db.Offset(option.Offset)
		}
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
