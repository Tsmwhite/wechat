package redis

import (
	"time"
	"wechat/core/log"
)

func Get(key, uuid string, dest interface{}) {
	key = key + uuid
	err := Init().Get(Ctx, key).Scan(dest)
	if err != nil {
		log.Error.Println("Redis Get Key "+key+" Error:", err)
	}
}

func Set(key, uuid string, value interface{}, expiredTime ...time.Duration) {
	expired := time.Minute * 5
	if len(expiredTime) > 0 {
		expired = expiredTime[0]
	}
	key = key + uuid
	err := Init().Set(Ctx, key, value, expired).Err()
	if err != nil {
		log.Error.Println("Redis Set Key "+key+" Error:", err)
	}
}
