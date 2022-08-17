package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"wechat/core/log"
	"wechat/core/message"
)

type OtherHandle struct {
	Type    int
	Options []interface{}
	Handle  func(...interface{})
}

type WsClient struct {
	id              string
	conn            *websocket.Conn
	Msg             message.Messenger
	PullMsgChan     chan message.Messenger   // 接受向web端发送消息的通道
	PushMsgChan     chan<- message.Messenger // 读取web端发送的消息，推送到消息处理中心
	UnregisterChan  chan<- *WsClient         // 读取web端发送的消息，推送到消息处理中心
	OtherHandleChan chan<- *OtherHandle      // 读取web端发送的消息，推送到消息处理中心
}

func NewClient(
	conn *websocket.Conn,
	id string,
	msgCarrier message.Messenger,
	managerReceiveCh chan<- message.Messenger,
	managerUnregisterCh chan<- *WsClient) *WsClient {
	return &WsClient{
		id:              id,
		conn:            conn,
		Msg:             msgCarrier,
		PullMsgChan:     make(chan message.Messenger, 100),
		PushMsgChan:     managerReceiveCh,
		UnregisterChan:  managerUnregisterCh,
		OtherHandleChan: make(chan *OtherHandle),
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
	c.UnregisterChan <- c
	if err := c.conn.Close(); err != nil {
		log.Error.Println("WsClient Close Error:", err)
	}
}

//读取客户端消息
func (c *WsClient) read() {
	defer func() { c.Close() }()
	for {
		//读取消息
		msg := c.Msg.New()
		err := c.conn.ReadJSON(msg)
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			return
		}
		switch msg.GetType() {
		case message.TypeHeartbeat:
			c.Pong()
			continue
		}
		if msg.GetContent() == "" || msg.GetRecipientsUuid() == nil {
			continue
		}
		// 将消息塞入消息中心
		if msg.GetSenderUuid() == "" {
			msg.SetSenderUuid(c.id)
		}
		fmt.Println("msg", msg)
		msg.SetReceiveTime()
		c.PushMsgChan <- msg
	}
}

//将消息发送至客户端
func (c *WsClient) write() {
	defer func() { c.Close() }()
	for {
		select {
		case msg, ok := <-c.PullMsgChan:
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

func (c *WsClient) Pong() {
	pongMsg := map[string]interface{}{
		"type":    0,
		"content": "pong",
	}
	err := c.conn.WriteJSON(pongMsg)
	if err != nil {
		log.Error.Println("Connect Pong Error:", err)
	}
}
