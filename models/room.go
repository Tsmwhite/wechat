package model

import (
	"errors"
	"wechat/core/encrypt"
	"wechat/core/member"
)

const (
	IsPrivate = iota
	IsGroup
)

type Room struct {
	Uuid        string
	Title       string
	Type        int
	Members     map[string]bool
	Creator     string
	Description string
	CreateTime  int
}

func New(title string, creator member.Member, members map[string]bool, _type int) *Room {
	return &Room{
		Uuid:    encrypt.CreateUuid(),
		Title:   title,
		Type:    _type,
		Members: members,
		Creator: creator.GetUuid(),
	}
}

func (r *Room) GetMembers() []string {
	return nil
}

func (r *Room) Join(m member.Member) error {
	uuid := m.GetUuid()
	if r.Members[uuid] {
		return errors.New("")
	}
	r.Members[uuid] = true
	return nil
}

func (r *Room) Quit(m member.Member) error {
	uuid := m.GetUuid()
	if !r.Members[uuid] {
		return errors.New("")
	}
	delete(r.Members, uuid)
	return nil
}