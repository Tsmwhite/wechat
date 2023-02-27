package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"wechat/config"
	"wechat/core/encrypt"
	model "wechat/models"
)

func main() {
	question := bufio.NewScanner(os.Stdin)
	fmt.Println("what can i do for you")
	for question.Scan() {
		line := question.Text()
		ReqChatGPT(line)
	}
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
