package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

var Ctx  = context.Background()

func Init() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return rdb
}
