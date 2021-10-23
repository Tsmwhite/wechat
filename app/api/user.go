package api

import (
	"github.com/gin-gonic/gin"
	"wechat/app/res"
	"wechat/services"
)

func Register(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {
	var request = ctx.Request.PostForm
	account := request.Get("account")
	password := request.Get("password")
	if account == "" {
		res.ErrCode(ctx, res.TipNotInputAccount)
	}
	if password == "" {
		res.ErrCode(ctx, res.TipNotInputPassword)
	}
	req := &services.LoginRequest{
		Account:  account,
		Password: password,
	}
	if err, data := services.LoginVerify(req); err != nil {
		res.Success(ctx, data)
	} else {
		res.Error(ctx, err)
	}
}
