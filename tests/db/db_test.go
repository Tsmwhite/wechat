package db

import (
	"fmt"
	"testing"
	model "wechat/models"
)

func TestCondition(t *testing.T) {
	var users,users2 []model.User
	model.FindAll(&model.Condition{Table: "users"}, &users)
	fmt.Println("users", users)

	model.FindAll(&model.Condition{}, &users2)
	fmt.Println("users", users)

	result := []map[string]interface{}{}
	model.FindAll(&model.Condition{Table: "users",Limit: 3}, &result)
	fmt.Println("users", result, len(result))
}
