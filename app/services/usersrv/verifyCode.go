package usersrv

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"wechat/helper/format"
	"wechat/helper/mail"
	model "wechat/models"
)

type SendCodeRequest struct {
	Account string
	Type    int
}

func SendCode(req *SendCodeRequest) error {
	now := time.Now().Unix()
	code := createCode(6)
	res := model.DB.Create(&model.VerifyCode{
		Account: req.Account,
		Code:    code,
		Type:    req.Type,
		CreateTime: now,
		UpdateTime: now,
		ExpireTime: now + 5 * 60,
	})
	if res.Error != nil {
		return res.Error
	}
	switch {
	case format.IsMail(req.Account):
		body := fmt.Sprintf("<div><label>Verification code:</label> <span style='color:#c60023;font-weight:500;'>%s</div><div>Valid for five minutes</div>", code)
		if err := mail.New().SetMessage(req.Account, "TsmWhite Chat Verification Code", body).Send(); err != nil {
			return err
		}
	case format.IsMobile(req.Account):
		//TODO
	}
	return nil
}

func VerifyCode(account, code string) error {
	codeInfo := &model.VerifyCode{}
	model.DB.Raw("SELECT * FROM `verify_codes` WHERE `account` = ? AND `code` = ? ORDER BY `id` DESC LIMIT 1", account, code).Scan(codeInfo)
	if codeInfo.Id == 0 {
		return errors.New("验证码错误")
	}
	if codeInfo.Status != model.VerifyCodeStatusActive {
		return errors.New("验证码已失效")
	}
	now := time.Now().Unix()
	if now > codeInfo.ExpireTime {
		return errors.New("验证码已过期")
	}
	model.DB.Table("verify_codes").Updates(map[string]interface{}{
		"status":      model.VerifyCodeStatusUsed,
		"update_time": now,
	})
	return nil
}

func createCode(length int) (code string) {
	str := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	max := len(str)
	var key int
	for i := 0; i < length; i++ {
		randS := rand.New(rand.NewSource(time.Now().UnixNano()))
		key = randS.Intn(max)
		code += str[key]
	}
	return code
}
