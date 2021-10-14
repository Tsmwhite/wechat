package roomer

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"wechat/lib/member"
)

type Roomer interface {
	GetMembers() []string
	Join(member.Member) error
	Quit(member.Member) error
}

const (
	IsPrivate = iota
	IsGroup
)

type Room struct {
	Uuid    string
	Title   string
	Type    int
	Members map[string]bool
	Creator string
}

func New(title string, creator member.Member, members map[string]bool, _type int) *Room {
	return &Room{
		Uuid:    uuid.NewV4().String(),
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