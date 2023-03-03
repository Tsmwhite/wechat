package model

type ReceiveMessage struct {
	Id        int64  `json:"id"`
	MsgUuid   string `json:"msg_uuid"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	IsRead    int64  `json:"is_read"`
	IsDel     int64  `json:"is_del"`
}
