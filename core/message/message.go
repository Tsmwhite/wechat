package message

import (
	"wechat/core/roomer"
)

type Messenger interface {
	New() Messenger
	Send()
	Revoke()
	GetContent() string
	GetStatus() int
	GetType() int
	GetRecipientsUuid() []string
	GetSenderUuid() string
	GetRoom() roomer.Roomer
	SetSenderUuid(string)
	SetContent(string)
	SetType(int)
	SetReceiveTime()
	Save()
	WhetherToRecord() bool
	Format()
	//Handle(*wsClient.WsClient)
}
