package services

import (
	"wechat/app/res"
	"wechat/lib/encrypt"
	"wechat/lib/encrypt/token"
	model "wechat/models"
)

type LoginRequest struct {
	Account, Password string
}

type RegisterRequest struct {
	LoginRequest
	PasswordConfirm string
	VerifyCode      string
}

func Login(req *LoginRequest) (error, *model.ShowAppUser) {
	user := model.GetUserByAccount(req.Account)
	if !user.Exist() {
		return res.GetTip(res.TipNotExistAccount), nil
	}
	psMd5 := encrypt.PasswordMd5(req.Password, user.Salt)
	if psMd5 != user.Password {
		return res.GetTip(res.TipPasswordErr), nil
	}
	// ok
	data := user.ShowAppUser
	t := token.New(user.GetUuid())
	data.Token = string(t)
	return nil, &user.ShowAppUser
}

func Register(req *RegisterRequest) (error, *model.ShowAppUser) {

	return nil, nil
}
