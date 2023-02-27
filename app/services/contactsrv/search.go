package contactsrv

import (
	model "wechat/models"
)

type SearchRequest struct {
	Keyword string
}

func SearchFriends(req *SearchRequest, currentUser *model.User) []map[string]interface{} {
	likeKeyword := "%" + req.Keyword + "%"
	var data []map[string]interface{}
	model.DB().Table("users u").
		Select("u.avatar,u.uuid,u.name,f.id as fid,c.user,c.friend,c.remark,c.type").
		Joins("INNER JOIN `friends` f ON u.uuid = f.friend AND f.user = ? ", currentUser.Uuid).
		Joins("INNER JOIN `contacts` c ON u.uuid = c.friend AND c.user = ? ", currentUser.Uuid).
		Where("(u.mail = ? OR f.name LIKE ? OR f.remark LIKE ?) AND u.is_del = 0", req.Keyword,likeKeyword,likeKeyword).
		Scan(&data)
	return data
}

func SearchUser(req *SearchRequest, currentUser *model.User) []map[string]interface{} {
	var data []map[string]interface{}
	model.DB().Table("users u").
		Select("u.avatar,u.uuid,u.name,f.id as fid,c.user,c.friend,c.remark,c.type").
		Joins("LEFT JOIN `friends` f ON u.uuid = f.friend AND f.user = ? ", currentUser.Uuid).
		Joins("LEFT JOIN `contacts` c ON u.uuid = c.friend AND c.user = ?", currentUser.Uuid).
		Where("u.mail = ? AND u.is_del = 0", req.Keyword).
		Scan(&data)
	return data
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
