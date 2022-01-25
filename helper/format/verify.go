package format

import "regexp"

const (
	MobileReg = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	MailReg   = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
)

func IsMobile(mobile string) bool {
	reg := regexp.MustCompile(MobileReg)
	return reg.MatchString(mobile)
}

func IsMail(mail string) bool {
	reg := regexp.MustCompile(MailReg)
	return reg.MatchString(mail)
}
