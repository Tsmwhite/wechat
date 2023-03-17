package message

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services"
	"wechat/app/services/messagesrv"
)

func GetHistory(ctx *gin.Context) {
	req := new(messagesrv.HistoryRequest)
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	result, err := messagesrv.History(req, api.GetCurrentUser(ctx))
	if err != nil {
		res.Error(ctx, err)
	} else {
		res.Success(ctx, result)
	}
}

func GetFriendNotice(ctx *gin.Context) {
	pagination := new(services.Pagination)
	if err := api.VerifyParams(ctx, pagination, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	result, err := messagesrv.FriendNotice(pagination, api.GetCurrentUser(ctx))
	if err != nil {
		res.Error(ctx, err)
	} else {
		res.Success(ctx, result)
	}
}

func ReadMark(ctx *gin.Context) {
	req := new(messagesrv.ReadMarkRequest)
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	err := messagesrv.ReadMark(req, api.GetCurrentUser(ctx))
	if err != nil {
		res.Error(ctx, err)
	} else {
		res.Success(ctx, res.Nil, "ok")
	}
}
