package user

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
	"wechat/app/services/contactsrv"
	"wechat/config"
	"wechat/core/encrypt/token"
	"wechat/core/message"
	model "wechat/models"
)

var account = flag.String("acc", "", "用户账户")
var password = flag.String("pwd", "123456", "用户密码")
var uuid = flag.String("uuid", "", "uuid")

func TestGetToken(t *testing.T) {
	flag.Parse()
	fmt.Println("uuid", *uuid)
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

func TestMergeNames(t *testing.T) {
	nameMap := make(map[string]bool)
	for i := 1; i < 4; i++ {
		res, err := ioutil.ReadFile("names0" + strconv.Itoa(i) + ".json")
		if err != nil {
			fmt.Println("ReadFile Error:", err)
			continue
		}
		var names []string
		err = json.Unmarshal(res, &names)
		if err != nil {
			fmt.Println("Json Unmarshal Error:", err)
			continue
		}
		for _, name := range names {
			nameMap[name] = true
		}
	}
	var names []string
	for name := range nameMap {
		names = append(names, name)
	}
	fmt.Println("names len", len(names))
	namesJson, err := json.Marshal(names)
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}
	err = ioutil.WriteFile("names.json", namesJson, os.FileMode(0755))
	if err != nil {
		fmt.Println("WriteFile Error:", err)
		return
	}
}

func TestRandomNames(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wait := sync.WaitGroup{}
	nameMap := make(map[string]bool)
	nameWriteCh := make(chan string, 1000)
	go func() {
		for name := range nameWriteCh {
			if name != "" {
				nameMap[name] = true
			}
		}
	}()
	for i := 0; i < 10000; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			res, err := http.Get("https://www.qmsjmfb.com/")
			if err != nil {
				fmt.Println("Request Error:", err)
				return
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				fmt.Println("status code error: ", res.StatusCode, res.Status)
			}
			doc, err := goquery.NewDocumentFromReader(res.Body)
			doc.Find(".name_box.container ul li").Each(func(i int, s *goquery.Selection) {
				nameWriteCh <- s.Text()
			})
		}()
	}
	wait.Wait()
	for len(nameWriteCh) > 0 {
	}
	var names []string
	for name := range nameMap {
		names = append(names, name)
	}
	fmt.Println("names len", len(names))
	namesJson, err := json.Marshal(names)
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}
	err = ioutil.WriteFile("names.json", namesJson, os.FileMode(0755))
	if err != nil {
		fmt.Println("WriteFile Error:", err)
		return
	}
}

func TestMaxLoginReq(t *testing.T) {
	config.DBEnv.Dsn = "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
	var users []*model.User
	model.Find(&model.Condition{
		Table: "users",
		Limit: 100000,
	}, &users)

	okCounter := make(chan int)
	wait := sync.WaitGroup{}
	for i := 0; i < len(users); i++ {
		wait.Add(1)
		go func(index int) {
			defer wait.Done()
			user := users[index]
			res, err := http.PostForm(
				"http://127.0.0.1:8099/login",
				url.Values{
					"account":  {user.Mail},
					"password": {user.PassLook},
				},
			)
			if err != nil {
				fmt.Println("http Error:", err)
				return
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("body Read Error:", err)
				return
			}
			data := make(map[string]interface{})
			err = json.Unmarshal(body, &data)
			resStatus := data["ok"].(bool)
			if !resStatus {
				fmt.Println("Request Error:", data["err_msg"])
				return
			}
			okCounter <- 1
		}(i)
	}
	okC := 0
	go func() {
		for range okCounter {
			okC += 1
		}
	}()
	wait.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("okC", okC)
}

func TestCreateFriendRelation(t *testing.T) {
	config.DBEnv.Dsn = "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
	user := model.GetUserByAccount("577689878@qq.com")
	friends := []string{
		"thesmallwhiteme@163.com",
		//"577689878@qq.com",
		"thesmallwhiteme@gmail.com",
		"testMail3@163.com",
		"testMail4@163.com",
		"testMail5@163.com",
		"testMail6@163.com",
		"testMail7@163.com",
		"testMail8@163.com",
		"testMail9@163.com",
		"testMail10@163.com",
		"testMail11@163.com",
		"testMail12@163.com",
		"testMail13@163.com",
		"testMail14@163.com",
		"testMail15@163.com",
		"testMail16@163.com",
		"testMail17@163.com",
	}
	//发送好友申请
	for _, account := range friends {
		u := model.GetUserByAccount(account)
		err := contactsrv.AddFriend(&contactsrv.AddFriendRequest{
			Uuid:    user.Uuid,
			Content: "洛 Long time no see," + u.Name,
		}, u)
		if err != nil {
			t.Log("AddFriend Err", err)
		}
	}

	// 处理好友申请
	var applyList []model.Message
	model.Find(&model.Condition{
		Table: model.HandleMessageTable,
		Where: map[string]interface{}{
			"recipient": user.Uuid,
			"status":    message.AddFriendStatusNormal,
		},
	}, &applyList)

	for _, msg := range applyList {
		err := contactsrv.AddFriendHandle(&contactsrv.AddFriendHandleRequest{
			Uuid:   msg.Uuid,
			Status: message.AddFriendStatusAgree,
		}, user)
		if err != nil {
			t.Log("AddFriendHandle Err", err)
		}
	}
}

func TestGetUserUnreadMsgCount(t *testing.T) {
	config.DBEnv.Dsn = "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
	r := model.GetRoomBuyUuid("f272932fb0b5d52a4eec97d2d00b5d23")
	u := model.GetUserByUuid("91dc76fd6b8b4a7e8965fe5c7859db01")
	c := r.GetUnreadMsgCountByUser(u)
	fmt.Println("c", c)
}
