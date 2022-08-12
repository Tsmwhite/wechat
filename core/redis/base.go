package redis

import (
	"encoding/json"
	"time"
	"wechat/core/log"
)

func Get(key, uuid string, dest interface{}) {
	key = key + uuid
	res := Init().Get(Ctx, key).Val()
	if res == "" {
		return
	}
	err := json.Unmarshal([]byte(res), dest)
	if err != nil {
		log.Error.Println("Redis Get Unmarshal Key "+key+" Error:", err)
	}
}

func Set(key, uuid string, value interface{}, expiredTime ...time.Duration) {
	expired := time.Minute * 5
	if len(expiredTime) > 0 {
		expired = expiredTime[0]
	}
	key = key + uuid
	valByte, err := json.Marshal(value)
	err = Init().Set(Ctx, key, valByte, expired).Err()
	if err != nil {
		log.Error.Println("Redis Set Key "+key+" Error:", err)
	}
}
