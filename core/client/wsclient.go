package client

import (
	"github.com/gorilla/websocket"
	"wechat/core/log"
	"wechat/core/message"
)

type WsClient struct {
	id   string
	conn *websocket.Conn
	Msg  message.Messenger
	Send chan message.Messenger // 向该链接发送消息的通道
	Read chan message.Messenger // 读取该链接发送的消息，传递到消息处理中心
}

func NewClient(conn *websocket.Conn, id string, msgHandel message.Messenger, read chan message.Messenger) *WsClient {
	return &WsClient{
		id:   id,
		conn: conn,
		Msg:  msgHandel,
		Send: make(chan message.Messenger, 100),
		Read: read,
	}
}

func (c *WsClient) GetConnect() *websocket.Conn {
	return c.conn
}

func (c *WsClient) GetUuid() string {
	return c.id
}

func (c *WsClient) MessageProcess() {
	go c.read()
	go c.write()
}

func (c *WsClient) Close() {
	c.conn.Close()
}

//读取客户端消息
func (c *WsClient) read() {
	defer func() { c.Close() }()
	var msg = c.Msg.New()
	for {
		//读取消息
		err := c.conn.ReadJSON(msg)
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			return
		}
		if msg.GetRecipientsUuid() == nil || msg.GetContent() == "" {
			continue
		}

	}
}

//将消息发送至客户端
func (c *WsClient) write() {
	defer func() { c.Close() }()
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
