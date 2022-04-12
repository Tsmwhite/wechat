package model

import (
	"github.com/gorilla/websocket"
	"strings"
	wsClient "wechat/core/client"
	"wechat/core/message"
	"wechat/core/roomer"
)

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
	SendTime    int64 //客户端发送时间
	ReceiveTime int64 //服务端接收时间
	ForwardTime int64 //转发至接收者时间
	LogTime     int64 //消息入库时间
	Remark		string
}

func (m Message) New() message.Messenger {
	return &Message{}
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

func (m Message) GetType() int {
	return m.Type
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

func (m Message) SetSenderUuid(uuid string) {
	m.Sender = uuid
}

func (m Message) Handle(ws *wsClient.WsClient) {
	switch m.Type {
	case message.TypeHeartbeat:
		if err := ws.GetConnect().WriteMessage(websocket.PongMessage, nil); err != nil {
			ws.Close()
		}
	default:
		// 消息发送者为该连接
		m.SetSenderUuid(ws.GetUuid())
		ws.Read <- m
	}
}
