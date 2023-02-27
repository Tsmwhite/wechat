package model

const ContactsLimit = 500

type Contact struct {
	Id         int64  `json:"id"`
	Room       string `json:"room"`
	User       string `json:"user"`
	Friend     string `json:"friend"`
	Avatar     string `json:"avatar"`
	Remark     string `json:"remark"`
	Name       string `json:"name"`
	Type       int64  `json:"type"`
	Sort       int64  `json:"sort"`
	LastTime   int64  `json:"last_time"`
	IsDel      int64  `json:"is_del"`
	UpdateTime int64  `json:"update_time"`
}
