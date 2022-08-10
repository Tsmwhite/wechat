package model

import (
	"gorm.io/gorm"
	"time"
	"wechat/core/encrypt"
	"wechat/core/message"
	"wechat/core/roomer"
	"wechat/helper/format"
)

type Message struct {
	BaseModal
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Sender      string `json:"sender"`
	Recipient   string `json:"recipient"` //接受消息房间uuid
	Type        int    `json:"type"`
	SecondType  int    `json:"second_type"`
	Status      int    `json:"status"`
	Content     string `json:"content"`
	Parent      string `json:"parent"`
	Deletes     string `json:"deletes"`
	Reads       string `json:"reads"`
	SendTime    int64  `json:"send_time"`    //客户端发送时间
	ReceiveTime int64  `json:"receive_time"` //服务端接收时间
	ForwardTime int64  `json:"forward_time"` //转发至接收者时间
	LogTime     int64  `json:"log_time"`     //消息入库时间
	UpdateTime  int64  `json:"update_time"`
	Remark      string `json:"remark"`
}

// SetTable 设置表名 根据房间id分表（10张）
func (m *Message) SetTable() *gorm.DB {
	return DB().Table(format.HashCutTable("messages", m.Recipient, 10))
}

func (m *Message) New() message.Messenger {
	return &Message{}
}

func (m *Message) Send() {
	panic("implement me")
}

func (m *Message) Revoke() {
	panic("implement me")
}

func (m *Message) GetContent() string {
	return m.Content
}

func (m *Message) GetStatus() int {
	return m.Status
}

func (m *Message) GetType() int {
	return m.Type
}

func (m *Message) GetRecipientsUuid() []string {
	return GetRoomBuyUuid(m.Recipient).GetMembers()
}

func (m *Message) GetSenderUuid() string {
	return m.Sender
}

func (m *Message) GetRoom() roomer.Roomer {
	return nil
}

func (m *Message) SetSenderUuid(uuid string) {
	m.Sender = uuid
}

func (m *Message) SetType(typeVal int) {
	m.Type = typeVal
}

func (m *Message) SetContent(content string) {
	m.Content = content
}

func (m *Message) SetReceiveTime() {
	m.ReceiveTime = time.Now().Unix()
}

func (m *Message) Save() {
	if m.Id > 0 {
		m.SetTable().Save(m)
	} else {
		now := time.Now().Unix()
		m.Uuid = encrypt.CreateUuid()
		m.LogTime = now
		m.UpdateTime = now
		m.SetTable().Create(m)
	}
}
