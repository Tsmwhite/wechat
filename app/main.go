package main

import "wechat/app/route"

func main() {
	if err := route.Run(); err != nil {
		panic(err)
	}
}
