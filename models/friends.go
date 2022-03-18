package model

type Friends struct {
	Id int64 `json:"id"`
	User string `json:"user"`
	Friend string `json:"friend"`
	Remark string `json:"remark"`
	CreateTime int64 `json:"create_time"`
	Type int64 `json:"type"`
	Hot int64 `json:"hot"`
	IsDel int64 `json:"is_del"`
}