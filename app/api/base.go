package api

import (
	"github.com/gin-gonic/gin"
	"wechat/app/utils/validations"
)

func VerifyParams(ctx *gin.Context, carrier interface{}, messages map[string]string) error {
	return validations.Validate(carrier,messages, func(s string) string {
		return ctx.PostForm(s)
	})
}
