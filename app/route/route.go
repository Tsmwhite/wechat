package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/app/api/contact"
	"wechat/app/api/user"
	"wechat/app/middleware"
	"wechat/config"
)

var ginEngine *gin.Engine

func init() {
	ginEngine = gin.Default()
	ginEngine.Use(middleware.AccessLog)

	//web入口 home
	ginEngine.Static("/static", "./web/dist/static")
	ginEngine.LoadHTMLFiles("./web/dist/index.html")
	ginEngine.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	ginEngine.POST("/register", user.Register)
	ginEngine.POST("/login", user.Login)
	ginEngine.POST("/loginReg", user.LoginOrRegister)
	ginEngine.POST("/sendCode", user.SendVerifyCode)

	chatApi := ginEngine.Group("/chat", middleware.Authorization)
	{
		chatApi.POST("/addFriends", contact.AddFriend)
		chatApi.POST("/addFriendsHandle", contact.AddFriendsHandle)
		chatApi.POST("/friends", contact.FriendsList)
		chatApi.POST("/addContact", contact.AddContact)
		chatApi.POST("/contacts", contact.ContactsList)
		chatApi.POST("/searchUser", contact.SearchUser)
		chatApi.POST("/searchRoom", contact.SearchRoom)
		chatApi.POST("/searchFriend", contact.SearchFriends)

		chatApi.POST("/createGroup", contact.CreateGroup)
	}
}

func Run() error {
	fmt.Println("Listen:", config.WebSrvEnv.Port)
	return ginEngine.Run(":" + config.WebSrvEnv.Port)
}
