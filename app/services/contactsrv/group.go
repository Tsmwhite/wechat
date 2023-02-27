package contactsrv

import (
	"errors"
	"strings"
	"time"
	"wechat/core/encrypt"
	"wechat/core/log"
	model "wechat/models"
)

// CreateGroupRequest 创建群聊请求
type CreateGroupRequest struct {
	Title,
	Friends string
}

// 创建群聊参数验证
func (req *CreateGroupRequest) verify(creator *model.User) error {
	friends := strings.Split(req.Friends, ",")
	if len(friends) < 1 {
		return errors.New("缺少群聊成员信息")
	}
	if len(friends) > 120 {
		return errors.New("群聊成员最多120人")
	}
	var existFriends []model.Friend
	// 查询邀请人是否存在好友关系
	model.DB().Table("friends").
		Select("friend").
		Where("user = ?", creator.Uuid).
		Where("is_del = 0").
		Where("friend in ?", friends).
		Find(&existFriends)

	var isFriendsUuid []string
	for _, f := range existFriends {
		isFriendsUuid = append(isFriendsUuid, f.Friend)
	}
	if len(isFriendsUuid) < 1 {
		return errors.New("非好友关系无法邀请入群")
	}
	req.Friends = strings.Join(isFriendsUuid, ",")
	return nil
}

// CreateGroup 创建群聊
func CreateGroup(req *CreateGroupRequest, currentUser *model.User) error {
	if err := req.verify(currentUser); err != nil {
		return err
	}
	users := strings.Split(req.Friends, ",")
	users = append(users, currentUser.Uuid)
	// 创建房间
	roomData := &model.Room{
		Uuid:       encrypt.CreateUuid(),
		Title:      req.Title,
		Members:    req.Friends,
		Creator:    currentUser.Uuid + "," + currentUser.Uuid,
		Type:       model.IsGroup,
		CreateTime: time.Now().Unix(),
		MemberNum:  len(users),
	}
	err := model.DB().Create(roomData).Error
	if err != nil {
		log.Error.Println("CreateGroup CreateRoom Error:", err)
		return errors.New("创建群聊失败")
	}

	now := time.Now().Unix()
	var contactsRelation []model.Contact
	for _, user := range users {
		contactsRelation = append(contactsRelation, model.Contact{
			Room:       roomData.Uuid,
			User:       user,
			Friend:     roomData.Uuid,
			Name:       roomData.Title,
			Type:       model.IsGroup,
			LastTime:   now,
			UpdateTime: now,
		})
	}
	err = model.DB().Create(&contactsRelation).Error
	if err != nil {
		log.Error.Println("CreateGroup ContactsRelation Error:", err)
		return errors.New("邀请好友入群失败")
	}
	return nil
}
