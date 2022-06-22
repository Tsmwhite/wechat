package contact

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services/contactsrv"
	"wechat/app/utils/validations"
)

func verify(ctx *gin.Context) *contactsrv.SearchRequest {
	req := &contactsrv.SearchRequest{}
	if err := api.VerifyParams(ctx, req, map[string]string{validations.Required: "请输入搜索关键字"}); err != nil {
		res.Error(ctx, err)
		return nil
	}
	return req
}

func SearchFriends(ctx *gin.Context) {
	req := verify(ctx)
	if req != nil {
		
	}
}

func SearchUser(ctx *gin.Context) {

}

func SearchRoom(ctx *gin.Context) {

}
