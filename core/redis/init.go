package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"wechat/config"
	"wechat/core/log"
)

var rdb *redis.Client
var Ctx = context.Background()

func Init() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.RedisEnv.Addr,
			Password: config.RedisEnv.Password,
			DB:       0, // use default DB
		})
	}
	return rdb
}

func Lock(key string) bool {
	ok, err := rdb.SetNX(Ctx, "lock:"+key, 1, time.Minute).Result()
	if err != nil {
		log.Error.Println("redis lock error:" + err.Error())
	}
	return ok
}
