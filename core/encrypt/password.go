package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// GetRandomChar 生成随机字符串
func GetRandomChar(length int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bytesLen := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

func PasswordMd5(password, salt string) string {
	return Md5(Md5(password) + salt)
}

func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}