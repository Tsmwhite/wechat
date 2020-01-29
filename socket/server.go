package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

//建立连接后的返回token
type Token string

type Server struct {
	Addr 		string
	WSPath		string
	upgrader	*websocket.Upgrader
	AuthToken	func(token Token) (id int,err error)
}

//
var ManagerDefault *ClientManager

//开启一个服务
func (s *Server) ListenServer()( err error){
	fmt.Println("open listen")
	//开启一个管理员
	ManagerDefault = RunMangerServer()
	s.upgrader = &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	http.HandleFunc(s.WSPath,s.WsHandler)
	fmt.Println("open WsHandler")
	err = http.ListenAndServe(s.Addr,nil)
	fmt.Println("err",err)
	return err
}

func (s *Server) WsHandler (w http.ResponseWriter, r *http.Request) {
	header := r.Header
	if header == nil || header["Sec-Websocket-Protocol"] == nil || header["Sec-Websocket-Protocol"][0] == "" {
		fmt.Println("暂无权限操作")
		return
	}

	//token
	token := Token(header["Sec-Websocket-Protocol"][0])
	user,err := s.AuthToken(token)
	if err != nil {
		fmt.Println("暂无权限操作")
		return
	}

	//返回token
	response := &http.Response{
		Header:make(map[string][]string),
	}
	response.Header.Add("Sec-Websocket-Protocol",string(token))

	// 完成ws协议的握手操作
	//将http升级为websocket
	conn,err := s.upgrader.Upgrade(w,r,response.Header)
	if err != nil {
		http.NotFound(w,r)
	}

	//创建客户端
	c := NewClient(user,conn)
	//注册客户端
	ManagerDefault.Connect <- c
}