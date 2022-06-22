package contactsrv

import (
	model "wechat/models"
)

type SearchRequest struct {
	Keyword string
}

func SearchFriends(req *SearchRequest, user *model.User) []*model.Friend {
	var friends []*model.Friend
	//sql := "SELECT `uuid`,`friend`,`remark`,`name`,`avatar`,`type` FROM `friends` WHERE `is_del` = 0 AND `user` = ? AND (`name` = ? OR remark = ?) LIMIT 100"
	model.DB().
		Table("friends").
		Select("uuid").Where("is_del = 0").
		Where("user = ?", user.Uuid).
		Where("name = ? OR remark = ?", req.Keyword, req.Keyword).
		Limit(100).Scan(friends)
	return friends
}

func SearchUser(req *SearchRequest, user *model.User) {
	var friends []*model.Friend
	sql := "SELECT `uuid`,`friend`,`remark`,`name`,`avatar`,`type` FROM `friends` WHERE `is_del` = 0 AND `user` = ? AND (`name` = ? OR remark = ?) LIMIT 100"
	model.DB().Raw(sql, user.Uuid, req.Keyword, req.Keyword).Scan(friends)
}

func SearchRoom(req *SearchRequest, user *model.User) {

}
