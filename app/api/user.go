package api

import (
	"github.com/gin-gonic/gin"
	"wechat/app/res"
	"wechat/services"
)

func Register(ctx *gin.Context) {
	params := []string{"account", "password", "passwordConfirm", "verifyCode"}
	tipCode := make(map[string]int)
	tipCode["account"] = res.TipNotInputAccount
	tipCode["password"] = res.TipNotInputPassword
	tipCode["passwordConfirm"] = res.TipNotInputPasswordConfirm
	tipCode["verifyCode"] = res.TipNotInputCode
	if ok, data := VerifyParams(ctx, params, tipCode); ok {
		req := &services.RegisterRequest{}
		req.Account = data["account"]
		req.Password = data["password"]
		req.PasswordConfirm = data["passwordConfirm"]
		req.VerifyCode = data["verifyCode"]
		if err, user := services.Register(req); err != nil {
			res.Success(ctx, user)
		} else {
			res.Error(ctx, err)
		}
	}
}

func Login(ctx *gin.Context) {
	params := []string{"account", "password", "passwordConfirm", "verifyCode"}
	tipCode := make(map[string]int)
	tipCode["account"] = res.TipNotInputAccount
	tipCode["password"] = res.TipNotInputPassword
	if ok, data := VerifyParams(ctx, params, tipCode); ok {
		req := &services.LoginRequest{
			Account:  data["account"],
			Password: data["password"],
		}
		if err, user := services.Login(req); err != nil {
			res.Success(ctx, user)
		} else {
			res.Error(ctx, err)
		}
	}
}
