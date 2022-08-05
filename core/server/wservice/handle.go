package wservice

import (
	"fmt"
	"wechat/core/client"
	"wechat/core/message"
)

// HandleRegister 处理注册连接
func HandleRegister(m Manager, conn *client.WsClient) {
	uuid := conn.GetUuid()
	exist := false
	clients := m.GetClient(uuid)
	for _, c := range clients {
		if c == conn {
			exist = true
			break
		}
	}
	if !exist {
		m.SetClient(uuid, append(clients, conn))
	}
}

// HandleUnregister 处理断开连接
func HandleUnregister(m Manager, conn *client.WsClient) {
	uuid := conn.GetUuid()
	close(conn.Send)
	var temp []*client.WsClient
	for _, c := range m.GetClient(uuid) {
		if c != conn {
			temp = append(temp, c)
		}
	}
	if len(temp) < 1 {
		m.DelClient(uuid)
	} else {
		m.SetClient(uuid, temp)
	}
}

// HandleBroadcast 转发消息至客户端
func HandleBroadcast(m Manager, msg message.Messenger) {
	fmt.Println("msg", msg.GetContent(), msg.GetRecipientsUuid())
	for _, mem := range msg.GetRecipientsUuid() {
		if mem != msg.GetSenderUuid() {
			for _, c := range m.GetClient(mem) {
				c.Send <- msg
			}
		}
	}
}
