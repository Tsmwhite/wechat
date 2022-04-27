package model

import (
	"encoding/json"
	"strings"
	"wechat/core/message"
	"wechat/core/roomer"
)

type Message struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Sender      string `json:"sender"`
	Recipient   string `json:"recipient"`
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

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
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

//func (m Message) Handle(ws *wsClient.WsClient) {
//	switch m.Type {
//	case message.TypeHeartbeat:
//		if err := ws.GetConnect().WriteMessage(websocket.PongMessage, nil); err != nil {
//			ws.Close()
//		}
//	default:
//		// 消息发送者为该连接
//		m.SetSenderUuid(ws.GetUuid())
//		ws.Read <- m
//	}
//}
