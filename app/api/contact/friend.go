package contact

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services/contactsrv"
)

func AddFriend(ctx *gin.Context) {
	req := &contactsrv.AddFriendRequest{}
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	if err := contactsrv.AddFriend(req, api.GetCurrentUser(ctx)); err != nil {
		res.Error(ctx, err)
		return
	}
	msg := ""
	if req.Type == 0 {
		msg = "已发送好友请求"
	} else {
		msg = "已发送入群申请"
	}
	res.Success(ctx, res.Nil, msg)
}

func AddFriendsHandle(ctx *gin.Context) {
	req := &contactsrv.AddFriendHandleRequest{}
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	if err := contactsrv.AddFriendHandle(req, api.GetCurrentUser(ctx)); err != nil {
		res.Error(ctx, err)
		return
	}
	res.Success(ctx, res.Nil)
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

func GetFriendInfo(ctx *gin.Context) {
	friendUuid := ctx.PostForm("friend")
	info := contactsrv.GetFriendInfo(friendUuid, api.GetCurrentUser(ctx))
	res.Success(ctx, info)
}
