package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"wechat/core/log"
	"wechat/core/message"
)

type WsClient struct {
	id   string
	conn *websocket.Conn
	Msg  message.Messenger
	Send chan message.Messenger   // 接受向web端发送消息的通道
	Read chan<- message.Messenger // 读取web端发送的消息，传递到消息处理中心
}

func NewClient(conn *websocket.Conn, id string, msgCarrier message.Messenger, managerReceiveCh chan<- message.Messenger) *WsClient {
	return &WsClient{
		id:   id,
		conn: conn,
		Msg:  msgCarrier,
		Send: make(chan message.Messenger, 100),
		Read: managerReceiveCh,
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
	if err := c.conn.Close(); err != nil {
		log.Error.Println("WsClient Close Error:", err)
	}
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
		if msg.GetType() == message.TypeHeartbeat || msg.GetContent() == "" || msg.GetRecipientsUuid() == nil {
			continue
		}
		// 将消息塞入消息中心
		fmt.Println("msg.GetSenderUuid()", msg.GetSenderUuid())
		if msg.GetSenderUuid() == "" {
			fmt.Println("c.id", c.id)
			msg.SetSenderUuid(c.id)
		}
		fmt.Println("msg", msg)
		c.Read <- msg
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
