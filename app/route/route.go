package route

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api/contact"
	"wechat/app/api/user"
	"wechat/app/middleware"
)

var ginEngine *gin.Engine

func init() {
	ginEngine = gin.Default()
	ginEngine.Use(middleware.AccessLog)

	ginEngine.POST("/register", user.Register)
	ginEngine.POST("/login", user.Login)

	chatApi := ginEngine.Group("/chat", middleware.Authorization)
	{
		chatApi.POST("/addFriends", contact.AddFriend)
		chatApi.POST("/friends", contact.FriendsList)
		chatApi.POST("/addContact", contact.AddContact)
		chatApi.POST("/contacts", contact.ContactsList)
	}
}

func Run() error {
	return ginEngine.Run(":8099")
}
