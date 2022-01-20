package user

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services/usersrv"
)

const (
	TipNotInputAccount = iota + 100
	TipNotInputPassword
	TipNotInputPasswordConfirm
	TipNotInputCode
	TipNotExistAccount
	TipPasswordErr
	TipAccountOrPassErr
)

var tipsMap = map[int]string{
	TipNotInputAccount:         "请输入登录账户",
	TipNotInputPassword:        "请输入账户密码",
	TipNotInputPasswordConfirm: "请确认密码",
	TipNotInputCode:            "请输入验证码",
	TipNotExistAccount:         "账户不存在",
	TipPasswordErr:             "密码不正确",
	TipAccountOrPassErr:        "账户或密码不正确",
}

func Register(ctx *gin.Context) {
	params := []string{"account", "password", "passwordConfirm", "verifyCode"}
	tipCode := make(map[string]string)
	tipCode["account"] = tipsMap[TipNotInputAccount]
	tipCode["password"] = tipsMap[TipNotInputPassword]
	tipCode["passwordConfirm"] = tipsMap[TipNotInputPasswordConfirm]
	tipCode["verifyCode"] = tipsMap[TipNotInputCode]
	if err, data := api.VerifyParams(ctx, params, tipCode); err == nil {
		req := &usersrv.RegisterRequest{}
		req.Account = data["account"]
		req.Password = data["password"]
		req.PasswordConfirm = data["passwordConfirm"]
		req.VerifyCode = data["verifyCode"]
		if err, user := usersrv.Register(req); err == nil {
			res.Success(ctx, user)
		} else {
			res.Error(ctx, err)
		}
	} else {
		res.Error(ctx, err)
	}
}

func Login(ctx *gin.Context) {
	params := []string{"account", "password", "passwordConfirm", "verifyCode"}
	tipCode := make(map[string]string)
	tipCode["account"] = tipsMap[TipNotInputAccount]
	tipCode["password"] = tipsMap[TipNotInputPassword]
	if err, data := api.VerifyParams(ctx, params, tipCode); err == nil {
		req := &usersrv.LoginRequest{
			Account:  data["account"],
			Password: data["password"],
		}
		if err, user := usersrv.Login(req); err == nil {
			res.Success(ctx, user)
		} else {
			res.Error(ctx, err)
		}
	} else {
		res.Error(ctx, err)
	}
}
