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
	//Handle(*wsClient.WsClient)
}
