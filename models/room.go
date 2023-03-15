package model

import (
	"strings"
	"wechat/core/log"
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
		First(&Condition{
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

func (r *Room) GetMembers() []string {
	return strings.Split(r.Members, ",")
}

func (r *Room) Join(m member.Member) error {
	return nil
}

func (r *Room) Quit(m member.Member) error {
	return nil
}

func (r *Room) GetUnreadMsgCountByUser(u *User) int64 {
	if r.Type == IsPrivate {
		return getPrivateRoomUnreadMsgCount(r, u)
	} else {
		return getGroupRoomUnreadMsgCount(r, u)
	}
}

func getPrivateRoomUnreadMsgCount(r *Room, u *User) int64 {
	table := GetTableName("messages", r.Uuid)
	var count int64
	err := DB().Table(table).
		Where("id > ANY (SELECT IFNULL(MAX(id),0) FROM `"+table+"` WHERE `recipient` = ?  AND `sender` = ?)", r.Uuid, u.Uuid).
		Where("recipient = ?", r.Uuid).
		Where("sender <> ?", u.Uuid).
		Where("`reads` <> ?", u.Uuid).
		Count(&count).Error
	if err != nil {
		log.PrintlnErr("getPrivateRoomUnreadMsgCount Error:", err)
	}
	return count
}

func getGroupRoomUnreadMsgCount(r *Room, u *User) int64 {
	return 0
}
