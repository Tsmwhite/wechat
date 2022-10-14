package messagesrv

import (
	"errors"
	"strings"
	"wechat/app/services"
	"wechat/core/message"
	model "wechat/models"
)

type HistoryRequest struct {
	RoomUuid string
	Keyword  string `validate:"noRequired"`
	Page     int    `validate:"noRequired"`
	Size     int    `validate:"noRequired"`
	LastId   int    `validate:"noRequired"`
}

// History 查询历史消息
func History(req *HistoryRequest) ([]map[string]interface{}, error) {
	result := services.NewResult()
	room := model.GetRoomBuyUuid(req.RoomUuid)
	if room.Id == 0 {
		return result, errors.New("房间不存在或状态异常")
	}
	msg := &model.Message{
		Recipient: req.RoomUuid,
	}
	if req.Size == 0 {
		req.Size = 50
	}

	// 仅支持查询近3个月消息
	timeLimit := services.TimeLimit()
	query := msg.SetTable().
		Where("recipient = ?", req.RoomUuid).
		Where("is_del = ?", 0).
		Where("send_time > ?", timeLimit)
	if req.LastId > 0 {
		query.Where("id < ?", req.LastId)
	}
	keyword := strings.Trim(req.Keyword, "")
	if keyword != "" {
		query.Where("content LIKE ?", "%"+req.Keyword+"%")
	}
	query.Limit(req.Size).Order("id DESC").Find(&result)
	return result, nil
}

// FriendNotice 查询添加好友通知
func FriendNotice(pagination *services.Pagination, user *model.User) ([]map[string]interface{}, error) {
	result := services.NewResult()
	if pagination.Size == 0 {
		pagination.Size = 50
	}
	// 仅支持查询近3个月消息
	timeLimit := services.TimeLimit()
	table := model.GetTableName("messages", user.Uuid)
	var sql string
	var args []interface{}
	if pagination.LastId > 0 {
		sql = `SELECT m.* FROM (
				SELECT * FROM ` + table + `
					WHERE id < ? AND 
					type = ? AND 
					recipient = ? AND 
					is_del = 0 AND 
					send_time > ?
					GROUP BY sender
					ORDER BY id DESC
               ) as m GROUP BY m.sender LIMIT ?`
		args = append(args, pagination.LastId)
	} else {
		sql = `SELECT m.* FROM (
				SELECT * FROM ` + table + `
					WHERE type = ? AND 
					recipient = ? AND 
					is_del = 0 AND 
					send_time > ?
					GROUP BY sender
					ORDER BY id DESC
               ) as m GROUP BY m.sender LIMIT ?`
	}
	args = append(args, message.TypeAddFriendReq, user.Uuid, timeLimit, pagination.Size)
	model.DB().Raw(sql, args...).Scan(&result)
	return result, nil
}
