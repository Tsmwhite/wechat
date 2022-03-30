package contactsrv

import (
	"fmt"
	"wechat/app/api"
	model "wechat/models"
)

type AddContactRequest struct {
	Contact string
	Type    int
}

type GetContactsRequest struct {
	Keyword string `validate:"noRequired & filter"`
}

func GetContacts(req *GetContactsRequest, user *model.User) []map[string]interface{} {
	condition := model.NewCondition()
	condition.Table = "contacts"
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
