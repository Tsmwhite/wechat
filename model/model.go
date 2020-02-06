package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	USER_NAME    = "root"
	PASS_WORD    = "lxy196914"
	HOST         = "127.0.0.1"
	PORT         = "3306"
	DB_NAME      = "thewhite"
	CHARSET      = "utf8"
	MaxOpenConns = 100
	MaxIdleConns = 20
	MaxLifeTime  = 100 * time.Second
)

var _DB_ *sql.DB
var DB_ERR error

func init() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("model/init Error: ", err)
			log.Println("model/init Error: " + DB_ERR.Error())
		}
	}()
	driver := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DB_NAME, CHARSET)

	// 打开连接失败
	_DB_, DB_ERR = sql.Open("mysql", driver)
	//defer MysqlDb.Close();
	if DB_ERR != nil {
		panic("Database Connection Error: " + DB_ERR.Error())
	}

	// 最大连接数
	_DB_.SetMaxOpenConns(MaxOpenConns)
	// 闲置连接数
	_DB_.SetMaxIdleConns(MaxIdleConns)
	// 最大连接周期
	_DB_.SetConnMaxLifetime(MaxLifeTime)

	if DB_ERR = _DB_.Ping(); nil != DB_ERR {
		panic("Database Connection Error(ping): " + DB_ERR.Error())
	}

}

type Member struct {
	Id           int64  `json:"id" key:"AUTO_INCREMENT"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	Mail         string `json:"mail"`
	Salt         string `json:"salt"`
	Password     string `json:"password"`
	PassLook     string `json:"passlook"`
	Headimg      string `json:"headimg"`
	RoleId       int64  `json:"role_id"`
	IsHid        int64  `json:"is_hid"`
	IsDel        int64  `json:"is_del"`
	RegisterTime int64  `json:"register_time"`
	LoginTime    int64  `json:"login_time"`
	UpdateTime   int64  `json:"update_time"`
	RegisterIp   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
}
