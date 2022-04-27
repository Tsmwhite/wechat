package contactsrv

import (
	"errors"
	"fmt"
	"time"
	"wechat/app/api"
	"wechat/core/encrypt"
	"wechat/core/log"
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

// GetFriends 获取好友列表
func GetFriends(req *FriendsRequest, user *model.User) []map[string]interface{} {
	condition := &model.Condition{
		Table: "friends",
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

// AddFriend 添加好友请求
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
	// 4、构造好友申请
	msg := &model.Message{
		Uuid:      encrypt.CreateUuid(),
		Sender:    user.Uuid,
		Recipient: addU.Uuid,
		Type:      message.TypeAddFriendReq,
		Status:    message.AddFriendStatusNormal,
		SendTime:  time.Now().Unix(),
		Remark:    req.Remark,
	}
	return sendAddFriendRequest(msg)
}

// AddFriendHandle 处理添加好友请求
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
		// 变更好友申请消息状态
		msg.Status = message.AddFriendStatusReject
		msg.UpdateTime = time.Now().Unix()
		model.DB.Save(msg)
	case message.AddFriendStatusAgree:
		if err := agreeFriendRequest(req, msg, user); err != nil {
			return err
		}
	}
	return nil
}

// agreeFriendRequest 同意好友请求
func agreeFriendRequest(req *AddFriendHandleRequest, msg *model.Message, receipt *model.User) error {
	// 查询发送请求用户信息
	sender := model.GetUserByUuid(msg.Sender)
	if sender.Id == 0 {
		return errors.New("该用户信息异常")
	}
	lockKey := CreateRoomKey(sender, receipt)
	lockKey = "AddFriendAgreeHandle:" + lockKey
	if err := redis.Init().SetNX(redis.Ctx, lockKey, 1, time.Minute*2).Err(); err != nil {
		return err
	}
	defer func() {
		redis.Init().Del(redis.Ctx, lockKey)
	}()
	now := time.Now().Unix()
	// 好友关系数据
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
	// 查询二人是否存在房间
	room := GetRoom(sender, receipt)
	tx := model.DB.Begin()
	return func() error {
		defer func() {
			if e := recover(); e != nil {
				tx.Rollback()
			}
		}()
		// 建立房间
		var err error
		if room.Id == 0 {
			if err = CreatePrivateRoomTx(tx, sender, receipt); err != nil {
				tx.Rollback()
				log.Error.Println("agreeFriendRequest -> CreatePrivateRoomTx Error:", err)
				return errors.New("建立房间失败")
			}
		}
		// 添加好友关系
		if err = tx.Create(&friends).Error; err != nil {
			tx.Rollback()
			log.Error.Println("agreeFriendRequest ->  Error:", err)
			return errors.New("建立好友关系失败")
		}
		// 建立双方联系人
		if err = CreateContactsTx(sender, receipt, tx); err != nil {
			tx.Rollback()
			log.Error.Println("agreeFriendRequest -> CreateContactsTx Error:", err)
			return errors.New("建立联系人失败")
		}
		// 变更好友申请消息状态
		msg.Status = message.AddFriendStatusAgree
		msg.UpdateTime = now
		if err = tx.Save(msg).Error; err != nil {
			tx.Rollback()
			log.Error.Println("agreeFriendRequest ->  Error:", err)
			return errors.New("变更好友申请状态失败")
		}
		return tx.Commit().Error
	}()
}

// sendAddFriendRequest 发送添加好友申请
func sendAddFriendRequest(msg *model.Message) error {
	return model.DB.Create(&msg).Error
	//err := redis.Init().LPush(redis.Ctx, env.AddFriendRequestHandel, msg).Err()
	//fmt.Println("sendAddFriendRequest", err)
}

// sendAddFriendRequestLimit 发送好友申请限制
func sendAddFriendRequestLimit() error {
	if redis.Init().LLen(redis.Ctx, env.AddFriendRequestHandel).Val() > 100 {
		return errors.New("当前访问人数过多，请稍后重试")
	}
	return nil
}
