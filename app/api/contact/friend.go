package contact

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services/contactsrv"
)

func AddFriend(ctx *gin.Context) {

}

func FriendsList(ctx *gin.Context) {
	req := &contactsrv.FriendsRequest{}
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	list := contactsrv.GetFriends(req, api.GetCurrentUser(ctx))
	res.Success(ctx, list)
}
