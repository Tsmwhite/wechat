package route

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api/user"
	"wechat/app/middleware"
)

var ginEngine *gin.Engine

func init() {
	ginEngine = gin.Default()
	ginEngine.Use(middleware.AccessLog)

	ginEngine.POST("/register", user.Register)
	ginEngine.POST("/login", user.Login)

	ginEngine.Group("/chat", middleware.Authorization)
	{

	}
}

func Run() error {
	return ginEngine.Run(":8099")
}
