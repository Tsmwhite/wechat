package wservice

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"wechat/config"
	wsClient "wechat/lib/client"
)

type Config struct {
	router, port string
}

type Option struct {
	ClientManage Manager
	Config       *Config
	CheckToken   func(string) (ok bool, uuid string)
	CheckOrigin  func(*http.Request) bool
}

var DefaultOption = &Option{
	ClientManage: NewManager(),
	Config: &Config{
		router: config.CONFIG.Server.Router,
		port:   config.CONFIG.Server.Port,
	},
	CheckToken: func(s string) (ok bool, uuid string) {
		return
	},
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

var _option *Option
var _manager Manager
var _config = &Config{
	router: config.CONFIG.Server.Router,
	port:   config.CONFIG.Server.Port,
}

func Run(option *Option) {
	_option = option
	_manager = option.ClientManage
	_config = option.Config
	go _manager.Run()
	//注册默认路由为 /ws ，并使用wsHandler这个方法
	http.HandleFunc(_config.router, wsHandler)
	//监听端口
	fmt.Println("Listen:", _config.port, _config.router)
	if err := http.ListenAndServe(_config.port, nil); err != nil {
		panic(err)
	}
}

func wsHandler(res http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Sec-WebSocket-Protocol")
	if token == "" {
		return
	}
	var _uuid string
	var ok bool
	if ok, _uuid = _option.CheckToken(token); !ok {
		return
	}

	//将http协议升级成websocket协议
	resHeader := http.Header{}
	resHeader = make(map[string][]string)
	resHeader.Add("Sec-Websocket-Protocol", token)
	conn, err := (&websocket.Upgrader{CheckOrigin: _option.CheckOrigin}).Upgrade(res, req, resHeader)
	if err != nil {
		return
	}
	//每一次连接都会新开一个client
	client := wsClient.NewClient(conn, _uuid, _manager.GetBroadcastChan())
	//注册一个新的链接
	_manager.Register(client)
}
