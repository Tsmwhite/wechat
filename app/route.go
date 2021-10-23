package app

import (
	"github.com/gin-gonic/gin"
	"wechat/http/api"
	"wechat/http/middleware"
)

var ginEngine *gin.Engine

func init() {
	ginEngine = gin.Default()
	ginEngine.Use(middleware.AccessLog)

	ginEngine.POST("/register", api.Register)
	ginEngine.POST("/login", api.Login)

	ginEngine.Group("/chat", middleware.Authorization)
	{

	}
}

func Run() {
	ginEngine.Run(":8099")
}
