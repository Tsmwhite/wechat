package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/app/api/contact"
	"wechat/app/api/message"
	"wechat/app/api/user"
	"wechat/app/middleware"
	"wechat/config"
)

var ginEngine *gin.Engine

func _init() {
	ginEngine = gin.Default()
	ginEngine.Use(middleware.AccessLog)

	//web入口 home
	ginEngine.Static("/static", "./web/dist/static")
	ginEngine.Static("/.well-known", "./web/public/.well-known")
	ginEngine.LoadHTMLFiles("./web/dist/index.html")
	ginEngine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	ginEngine.POST("/register", user.Register)
	ginEngine.POST("/login", user.Login)
	ginEngine.POST("/loginReg", user.LoginOrRegister)
	ginEngine.POST("/sendCode", user.SendVerifyCode)
	ginEngine.GET("/getUserInfo", middleware.Authorization, user.GetUserInfo)

	chatApi := ginEngine.Group("/chat", middleware.Authorization)
	{
		chatApi.POST("/addFriends", contact.AddFriend)
		chatApi.POST("/addFriendsHandle", contact.AddFriendsHandle)
		chatApi.POST("/friends", contact.FriendsList)
		chatApi.POST("/addContact", contact.AddContact)
		chatApi.POST("/contacts", contact.ContactsList)
		chatApi.POST("/getContactInfo", contact.GetContactInfo)
		chatApi.POST("/searchUser", contact.SearchUser)
		chatApi.POST("/searchRoom", contact.SearchRoom)
		chatApi.POST("/searchFriend", contact.SearchFriends)
		chatApi.POST("/getFriendInfo", contact.GetFriendInfo)

		chatApi.POST("/createGroup", contact.CreateGroup)
		chatApi.POST("/inviteJoinGroup", contact.InviteJoin)
		chatApi.POST("/applyJoinGroup", contact.ApplyJoin)
		chatApi.POST("/joinGroupHandle", contact.JoinGroupHandle)

		chatApi.POST("/getHistory", message.GetHistory)
		chatApi.POST("/getFriendNotice", message.GetFriendNotice)
	}
}

func Run() error {
	if config.IsPro() {
		gin.SetMode(gin.ReleaseMode)
	}
	_init()
	if config.WebSrvEnv.Tls {
		return ginEngine.RunTLS(":"+config.WebSrvEnv.Port, config.WebSrvEnv.CertFile, config.WebSrvEnv.KeyFile)
	} else {
		fmt.Println("Listen:", config.WebSrvEnv.Port)
	}
	return ginEngine.Run(":" + config.WebSrvEnv.Port)
}
