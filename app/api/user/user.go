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
	tipCode := make(map[string]string)
	tipCode["Account"] = tipsMap[TipNotInputAccount]
	tipCode["Password"] = tipsMap[TipNotInputPassword]
	tipCode["PasswordConfirm"] = tipsMap[TipNotInputPasswordConfirm]
	tipCode["VerifyCode"] = tipsMap[TipNotInputCode]
	req := &usersrv.RegisterRequest{}
	if err := api.VerifyParams(ctx, req, tipCode); err == nil {
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
	tipCode := make(map[string]string)
	tipCode["Account"] = tipsMap[TipNotInputAccount]
	tipCode["Password"] = tipsMap[TipNotInputPassword]
	req := &usersrv.LoginRequest{}
	if err := api.VerifyParams(ctx, req, tipCode); err == nil {
		if err, user := usersrv.Login(req); err == nil {
			res.Success(ctx, user)
		} else {
			res.Error(ctx, err)
		}
	} else {
		res.Error(ctx, err)
	}
}
