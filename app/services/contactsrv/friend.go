package contactsrv

import (
	"fmt"
	"wechat/app/api"
	model "wechat/models"
)

type FriendsRequest struct {
	Keyword string `validate:"noRequired & filter"`
}

type FriendsResponse struct{}

func GetFriends(req *FriendsRequest, user *model.User) []map[string]interface{} {
	condition := &model.Condition{
		Table: "friends",
	}
	condition.Where = map[string]interface{}{
		"user": user.Uuid,
	}
	if req.Keyword != "" {
		condition.SqlStr = fmt.Sprintf("`name` LIKE '%%%s%%' OR remark LIKE '%%%s%%'", req.Keyword, req.Keyword)
	}
	result := api.NewResult()
	model.FindAll(condition, &result)
	return result
}
