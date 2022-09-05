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

func SearchUser(req *SearchRequest, currentUser *model.User) []*model.ShowAppUser {
	var users []*model.User
	var resList []*model.ShowAppUser
	model.Find(&model.Condition{
		Table: "users",
		Where: map[string]interface{}{
			"mail": req.Keyword,
		},
	}, &users)
	if len(users) > 0 {
		for _, user := range users {
			resList = append(resList, user.ShowAppUser())
		}
	}
	return resList
}

func SearchRoom(req *SearchRequest, currentUser *model.User) []*model.Room {
	var rooms []*model.Room
	model.First(&model.Condition{
		Table: "rooms",
		Where: map[string]interface{}{
			"uuid": req.Keyword,
		},
	}, &rooms)
	return rooms
}
