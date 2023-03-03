package model

import "wechat/helper/format"

const MessageTableCount = 10
const ReceiveMessageTableCount = 10

var tableCountMap = map[string]int{
	"messages":         MessageTableCount,
	"receive_messages": ReceiveMessageTableCount,
}

func GetTableName(sourceTable, uuid string) string {
	return format.HashCutTable(sourceTable, uuid, tableCountMap[sourceTable])
}
