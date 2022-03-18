package middleware

import (
	"github.com/gin-gonic/gin"
	"wechat/app/res"
	"wechat/core/encrypt/token"
	model "wechat/models"
)

const AuthTokenKey = "authorization"
const CurrentUserKey = "current-user"
const TokenError = "token错误"
const TokenErrorCode = 10004

func Authorization(ctx *gin.Context) {
	authToken := token.Token(ctx.Request.Header.Get(AuthTokenKey))
	user := &model.User{}
	if authToken.Check(user) {
		ctx.Set(AuthTokenKey, authToken)
		ctx.Set(CurrentUserKey, user)
		ctx.Next()
	} else {
		res.ErrorStr(ctx, TokenError, TokenErrorCode)
		ctx.Abort()
	}
}
