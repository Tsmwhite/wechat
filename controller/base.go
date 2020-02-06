package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	//"strings"
	"time"

	"wechat/model"
)

const (
	UPLOAD_HEADIMG_PATH = "Uploads/Headimg/" //头像文件目录
)

var MemberId int
var LoginTime int
var MEMBER_INFO *model.Member

func GetClientIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

//返回错误信息
func ResError(res interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "0",
		"info":   res,
	})
	c.Abort()
}

//返回请求内容
func ResSucces(res interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "1",
		"info":   res,
	})
	c.Abort()
}

func GetSaLt() string {
	return GetRandomChar(8)
}

//生成随机字符串
func GetRandomChar(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	bytesLen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

// 计算字符串的md5值
func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}

//根据用户id当前时间戳生成用户密钥
func GetUserKey(member_id int) string {
	time := time.Now().Unix()
	time_str := strconv.FormatInt(time, 10)
	member_str := strconv.Itoa(member_id)
	orig := member_str + "_" + time_str

	key := []byte{0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
		0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
	}
	decrypt_str, err := Encrypt(orig, key)
	if err == nil {
		return decrypt_str
	}
	return ""
}

//解析用户密钥返回用户信息
//func ResUserKeyMsg(encrypted string) (error string){
//	key := []byte{0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
//		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
//		0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
//		0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
//	}
//
//	msg,err := Decrypt(encrypted,key)
//	if err == nil{
//		arr := strings.Split(msg,"_")
//		member_id,_ := strconv.Atoi(arr[0])
//		LoginTime,_  = strconv.Atoi(arr[1])
//		mem,err_ := model.GetMemberInfo("id",member_id)
//		if err_ != nil{
//			error = "账户不存在"
//		}
//		if mem.Id > 0{
//			MEMBER_INFO = mem
//			MemberId  = MEMBER_INFO.Id
//		}else{
//			error = "账户不存在"
//		}
//	}else{
//		error = "账户不存在"
//	}
//	return error
//}

//生成MD5密码
func PassWrodMd5(password, salt string) string {
	return Md5(Md5(password) + salt)
}

//获取请求ip
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("test"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("test"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

//加密
func Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

//解密
func Decrypt(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

//写入日志
func WriteLog(content string) {
	//当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	//截取年月日
	filename := "Log/" + string([]byte(nowTime)[:10]) + ".log"
	content = nowTime + "\r\n" + content
	if f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644); err != nil && os.IsNotExist(err) {
		if f, err := os.Create(filename); err == nil {
			FileWrite(f, content)
		}
	} else {
		FileWrite(f, content)
	}
}

//文件写入
func FileWrite(f *os.File, content string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	_, err := fmt.Fprintln(f, content)
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func Handle404(c *gin.Context) {
	ResError("Not Fund 404", c)
}
