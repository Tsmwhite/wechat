package services

import "time"

func NewResult() []map[string]interface{} {
	return []map[string]interface{}{}
}

func TimeLimit() int64 {
	return time.Now().Unix() - 3*30*24*60*60
}

type Pagination struct {
	Page   int `validate:"noRequired"`
	Size   int `validate:"noRequired"`
	LastId int `validate:"noRequired"`
}
