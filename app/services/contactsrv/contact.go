package contactsrv

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"wechat/app/api"
	model "wechat/models"
)

type AddContactRequest struct {
	Contact   string
	Type      int
	isPrivate int `validate:"noRequired"` // 0 私聊；1 群聊
}

type GetContactsRequest struct {
	Keyword string `validate:"noRequired & filter"`
}

// GetContacts 获取联系列表
func GetContacts(req *GetContactsRequest, user *model.User) []map[string]interface{} {
	condition := &model.Condition{
		Table: "contacts",
		Where: map[string]interface{}{
			"user": user.Uuid,
		},
	}
	if req.Keyword != "" {
		condition.SqlStr = fmt.Sprintf("`name` LIKE '%%%s%%' OR remark LIKE '%%%s%%'", req.Keyword, req.Keyword)
	}
	result := api.NewResult()
	model.FindAll(condition, &result)
	return result
}

// CreateContactsTx 建立双方联系人关系
func CreateContactsTx(u1, u2 *model.User, tx *gorm.DB) error {
	err := CreateContactTx(u1, u2, tx)
	if err != nil {
		return err
	}
	err = CreateContactTx(u2, u1, tx)
	if err != nil {
		return err
	}
	return nil
}

// AddContact 添加到联系人列表
func AddContact(req *AddContactRequest, user *model.User) error {
	// 查询联系人信息
	cInfo := model.GetUserByUuid(req.Contact)
	if cInfo.Id == 0 {
		return errors.New("用户信息异常")
	}
	return CreateContactTx(user, cInfo, model.DB)
}

// CreateContactTx 添加为联系人
func CreateContactTx(user, contacts *model.User, tx *gorm.DB) error {
	c := &model.Contact{}
	where := map[string]interface{}{
		"user":   user.Uuid,
		"object": contacts.Uuid,
	}
	// 查询是否存在联系人
	model.Find(&model.Condition{
		Table:      "contacts",
		IncludeDel: true,
		Where:      where,
	}, c)

	// 查询好友信息
	friend := &model.Friend{}
	tx.Table("friends").Where(where).First(friend)
	c.Avatar = contacts.Avatar
	c.Name = contacts.Name
	c.Remark = friend.Remark
	// 已存在该联系人
	if c.Id > 0 {
		// 删除了则直接变更状态
		if c.IsDel == model.IsDel {
			c.IsDel = model.IsNoDel
			c.UpdateTime = time.Now().Unix()
			return tx.Save(c).Error
		}
	} else {
		c.User = user.Uuid
		c.Object = contacts.Uuid
		c.Room = CreateRoomKey(user, contacts)
		return tx.Create(c).Error
	}
	return nil
}
