package wservice

import (
	"wechat/core/client"
	"wechat/core/message"
)

// Handle 处理消息
func (m *ClientManager) Handle() {
	for {
		select {
		case conn := <-m.register:
			// 注册连接
			HandleRegister(m, conn)
		case conn := <-m.unregister:
			// 退出连接
			HandleUnregister(m, conn)
		case msg := <-m.broadcast:
			// 录入消息
			m.writeMsg <- msg
			// 转发消息
			HandleBroadcast(m, msg)
		}
	}
}

// Log 记录消息（入库）
func (m *ClientManager) Log() {
	for {
		select {
		case msg := <-m.writeMsg:
			msg.Save()
		}
	}
}

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
		conn.MessageProcess()
		m.SetClient(uuid, append(clients, conn))
	}
}

// HandleUnregister 处理断开连接
func HandleUnregister(m Manager, conn *client.WsClient) {
	uuid := conn.GetUuid()
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
	for _, mem := range msg.GetRecipientsUuid() {
		if mem != msg.GetSenderUuid() {
			for _, c := range m.GetClient(mem) {
				c.PullMsgChan <- msg
			}
		}
	}
}
