package main

import (
	"fmt"
	"wechat/lib/encrypt/token"
	"wechat/lib/server/wservice"
	model "wechat/models"
)

func main() {
	//fmt.Println(token.New("65dcd20109ab4458b1ec6e859f8de6c6"))
	option := wservice.DefaultOption
	option.CheckToken = func(s string) (ok bool, uuid string) {
		user := model.NewUser()
		if token.Token(s).Check(user) {
			ok = true
			uuid = user.GetUuid()
		}
		fmt.Println("ok", ok, uuid)
		return
	}
	wservice.RunServer(option)
}
