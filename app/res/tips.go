package res

import "errors"

const (
	TipNotInputAccount = iota + 100
	TipNotInputPassword
	TipNotInputPasswordConfirm
	TipNotInputCode
	TipNotExistAccount
	TipPasswordErr
	TipAccountOrPassErr
	TipAccountNameExist
)

var tipsMap = map[int]string{
	TipNotInputAccount:         "请输入登录账户",
	TipNotInputPassword:        "请输入账户密码",
	TipNotInputPasswordConfirm: "请确认密码",
	TipNotInputCode:            "请输入验证码",
	TipNotExistAccount:         "账户不存在",
	TipPasswordErr:             "密码不正确",
	TipAccountOrPassErr:        "账户或密码不正确",
	TipAccountNameExist:        "账户名已存在",
}

func GetTip(code int) error {
	return errors.New(tipsMap[code])
}
