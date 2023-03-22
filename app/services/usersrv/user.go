package usersrv

import (
	"errors"
	"wechat/app/res"
	"wechat/core/encrypt"
	"wechat/core/encrypt/token"
	"wechat/helper/format"
	model "wechat/models"
)

type LoginRequest struct {
	Account    string
	Password   string `validate:"noRequired & trim"`
	VerifyCode string `validate:"noRequired & trim"`
}

type RegisterRequest struct {
	Account         string `validate:"isMobile"`
	Password        string `validate:"trim & maxLen=20"`
	PasswordConfirm string `validate:"trim"`
	VerifyCode      string `validate:"trim"`
}

type LoginOrRegisterRequest struct {
	Account    string
	VerifyCode string
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
	data := user.ShowAppUser()
	data.Token = string(token.New(user.GetUuid()))
	return nil, data
}

func Register(req *RegisterRequest) (error, *model.ShowAppUser) {
	if model.GetUserByAccount(req.Account).Exist() {
		return res.GetTip(res.TipAccountNameExist), nil
	}
	user := model.NewUser()
	user.Mobile = req.Account
	user.PassLook = req.Password
	err := user.Create()
	if err != nil {
		return err, nil
	}
	data := user.ShowAppUser()
	data.Token = string(token.New(user.GetUuid()))
	return nil, data
}

func LoginOrRegister(req *LoginOrRegisterRequest) (error, *model.ShowAppUser) {
	user := model.GetUserByAccount(req.Account)
	if !user.Exist() {
		user = model.NewUser()
		user.Uuid = encrypt.CreateUuid()
		if format.IsMobile(req.Account) {
			user.Mobile = req.Account
		} else if format.IsMail(req.Account) {
			user.Mail = req.Account
		} else {
			return errors.New("账户格式不正确"), nil
		}
		user.PassLook = "123456"
		err := user.Create()
		if err != nil {
			return err, nil
		}
	}
	data := user.ShowAppUser()
	data.Token = string(token.New(user.GetUuid()))
	return nil, data
}

