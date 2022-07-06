package contact

import (
	"github.com/gin-gonic/gin"
	"wechat/app/api"
	"wechat/app/res"
	"wechat/app/services/contactsrv"
)

func CreateGroup(ctx *gin.Context) {
	req := &contactsrv.CreateGroupRequest{}
	if err := api.VerifyParams(ctx, req, nil); err != nil {
		res.Error(ctx, err)
	} else {
		err = contactsrv.CreateGroup(req, api.GetCurrentUser(ctx))
		if err != nil {
			res.Error(ctx, err)
		} else {
			res.Success(ctx, res.Nil, "创建完成")
		}
	}
}

func InviteJoin(ctx *gin.Context) {

}

func ApplyJoin(ctx *gin.Context) {

}

func JoinGroupHandle(ctx *gin.Context) {

}
