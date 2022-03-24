package middleware

import (
	"github.com/gin-gonic/gin"
	"wechat/app/res"
	"wechat/core/encrypt/token"
	"wechat/env"
	model "wechat/models"
)

func Authorization(ctx *gin.Context) {
	authToken := token.Token(ctx.Request.Header.Get(env.AppAuthTokenKey))
	user := &model.User{}
	if authToken.Check(user) {
		ctx.Set(env.AppAuthTokenKey, authToken)
		ctx.Set(env.AppCurrentUserKey, user)
		ctx.Next()
	} else {
		//ctx.Data(http.StatusForbidden, "text/html", []byte("<h1>403 Forbidden</h1>"))
		res.ErrorStr(ctx, env.AppTokenError, env.AppTokenErrorCode)
		ctx.Abort()
	}
}
