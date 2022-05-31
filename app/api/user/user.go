package user

import (
	"errors"
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
	TipPasswordConfirmErr
)

var tipsMap = map[int]string{
	TipNotInputAccount:         "请输入登录账户",
	TipNotInputPassword:        "请输入账户密码",
	TipNotInputPasswordConfirm: "请确认密码",
	TipNotInputCode:            "请输入验证码",
	TipNotExistAccount:         "账户不存在",
	TipPasswordErr:             "密码不正确",
	TipAccountOrPassErr:        "账户或密码不正确",
	TipPasswordConfirmErr:      "密码确认不一致",
}

func Register(ctx *gin.Context) {
	tipCode := make(map[string]string)
	tipCode["Account"] = tipsMap[TipNotInputAccount]
	tipCode["Password"] = tipsMap[TipNotInputPassword]
	tipCode["PasswordConfirm"] = tipsMap[TipNotInputPasswordConfirm]
	tipCode["VerifyCode"] = tipsMap[TipNotInputCode]
	req := &usersrv.RegisterRequest{}
	if err := api.VerifyParams(ctx, req, tipCode); err == nil {
		if req.Password != req.PasswordConfirm {
			res.Error(ctx, errors.New(tipsMap[TipPasswordConfirmErr]))
			return
		}
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

func LoginOrRegister(ctx *gin.Context) {
	tipCode := make(map[string]string)
	tipCode["Account"] = tipsMap[TipNotInputAccount]
	tipCode["VerifyCode"] = tipsMap[TipNotInputCode]
	req := &usersrv.LoginOrRegisterRequest{}
	err := api.VerifyParams(ctx, req, tipCode)
	if err != nil {
		res.Error(ctx, err)
	} else {
		if err = usersrv.VerifyCode(req.Account, req.VerifyCode); err != nil {
			res.Error(ctx, err)
		} else if err, user := usersrv.LoginOrRegister(req); err == nil {
			res.Success(ctx, user)
		} else {
			res.Error(ctx, err)
		}
	}
}

func SendVerifyCode(ctx *gin.Context) {
	req := &usersrv.SendCodeRequest{}
	err := api.VerifyParams(ctx, req, nil)
	if err != nil {
		res.Error(ctx, err)
	} else {
		if err := usersrv.SendCode(req); err == nil {
			res.Success(ctx, res.Nil, "已发送，请注意查收")
		} else {
			res.Error(ctx, err)
		}
	}
}
