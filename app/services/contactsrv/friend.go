package contactsrv

import (
	"errors"
	"fmt"
	"time"
	"wechat/app/api"
	"wechat/core/encrypt"
	"wechat/core/message"
	"wechat/core/redis"
	"wechat/env"
	model "wechat/models"
)

type FriendsRequest struct {
	Keyword string `validate:"noRequired & filter"`
}

type AddFriendRequest struct {
	Uuid   string // 添加用户或群组uuid
	Type   int    // 0 用户 1群组
	Remark string `validate:"noRequired & filter"`
}

type AddFriendHandleRequest struct {
	Uuid   string // 处理消息id
	Status int    // 0 拒绝 1 同意
	Remark string `validate:"noRequired & filter"`
}

func GetFriends(req *FriendsRequest, user *model.User) []map[string]interface{} {
	condition := model.NewCondition()
	condition.Table = "friends"
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

func AddFriend(req *AddFriendRequest, user *model.User) error {
	// 1、查询用户是否存在
	addU := model.NewUser()
	if !addU.CheckMemberByUuid(req.Uuid) {
		return errors.New("用户不存在")
	}
	// 2、查询用户是否存在好友关系
	friend := new(model.Friend)
	model.Find(&model.Condition{
		Table: "friends",
		Where: map[string]interface{}{
			"user":   user.Uuid,
			"friend": addU.Uuid,
		},
	}, friend)
	if friend.Id > 0 {
		return errors.New("已存在好友关系，请勿重复添加")
	}
	// 3、查询是否频繁发送请求
	if err := sendAddFriendRequestLimit(); err != nil {
		return err
	}
	// 4、发送好友申请
	sendAddFriendRequest(user, addU)
	return nil
}

func AddFriendHandle(req *AddFriendHandleRequest, user *model.User) error {
	// 1、查询消息信息
	msg := &model.Message{}
	model.Find(&model.Condition{
		Table: "messages",
		Where: map[string]interface{}{
			"uuid":      req.Uuid,
			"recipient": user.Uuid,
		},
	}, msg)
	if msg.Id == 0 {
		return errors.New("未找到消息")
	}
	if msg.Status != message.AddFriendStatusNormal {
		return errors.New("消息已过期")
	}
	switch req.Status {
	case message.AddFriendStatusReject:

	case message.AddFriendStatusAgree:
		if err := agreeFriendRequest(req, msg, user); err != nil {
			return err
		}
	}
	return nil
}

func agreeFriendRequest(req *AddFriendHandleRequest, msg *model.Message, receipt *model.User) error {
	now := time.Now().Unix()
	// 1、查询发送请求用户信息
	sender := model.GetUserByUuid(msg.Sender)
	if sender.Id == 0 {
		return errors.New("该用户信息异常")
	}
	// 2、添加好友关系并加入双方联系人
	senderFriend := model.Friend{
		User:       sender.Uuid,
		Friend:     receipt.Uuid,
		Name:       receipt.Name,
		Avatar:     receipt.Avatar,
		Remark:     msg.Remark,
		CreateTime: now,
	}
	receiptFriend := model.Friend{
		User:       receipt.Uuid,
		Friend:     sender.Uuid,
		Name:       sender.Name,
		Avatar:     sender.Avatar,
		Remark:     req.Remark,
		CreateTime: now,
	}
	friends := []model.Friend{
		senderFriend,
		receiptFriend,
	}
	model.DB.Create(&friends)
	return nil
}

// sendAddFriendRequest 发送添加好友申请
func sendAddFriendRequest(user, friend *model.User) {
	msg := model.Message{
		Uuid:      encrypt.CreateUuid(),
		Sender:    user.Uuid,
		Recipient: friend.Uuid,
		Type:      message.TypeAddFriendReq,
		Status:    message.AddFriendStatusNormal,
		SendTime:  time.Now().Unix(),
	}
	redis.Init().LPush(redis.Ctx, env.AddFriendRequestHandel, msg)
}

// sendAddFriendRequestLimit 发送好友申请限制
func sendAddFriendRequestLimit() error {
	if redis.Init().LLen(redis.Ctx, env.AddFriendRequestHandel).Val() > 100 {
		return errors.New("当前访问人数过多，请稍后重试")
	}
	return nil
}
