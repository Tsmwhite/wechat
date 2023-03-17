package model

import "gorm.io/gorm"

const MessageReadStatus = 1

type ReceiveMessage struct {
	Id         int64  `json:"id"`
	MsgUuid    string `json:"msg_uuid"`
	Sender     string `json:"sender"`
	Recipient  string `json:"recipient"` //接收消息用户
	Room       string `json:"room"`      //接收消息房间
	SecondType int    `json:"second_type"`
	Content    string `json:"content"`
	IsRead     int64  `json:"is_read"`
	IsDel      int64  `json:"is_del"`
}

// SetTable 设置表名 根据接收者id分表（20张）
func (m *ReceiveMessage) SetTable() *gorm.DB {
	return DB().Table(GetTableName("receive_messages", m.Recipient))
}
