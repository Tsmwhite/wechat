package socket

import(
	"fmt"
	"wechat/msg"
)


//客户端管理
type ClientManager struct {
	ClientAll 	map[int]*Client	//所有客户端
	Connect   	chan *Client	//新建连接
	Disconnect  chan *Client	//断开连接
	Broadcast	chan *msg.Msg	//消息
}

//创建一个管理员
func  RunMangerServer() *ClientManager {
	manager := &ClientManager{
		ClientAll:make(map[int]*Client),
		Connect:make(chan *Client),
		Disconnect:make(chan *Client),
		Broadcast:make(chan *msg.Msg),
	}
	go manager.start()
	return manager
}

//运行一个管理员
func (manager *ClientManager) start() {
	for{
		select {
		case  c,ok := <- manager.Connect:
			if ok {
				fmt.Println("Connect:",c)
				manager.RegisterConnect(c)
			}
		case  c,ok := <- manager.Disconnect:
			if ok {
				manager.LogoutConnect(c)
			}
		case  m,ok := <- manager.Broadcast:
			if ok {
				manager.PushBroadcast(m)
			}
		}
	}
}

//注册连接
func (manager *ClientManager) RegisterConnect(c *Client){
	id := c.Id
	manager.ClientAll[id] = c
}

//注销连接
func (manager *ClientManager) LogoutConnect(c *Client){
	id := c.Id
	delete(manager.ClientAll,id)
}

//发送消息
func (manager *ClientManager) PushBroadcast(msg *msg.Msg){
	for _,v := range manager.ClientAll {
		select {
			case v.SendMsg <- msg:
			default:
			close(v.SendMsg)
			delete(manager.ClientAll,v.Id)
		}
	}
}
