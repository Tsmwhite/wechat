package db

import (
	"fmt"
	"hash/crc32"
	"math"
	"testing"
	model "wechat/models"
)

func TestCondition(t *testing.T) {
	var users, users2 []model.User
	model.FindAll(&model.Condition{Table: "users"}, &users)
	fmt.Println("users", users)

	model.FindAll(&model.Condition{}, &users2)
	fmt.Println("users", users)

	var result []map[string]interface{}
	model.FindAll(&model.Condition{Table: "users", Limit: 3}, &result)
	fmt.Println("users", result, len(result))
}

func TestHashCutTable(t *testing.T) {
	hashVal := crc32.ChecksumIEEE([]byte("712056f6ca8e441a9fc8a5d29bda4006"))
	val := int(hashVal)
	if val > 2147483647 {
		val -= 4294967296
	}
	fmt.Println("val:", math.Abs(float64(val%10)))
}
