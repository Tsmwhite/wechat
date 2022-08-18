package contact

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services/contactsrv"
)

func AddContact(ctx *gin.Context) {

}

// ContactsList 获取联系人列表
func ContactsList(ctx *gin.Context) {
	req := &contactsrv.GetContactsRequest{}
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
		return
	}
	list := contactsrv.GetContacts(req, api.GetCurrentUser(ctx))
	res.Success(ctx, list)
}
