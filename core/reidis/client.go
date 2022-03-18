package reidis

import (
	"github.com/go-redis/redis/v8"
)

var DB *redis.Client

const RedisAddr = "120.26.1.206:6379"
const RedisPassword = "wxigNQxrwMm3rECCDAJBdtoO7gxHTnwd"

func Connect() *redis.Client {
	if DB == nil {
		DB = redis.NewClient(&redis.Options{
			Addr:     RedisAddr,
			Password: RedisPassword,
			DB:       0,
		})
	}
	return DB
}
