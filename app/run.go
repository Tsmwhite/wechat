package app

import (
	"wechat/app/route"
	"wechat/config"
	"wechat/core/log"
)

func Run() {
	if err := config.SetupServer(); err != nil {
		log.PrintlnErr("setup error:", err)
	}
	if err := route.Run(); err != nil {
		log.PrintlnErr("run error:", err)
	}
}
