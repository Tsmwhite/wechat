package main

import (
	"fmt"
	"wechat/socket"
)

func main() {
	fmt.Println("open")
	Start := &socket.Server{
		Addr:":8088",
		WSPath:"/ws",
		AuthToken: func(token socket.Token) (id int, err error) {
			fmt.Println(token)
			return
		},
	}
	err := Start.ListenServer()
	fmt.Println(err)
}

