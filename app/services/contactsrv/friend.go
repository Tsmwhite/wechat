package contactsrv

import (
	"errors"
	"fmt"
	"time"
	"wechat/app/services"
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
	Uuid    string // 添加用户或群组uuid
	Type    int    // 0 用户 1群组
	Content string // 验证信息
	Remark  string `validate:"noRequired & filter"`
}

type AddFriendHandleRequest struct {
	Uuid   string // 处理消息id
	Status int    // 0 拒绝 1 同意
	Remark string `validate:"noRequired & filter"`
}

// GetFriendInfo 获取好友信息
func GetFriendInfo(friendUuid string, user *model.User) *model.Friend {
	condition := &model.Condition{
		Table: "friends",
		Where: map[string]interface{}{
			"user":   user.Uuid,
			"friend": friendUuid,
		},
	}
	result := new(model.Friend)
	model.First(condition, &result)
	return result
}

// GetFriends 获取好友列表
func GetFriends(req *FriendsRequest, user *model.User) []map[string]interface{} {
	condition := &model.Condition{
		Table: "friends",
		Where: map[string]interface{}{
			"user": user.Uuid,
		},
		Limit: 500,
	}
	if req.Keyword != "" {
		condition.SqlStr = fmt.Sprintf("`name` LIKE '%%%s%%' OR remark LIKE '%%%s%%'", req.Keyword, req.Keyword)
	}
	result := services.NewResult()
	model.Find(condition, &result)
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
	model.First(&model.Condition{
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
	if err := sendAddFriendRequestLimit(req, user); err != nil {
		return err
	}
	// 4、构造好友申请
	nowTime := time.Now().Unix()
	msg := &model.Message{
		Uuid:       encrypt.CreateUuid(),
		Sender:     user.Uuid,
		Recipient:  addU.Uuid,
		Type:       message.TypeAddFriendReq,
		Status:     message.AddFriendStatusNormal,
		SendTime:   nowTime,
		LogTime:    nowTime,
		UpdateTime: nowTime,
		Content:    req.Content,
		Remark:     req.Remark,
	}
	if err := model.DB().Table(model.HandleMessageTable).Create(msg).Error; err != nil {
		return err
	}
	if err := redis.LPush(env.AddFriendRequestHandel, msg); err != nil {
		log.PrintlnErr("sendAddFriendRequest LPush Msg Error:", err)
	}
	return nil
}

// AddFriendHandle 处理添加好友请求
func AddFriendHandle(req *AddFriendHandleRequest, user *model.User) error {
	msgTable := model.HandleMessageTable
	// 1、查询消息信息
	msg := &model.Message{}
	model.First(&model.Condition{
		Table: msgTable,
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
	// 2、查询是否已存在好友关系
	friend := new(model.Friend)
	model.First(&model.Condition{
		Table: "friends",
		Where: map[string]interface{}{
			"user":   msg.Recipient,
			"friend": msg.Sender,
		},
	}, friend)
	if friend.Id > 0 {
		return errors.New("已存在好友关系，请勿重复添加")
	}
	switch req.Status {
	case message.AddFriendStatusReject:
		// 存在历史申请变更为失效状态
		err := model.DB().Table(msgTable).
			Where("id < ?", msg.Id).
			Where("recipient = ?", user.Uuid).
			Where("sender = ?", msg.Sender).
			Update("status", message.AddFriendStatusLose).Error
		if err != nil {
			log.PrintlnErr("AddFriendHandle -> AddFriendStatusReject -> Update history Error:", err)
			return errors.New("操作失败，请重试")
		}

		// 变更好友申请消息状态
		err = model.DB().Table(model.HandleMessageTable).
			Where("id = ?", msg.Id).
			Updates(map[string]interface{}{
				"status":      message.AddFriendStatusReject,
				"update_time": time.Now().Unix(),
			}).Error
		if err != nil {
			log.PrintlnErr("AddFriendHandle -> AddFriendStatusReject -> Update current Error:", err)
			return errors.New("操作失败，请重试")
		}

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
	room := GetPrivateRoom(sender, receipt)
	tx := model.NewDB().Begin()
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
				log.PrintlnErr("agreeFriendRequest -> CreatePrivateRoomTx Error:", err)
				return errors.New("建立房间失败")
			}
		}
		// 添加好友关系
		if err = tx.Create(&friends).Error; err != nil {
			tx.Rollback()
			log.PrintlnErr("agreeFriendRequest ->  Error:", err)
			return errors.New("建立好友关系失败")
		}
		// 建立双方联系人
		if err = CreateContactsTx(sender, receipt, tx); err != nil {
			tx.Rollback()
			log.PrintlnErr("agreeFriendRequest -> CreateContactsTx Error:", err)
			return errors.New("建立联系人失败")
		}
		// 变更好友申请消息状态
		err = tx.Table(model.HandleMessageTable).Where("uuid = ?", msg.Uuid).Updates(map[string]interface{}{
			"status":      message.AddFriendStatusAgree,
			"update_time": time.Now().Unix(),
		}).Error
		if err != nil {
			tx.Rollback()
			log.PrintlnErr("agreeFriendRequest ->  Error:", err)
			return errors.New("变更好友申请状态失败")
		}
		return tx.Commit().Error
	}()
}

// sendAddFriendRequestLimit 发送好友申请限制
func sendAddFriendRequestLimit(req *AddFriendRequest, user *model.User) error {
	// 查询最近10分钟是否存在未处理请求
	tenMinutesAgoUnix := time.Now().Unix() - (10 * 60)
	msg := &model.Message{}
	model.First(&model.Condition{
		Table: model.HandleMessageTable,
		Where: map[string]interface{}{
			"recipient": req.Uuid,
			"sender":    user.Uuid,
			"status":    message.AddFriendStatusNormal,
		},
		Expr: model.NewExpr("`send_time` > ?", tenMinutesAgoUnix),
	}, msg)

	if msg.Id > 0 {
		return errors.New("已发送申请，请勿频繁发送")
	}

	if redis.Init().LLen(redis.Ctx, env.AddFriendRequestHandel).Val() > 100 {
		return errors.New("当前访问人数过多，请稍后重试")
	}
	return nil
}
