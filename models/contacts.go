package model

type Contacts struct {
	Id        int64  `json:"id"`
	User      string `json:"user"`
	Object    string `json:"object"`
	IsPrivate int64  `json:"is_private"`
	Sort      int64  `json:"sort"`
	LastTime  int64  `json:"last_time"`
}
