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
	}
	result, err := messagesrv.History(req)
	if err != nil {
		res.Error(ctx, err)
	}
	res.Success(ctx, result)
}

func GetFriendNotice(ctx *gin.Context) {
	pagination := new(services.Pagination)
	if err := api.VerifyParams(ctx, pagination, nil); err != nil {
		res.Error(ctx, err)
	}
	result, err := messagesrv.FriendNotice(pagination, api.GetCurrentUser(ctx))
	if err != nil {
		res.Error(ctx, err)
	}
	res.Success(ctx, result)
}
