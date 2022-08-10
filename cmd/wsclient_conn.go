package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wechat/core/encrypt/token"
)

var addr = flag.String("addr", "localhost:8011", "http service address")
var user = flag.String("uuid", "731be9ef7b7e42feb884ee8faab4bdb1", "user uuid")
var source = flag.String("source", "message-source.txt", "消息源文件")

func main() {
	flag.Parse()
	connect()
}

func connect() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	var dialer *websocket.Dialer
	//通过Dialer连接websocket服务器
	conn, _, err := dialer.Dial(u.String(), http.Header{
		"Sec-WebSocket-Protocol": []string{token.New(*user).ToString()},
	})
	if err != nil {
		fmt.Println("Dial Error:", err)
		return
	}
	msgSource, err := ioutil.ReadFile("./static/file/" + *source)
	if err != nil {
		fmt.Println("ReadFile Error:", err)
		return
	}
	msgSourceStr := string(msgSource)
	msgSourceArr := strings.Split(msgSourceStr, "\n")
	for _, msg := range msgSourceArr {
		err := conn.WriteJSON(map[string]interface{}{
			"content":     msg,
			"recipient":   "712056f6ca8e441a9fc8a5d29bda4006",
			"second_type": 400,
			"send_time":   time.Now().Unix(),
			"sender":      *user,
			"type":        200,
		})
		if err != nil {
			fmt.Println("Conn WriteJSON:", err)
		}
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		fmt.Printf("received: %s\n", message)
	}
}
