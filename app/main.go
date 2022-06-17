package main

import (
	"wechat/app/route"
	"wechat/config"
	"wechat/core/log"
)

func main() {
	err := config.SetupServer()
	if err != nil {
		log.Error.Println("setup error:", err)
	}
	if err := route.Run(); err != nil {
		panic(err)
	}
}
