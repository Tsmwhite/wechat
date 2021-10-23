package res

import "errors"

const (
	TipNotInputAccount = iota + 100
	TipNotInputPassword
	TipNotExistAccount
	TipPasswordErr
	TipAccountOrPassErr
)

var tipsMap map[int]string

func init()  {
	tipsMap = make(map[int]string)
	tipsMap[TipNotInputAccount] = "请输入登录账户"
	tipsMap[TipNotInputPassword] = "请输入账户密码"
	tipsMap[TipNotExistAccount] = "账户不存在"
	tipsMap[TipPasswordErr] = "密码不正确"
	tipsMap[TipAccountOrPassErr] = "账户或密码不正确"
}

func GetTip(code int) error {
	return errors.New(tipsMap[code])
}
