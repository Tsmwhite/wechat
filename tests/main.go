package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wechat/app/services/contactsrv"
	"wechat/config"
	"wechat/core/encrypt"
	model "wechat/models"
)

func main() {
	initDB()
	createOverallGroup()
	//socket.Run()
}

func initDB() {
	config.DBEnv.Dsn = "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
}

func ReqChatGPT(question string) {
	request, err := http.NewRequest(
		"POST",
		"https://api.openai.com/v1/completions",
		strings.NewReader(`{"model": "text-davinci-003", "prompt": "`+question+`", "temperature": 0, "max_tokens": 2000}`))
	if err != nil {
		fmt.Println("Request Error:", err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer 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")
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response ReadAll Error:", err)
	}
	fmt.Println("Answer:", string(body))
}

func writeUuid() {
	config.DBEnv.Dsn = "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
	var usersCount int64
	model.DB().Table("users").Count(&usersCount)
	for i := 1; i <= int(usersCount); i++ {
		err := model.DB().Table("users").Where("id = ?", i).Update("uuid", encrypt.CreateUuid()).Error
		if err != nil {
			fmt.Println("err", err)
			time.Sleep(time.Second * 2)
			i--
		}
	}
}

func insertUsers() {
	names := getNames()
	namesLen := len(names)
	if namesLen < 1 {
		return
	}
	config.DBEnv.Dsn = "root:lxy196914@tcp(127.0.0.1:3306)/thewhite?charset=utf8&multiStatements=true"
	var users []*model.User
	var insert = func(users []*model.User) {
		err := model.DB().Create(&users).Error
		if err != nil {
			fmt.Println("Create error", err)
		}
	}
	for i := 288289; i < namesLen; i++ {
		name := names[i]
		user := new(model.User)
		user.Name = name
		user.PassLook = "123456"
		user.Mail = "testMail" + strconv.Itoa(i) + "@163.com"
		users = append(users, user)
		if len(users) > 1000 {
			insert(users)
			users = []*model.User{}
			time.Sleep(2 * time.Second)
		}
	}
	insert(users)
	fmt.Println("ok")
}

func getNames() (names []string) {
	res, err := ioutil.ReadFile("./user/names.json")
	if err != nil {
		fmt.Println("ReadFile Error:", err)
		return
	}
	err = json.Unmarshal(res, &names)
	if err != nil {
		fmt.Println("Json Unmarshal Error:", err)
		return
	}
	return
}

func createOverallGroup() {
	creator := new(model.User)
	model.DB().Table("users").Where("id = 1").First(&creator)
	var members []*model.User
	model.DB().Table("users").Where("id > 1").Limit(70).Find(&members)
	fmt.Println(members)
	var users []string
	for _, mem := range members {
		users = append(users, mem.Uuid)
	}
	fmt.Println("users",len(users))
	err := contactsrv.CreateGroup(&contactsrv.CreateGroupRequest{
		Title:   "尬聊啊！！！",
		Friends: strings.Join(users, ","),
	}, creator)
	println("err", err)
}
