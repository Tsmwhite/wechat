package api

import (
	"github.com/gin-gonic/gin"
	"wechat/app/res"
)

func VerifyParams(ctx *gin.Context, paramKeys []string, tips map[string]int) (bool, map[string]string) {
	params := ctx.Params
	resParams := make(map[string]string)
	ok := false
	temp := ""
	for _, key := range paramKeys {
		temp, ok = params.Get(key)
		if ok {
			resParams[key] = temp
		} else {
			res.ErrCode(ctx, tips[key])
			return false, nil
		}
	}
	return true, resParams
}
