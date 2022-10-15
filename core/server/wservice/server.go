package wservice

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"wechat/config"
	wsClient "wechat/core/client"
	"wechat/core/log"
	"wechat/core/redis"
	"wechat/env"
	model "wechat/models"
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

var _option *Option
var _manager Manager
var _config *Config

func NewOption() *Option {
	return &Option{
		ClientManage: NewManager(),
		Config: &Config{
			router: config.ServerEnv.Router,
			port:   config.ServerEnv.Port,
		},
		CheckToken: func(s string) (ok bool, uuid string) {
			return
		},
		CheckOrigin: func(request *http.Request) bool {
			return true
		},
	}
}

func Run(option *Option) {
	defer func() {
		if err := recover(); err != nil {
			log.Error.Println("Run Panic:", err)
		}
	}()
	_option = option
	_manager = option.ClientManage
	_config = option.Config
	go _manager.Run()
	go addFriendsRequestHandel()
	//注册默认路由为 /ws ，并使用wsHandler这个方法
	http.HandleFunc(_config.router, wsHandler)
	//监听端口
	fmt.Println("Listen:", _config.port, _config.router)
	if err := http.ListenAndServe(_config.port, nil); err != nil {
		log.Error.Println("ListenAndServe Error:", err)
	}
}

func addFriendsRequestHandel() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error.Println("addFriendsRequestHandel recover error", err)
				}
			}()
			count := redis.LLen(env.AddFriendRequestHandel)
			if count < 1 {
				return
			}
			msg := &model.Message{}
			err := redis.LPop(env.AddFriendRequestHandel, msg)
			if err != nil {
				log.Error.Println("addFriendsRequestHandel error", err)
				return
			}
			if msg.Id > 0 {
				_manager.Broadcast(msg)
			}
		}()
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
	client := wsClient.NewClient(conn, _uuid, &model.Message{}, _manager.GetBroadcastChan(), _manager.GetUnregisterChan())
	//注册一个新的链接
	_manager.Register(client)
}
