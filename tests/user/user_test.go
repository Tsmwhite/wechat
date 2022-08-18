package user

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"wechat/core/encrypt/token"
)

var account = flag.String("acc", "", "用户账户")
var password = flag.String("pwd", "123456", "用户密码")
var uuid =  flag.String("uuid", "", "uuid")

func TestGetToken(t *testing.T)  {
	flag.Parse()
	fmt.Println("uuid",*uuid)
	fmt.Println(token.New(*uuid))
}

func TestRegister(t *testing.T) {
	flag.Parse()
	if *account == "" {
		t.Error("缺少账户名")
	}
	if *password == "" {
		t.Error("缺少账户密码")
	}
	//account", "password", "passwordConfirm", "verifyCode
	res, err := http.PostForm(
		"http://127.0.0.1:8099/register",
		url.Values{
			"account":         {*account},
			"password":        {*password},
			"passwordConfirm": {*password},
			"verifyCode":      {"000000"}},
	)
	if err != nil {
		t.Error("http Error:", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("body Read Error:", err)
	}
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if data["err_code"] != "0" {
		t.Error("Request Error:", data["err_msg"])
	}
}

func TestLogin(t *testing.T) {
	flag.Parse()
	if *account == "" {
		t.Error("缺少账户名")
	}
	if *password == "" {
		t.Error("缺少账户密码")
	}
	res, err := http.PostForm(
		"http://127.0.0.1:8099/login",
		url.Values{
			"account":  {*account},
			"password": {*password},
		},
	)
	if err != nil {
		t.Error("http Error:", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("body Read Error:", err)
	}
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if data["err_code"] != "0" {
		t.Error("Request Error:", data["err_msg"])
	}
}
