package contactsrv

import (
	model "wechat/models"
)

type SearchRequest struct {
	Keyword string
}

func SearchFriends(req *SearchRequest, currentUser *model.User) []*model.Friend {
	var friends []*model.Friend
	likeKeyword := "%" + req.Keyword + "%"
	model.DB().
		Table("friends").
		Select([]string{"user", "friend", "remark", "remark", "name", "avatar", "type"}).
		Where("is_del = 0").
		Where("user = ?", currentUser.Uuid).
		Where("name LIKE ? OR remark LIKE ?", likeKeyword, likeKeyword).
		Limit(1000).Find(&friends)
	return friends
}

func SearchUser(req *SearchRequest, currentUser *model.User) *model.ShowAppUser {
	var user *model.User
	model.Find(&model.Condition{
		Table: "users",
		Where: map[string]interface{}{
			"mail": req.Keyword,
		},
	}, user)
	if user == nil {
		return nil
	}
	return user.ShowAppUser()
}

func SearchRoom(req *SearchRequest, currentUser *model.User) *model.Room {
	var room *model.Room
	model.Find(&model.Condition{
		Table: "rooms",
		Where: map[string]interface{}{
			"uuid": req.Keyword,
		},
	}, room)
	return room
}
