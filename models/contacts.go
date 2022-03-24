package model

const ContactsLimit = 500
type Contacts struct {
	Id        int64  `json:"id"`
	User      string `json:"user"`
	Object    string `json:"object"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	Name      string `json:"name"`
	IsPrivate int64  `json:"is_private"`
	Sort      int64  `json:"sort"`
	LastTime  int64  `json:"last_time"`
	IsDel     int64  `json:"is_del"`
}


