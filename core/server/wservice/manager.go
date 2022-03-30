package wservice

import (
	"fmt"
	"wechat/core/client"
	"wechat/core/message"
)

type Manager interface {
	Register(*client.WsClient)
	Unregister(*client.WsClient)
	Broadcast(message.Messenger)
	GetBroadcastChan() chan message.Messenger
	GetClient(string) *client.WsClient
	Run()
}

// ClientManager 客户端管理
type ClientManager struct {
	//客户端 map 储存并管理所有的长连接client
	clients map[string]*client.WsClient
	//web端发送来的的message我们用broadcast来接收，并最后分到client
	broadcast chan message.Messenger
	//新创建的长连接client
	register chan *client.WsClient
	//新注销的长连接client
	unregister chan *client.WsClient
}

func NewManager() *ClientManager {
	return &ClientManager{
		broadcast:  make(chan message.Messenger, 1000),
		register:   make(chan *client.WsClient, 100),
		unregister: make(chan *client.WsClient, 100),
		clients:    make(map[string]*client.WsClient),
	}
}

func (m *ClientManager) Register(c *client.WsClient) {
	m.register <- c
}

func (m *ClientManager) Unregister(c *client.WsClient) {
	m.unregister <- c
}

func (m *ClientManager) Broadcast(messenger message.Messenger) {
	m.broadcast <- messenger
}

func (m *ClientManager) GetClient(uuid string) *client.WsClient {
	return m.clients[uuid]
}

func (m *ClientManager) GetBroadcastChan() chan message.Messenger {
	return m.broadcast
}

func (m *ClientManager) Run() {
	for {
		select {
		case conn := <-m.register:
			m.clients[conn.GetUuid()] = conn
			conn.MessageProcess()
		case conn := <-m.unregister:
			if _, ok := m.clients[conn.GetUuid()]; ok {
				close(conn.Send)
				delete(m.clients, conn.GetUuid())
			}
		case msg := <-m.broadcast:
			fmt.Println("msg", msg.GetContent())
			for _, mem := range msg.GetRecipientsUuid() {
				if mem != msg.GetSenderUuid() {
					wsc, ok := m.clients[mem]
					if ok {
						wsc.Send <- msg
					}
				}
			}
		}
	}
}
