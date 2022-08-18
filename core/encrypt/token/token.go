package token

import (
	"fmt"
	"strings"
	"time"
	"wechat/core/encrypt"
	"wechat/core/member"
)

const tokenKey = "c2db2d1e0d8d0753d8878818da213721"
const tokenTimeout = 10 * 60

type Token string // uuid-expiredTime-updateTime

func New(uid string) Token {
	expiredTime := time.Now().Unix() + tokenTimeout
	updateTime := expiredTime - (2 * 60)
	text := fmt.Sprintf("%s-%d-%d", uid, expiredTime, updateTime)
	ciphertext, _ := encrypt.Encrypt(text, []byte(tokenKey))
	return Token(ciphertext)
}

func (token Token) Decode() []string {
	text, err := encrypt.Decrypt(string(token), []byte(tokenKey))
	if err != nil {
		return []string{}
	}
	values := strings.Split(text, "-")
	return values
}

func (token Token) GetUuid() string {
	values := token.Decode()
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func (token Token) Check(member member.Member) bool {
	_uuid := token.GetUuid()
	if _uuid == "" {
		return false
	}
	return member.CheckMemberByUuid(_uuid)
}

func (token Token) ToString() string {
	return string(token)
}
