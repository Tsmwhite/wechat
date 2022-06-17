package main

import (
	"fmt"
	"runtime"
	"wechat/config"
	"wechat/core/encrypt/token"
	"wechat/core/log"
	"wechat/core/server/wservice"
	model "wechat/models"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	err := config.SetupServer()
	if err != nil {
		log.Error.Println("setup error:", err)
	}
	option := wservice.NewOption()
	option.CheckToken = func(s string) (ok bool, uuid string) {
		user := model.NewUser()
		if token.Token(s).Check(user) {
			ok = true
			uuid = user.GetUuid()
		}
		fmt.Println("ok", ok, uuid)
		return
	}
	wservice.Run(option)
}
