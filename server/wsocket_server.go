package main

import (
	"fmt"
	"wechat/core/encrypt/token"
	"wechat/core/server/wservice"
	model "wechat/models"
)

func main() {
	//user := model.NewUser()
	//model.DB.Raw("SELECT * FROM `users`").Scan(user)
	//fmt.Println("user", user)
	//return
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
	wservice.Run(option)
}
