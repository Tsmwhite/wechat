package model

import (
	"strings"
	"wechat/core/member"
)

const (
	IsPrivate = iota
	IsGroup
)

type Room struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Type        int    `json:"type"`
	Members     string `json:"members"`
	Creator     string `json:"creator"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
	IsDel       int    `json:"is_del"`
}

func (r *Room) GetMembers() []string {
	return strings.Split(r.Members, ",")
}

func (r *Room) Join(m member.Member) error {

	return nil
}

func (r *Room) Quit(m member.Member) error {

	return nil
}