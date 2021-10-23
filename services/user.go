package services

import (
	"wechat/app/res"
	"wechat/lib/encrypt"
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

func LoginVerify(req *LoginRequest) (error, *model.User) {
	user := model.GetUserByAccount(req.Account)
	if !user.Exist() {
		return res.GetTip(res.TipNotExistAccount), nil
	}
	psMd5 := encrypt.PasswordMd5(req.Password, user.Salt)
	if psMd5 != user.Password {
		return res.GetTip(res.TipPasswordErr), nil
	}
	// ok
	return nil, user
}

func RegisterVerify(req *LoginRequest) (bool, error) {

	return true, nil
}
