package message

import (
	"strings"
	"wechat/core/roomer"
)

type Messenger interface {
	Send()
	Revoke()
	GetContent() string
	GetStatus() int
	GetRecipientsUuid() []string
	GetSenderUuid() string
	GetRoom() roomer.Roomer
}

type Message struct {
	Id          int
	Uuid        string
	Sender      string
	Recipient   string
	Type        int
	SecondType  int
	Status      int
	Content     string
	Parent      *Message
	Deletes     map[string]bool
	Reads       map[string]bool
	SendTime    int //客户端发送时间
	ReceiveTime int //服务端接收时间
	ForwardTime int //转发至接收者时间
	LogTime     int //消息入库时间
}

func (m Message) Send() {
	panic("implement me")
}

func (m Message) Revoke() {
	panic("implement me")
}

func (m Message) GetContent() string {
	return m.Content
}

func (m Message) GetStatus() int {
	return m.Status
}

func (m Message) GetRecipientsUuid() []string {
	return strings.Split(m.Recipient, ",")
}

func (m Message) GetSenderUuid() string {
	return m.Sender
}

func (m Message) GetRoom() roomer.Roomer {
	return nil
}
