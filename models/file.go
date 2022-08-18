package model

type File struct {
	BaseModal
	Id         int
	Uuid       string
	Path       string
	IsUse      int
	CreateTime int64
	UpdateTime int64
	IsDel      int
}
