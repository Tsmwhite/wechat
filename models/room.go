package model

import (
	"encoding/json"
	"strings"
	"wechat/core/member"
	"wechat/core/redis"
	"wechat/env"
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
	MemberNum   int    `json:"member_num"`
}

func GetRoomBuyUuid(roomId string) *Room {
	room := new(Room)
	redis.Get(env.RoomInfoKey, roomId, room)
	if room.Id == 0 {
		Find(&Condition{
			Table: "rooms",
			Where: map[string]interface{}{
				"uuid": roomId,
			},
		}, room)
		if room.Id == 0 {
			room.Id = -1
		}
		redis.Set(env.RoomInfoKey, roomId, room)
	}
	return room
}

func (r *Room) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
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
