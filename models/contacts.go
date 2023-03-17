package model

const ContactsLimit = 500

type Contact struct {
	Id         int64  `json:"id"`
	Room       string `json:"room"`
	User       string `json:"user"`
	Friend     string `json:"friend"`
	Remark     string `json:"remark"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Type       int64  `json:"type"`
	Sort       int64  `json:"sort"`
	LastTime    int64 `json:"last_time"`
	LastMsg    string `json:"last_msg"`
	IsDel      int64  `json:"is_del"`
	UpdateTime int64  `json:"update_time"`
}
