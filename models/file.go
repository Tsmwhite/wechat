package model

type File struct {
	Id         int
	Uuid       string
	Path       string
	IsUse      int
	CreateTime int64
	UpdateTime int64
	IsDel      int
}
