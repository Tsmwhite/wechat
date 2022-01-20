package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func VerifyParams(ctx *gin.Context, paramKeys []string, tips map[string]string) (error, map[string]string) {
	resParams := make(map[string]string)
	temp := ""
	for _, key := range paramKeys {
		temp = ctx.PostForm(key)
		if temp != "" {
			resParams[key] = temp
		} else {
			return errors.New(tips[key]), nil
		}
	}
	return nil, resParams
}
