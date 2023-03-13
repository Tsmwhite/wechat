package model

import "wechat/helper/format"

const MessageTableCount = 10
const ReceiveMessageTableCount = 20

var tableCountMap = map[string]int{
	"messages":         MessageTableCount, //根据接收房间分表
	"receive_messages": ReceiveMessageTableCount, //根据接收者分表
}

func GetTableName(sourceTable, uuid string) string {
	return format.HashCutTable(sourceTable, uuid, tableCountMap[sourceTable])
}
