package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"wechat/msg"
)

//状态
const(
	ON_LINE  int = 1
	OFF_LINE int = 2
)

//客户端信息
type Client struct {
	Id			int				 //链接id
	conn 		*websocket.Conn  //socket连接
	SendMsg		chan *msg.Msg	 //消息
	Status 		int				 //状态
	IsClose 	bool			 //是否已关闭
	CloseCh		chan byte		 //关闭
	CloseGuard	sync.Mutex		 //防止CloseCh被关闭多次
}

//创建一个客户端
func NewClient(id int,conn *websocket.Conn) *Client{
	client := Client{
		Id:        id,
		conn:      conn,
		SendMsg:   make(chan *msg.Msg),
		Status:    ON_LINE,
		IsClose:   false,
	}

	client.start()
	return &client
}

func (c *Client) start(){
	//将消息
	go c.Write()
	go c.Read()
}

//将收到得消息推送给客户端
func (c *Client) Write(){
	defer c.Close()
	for {
		select {
			case msg,ok := <- c.SendMsg:
				if !ok || c.conn.WriteJSON(msg) != nil {
					break
				}
			case <- c.CloseCh:
				break

		}
	}
}

//读取客户端得信息转接到消息管理员
func (c *Client) Read(){
	defer c.Close()
	for {
		_,msg,err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		data := string(msg)
		fmt.Println("msg:",data)
		//心跳检测回复
		if data == "ping" {
			err := c.conn.WriteMessage(websocket.TextMessage,[]byte("pong"))
			if err != nil {
				break
			}
		}
	}
}

//关闭连接
func (c *Client) Close(){
	if !c.IsClose {
		c.IsClose = true
		c.conn.Close()
	}
}