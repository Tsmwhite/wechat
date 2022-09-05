package model

import "wechat/helper/format"

const MessageTableCount = 10

var tableCountMap = map[string]int{
	"messages": MessageTableCount,
}

func GetTableName(sourceTable, uuid string) string {
	return format.HashCutTable(sourceTable, uuid, tableCountMap[sourceTable])
}
