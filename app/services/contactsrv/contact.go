package contactsrv

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"wechat/app/services"
	model "wechat/models"
)

type AddContactRequest struct {
	Contact   string
	Type      int
	isPrivate int `validate:"noRequired"` // 0 私聊；1 群聊
}

type GetContactsRequest struct {
	Keyword string `validate:"noRequired & filter"`
	Page    int    `validate:"noRequired"`
	Size    int    `validate:"noRequired"`
}

// GetContactInfoByRoom 获取联系人信息
func GetContactInfoByRoom(roomUuid string, user *model.User) *model.Contact {
	condition := &model.Condition{
		Table:  "contacts",
		Fields: []string{"contacts.*", "rooms.member_num"},
		Joins:  []string{"LEFT JOIN rooms ON contacts.room = rooms.uuid"},
		Where: map[string]interface{}{
			"user": user.Uuid,
			"room": roomUuid,
		},
	}
	result := new(model.Contact)
	model.First(condition, &result)
	return result
}

// GetContacts 获取联系列表
func GetContacts(req *GetContactsRequest, user *model.User) []map[string]interface{} {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 200
	}
	condition := &model.Condition{
		Table:  "contacts",
		Fields: []string{"contacts.*", "rooms.member_num"},
		Joins:  []string{"LEFT JOIN rooms ON contacts.room = rooms.uuid"},
		Where: map[string]interface{}{
			"user": user.Uuid,
		},
		Limit:  (req.Page - 1) * req.Size,
		Offset: req.Size,
		Order: "last_time DESC",
	}
	if req.Keyword != "" {
		condition.SqlStr = fmt.Sprintf("`name` LIKE '%%%s%%' OR remark LIKE '%%%s%%'", req.Keyword, req.Keyword)
	}
	result := services.NewResult()
	model.Find(condition, &result)
	if len(result) > 0 {
		var rooms []string
		for _, item := range result {
			rooms = append(rooms, fmt.Sprint(item["room"]))
		}
		// 未读消息
		var unread []map[string]interface{}
		table := model.GetTableName("receive_messages", user.Uuid)
		model.DB().Table(table).
			Select("room", "COUNT(`msg_uuid`) as unread_count").
			Where("room in ?", rooms).
			Where("recipient = ?", user.Uuid).
			Where("sender <> ?", user.Uuid).
			Where("is_del = 0").
			Where("is_read = 0").
			Group("room").
			Find(&unread)
		unreadMap := make(map[interface{}]interface{})
		for _, item := range unread {
			unreadMap[item["room"]] = item["unread_count"]
		}

		for _, item := range result {
			if val, ok := unreadMap[item["room"]]; ok {
				item["unread_count"] = val
			} else {
				item["unread_count"] = 0
			}
		}
	}
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
	return CreateContactTx(user, cInfo, model.DB())
}

// CreateContactTx 添加为联系人
func CreateContactTx(user, contacts *model.User, tx *gorm.DB) error {
	c := &model.Contact{}
	where := map[string]interface{}{
		"user":   user.Uuid,
		"friend": contacts.Uuid,
	}
	// 查询是否存在联系人
	model.First(&model.Condition{
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
		c.Friend = contacts.Uuid
		c.Room = CreateRoomKey(user, contacts)
		return tx.Create(c).Error
	}
	return nil
}
