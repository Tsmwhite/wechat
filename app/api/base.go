package api

import (
	"github.com/gin-gonic/gin"
	"wechat/app/utils/validations"
	env "wechat/env"
	model "wechat/models"
)

var currentUser *model.User

func NewResult() []map[string]interface{} {
	return []map[string]interface{}{}
}

func GetCurrentUser(ctx *gin.Context) *model.User {
	if currentUser == nil {
		user, _ := ctx.Get(env.AppCurrentUserKey)
		currentUser = user.(*model.User)
	}
	return currentUser
}

func VerifyParams(ctx *gin.Context, carrier interface{}, messages map[string]string) error {
	return validations.Validate(carrier, messages, func(s string) string {
		return ctx.PostForm(s)
	})
}
