package client

import (
	"github.com/gorilla/websocket"
	"wechat/core/log"
	"wechat/core/message"
)

func NewClient(conn *websocket.Conn, id string, msgChan chan message.Messenger) *WsClient {
	return &WsClient{
		id:   id,
		conn: conn,
		Send: make(chan message.Messenger),
		Read: msgChan,
	}
}

type WsClient struct {
	id   string
	conn *websocket.Conn
	Send chan message.Messenger
	Read chan message.Messenger
}

func (c *WsClient) GetUuid() string {
	return c.id
}

func (c *WsClient) MessageProcess() {
	go c.read()
	go c.write()
}

//读取客户端消息
func (c *WsClient) read() {
	defer func() { c.close() }()
	var msg = &message.Message{}
	for {
		//读取消息
		err := c.conn.ReadJSON(msg)
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			return
		}
		if msg.Recipient == "" || msg.Content == "" {
			continue
		}
		switch msg.Type {
		case message.TypeHeartbeat:
			if err = c.conn.WriteMessage(websocket.PongMessage, nil); err != nil {
				c.close()
			}
		default:
			// 消息发送者为该连接
			msg.Sender = c.id
			c.Read <- msg
		}
	}
}

//将消息发送至客户端
func (c *WsClient) write() {
	defer func() { c.close() }()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				log.Error.Println("client -> write -> WriteMessage", c)
				return
			}
			if err := c.conn.WriteJSON(msg); err != nil {
				log.Error.Println("client -> write -> WriteJSON", err, c)
				return
			}
		}
	}
}

func (c *WsClient) close() {
	c.conn.Close()
}
