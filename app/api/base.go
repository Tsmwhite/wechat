package api

import (
	"github.com/gin-gonic/gin"
	"wechat/app/utils/validations"
	"wechat/env"
	model "wechat/models"
)

func NewResult() []map[string]interface{} {
	return []map[string]interface{}{}
}

func GetCurrentUser(ctx *gin.Context) *model.User {
	user, _ := ctx.Get(env.AppCurrentUserKey)
	currentUser, ok := user.(*model.User)
	if !ok {
		return nil
	}
	return currentUser
}

func VerifyParams(ctx *gin.Context, carrier interface{}, messages map[string]string) error {
	return validations.Validate(carrier, messages, func(s string) string {
		return ctx.PostForm(s)
	})
}
