package messagesrv

import (
	"errors"
	"fmt"
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

type ReadMarkRequest struct {
	RoomUuid   string
	FriendUuid string
}

// History 查询历史消息
func History(req *HistoryRequest, user *model.User) ([]map[string]interface{}, error) {
	result := services.NewResult()
	room := model.GetRoomBuyUuid(req.RoomUuid)
	if room.Id == 0 {
		return result, errors.New("房间不存在或状态异常")
	}

	rmsg := &model.ReceiveMessage{
		Recipient: user.Uuid,
	}
	if req.Size == 0 {
		req.Size = 50
	}

	var rMsgs []*model.ReceiveMessage
	// 仅支持查询近3个月消息
	//timeLimit := services.TimeLimit()
	query := rmsg.SetTable().
		Select("id", "msg_uuid").
		Where("room = ?", req.RoomUuid).
		Where("recipient = ?", user.Uuid).
		Where("is_del = ?", 0)
	//Where("send_time > ?", timeLimit)
	if req.LastId > 0 {
		query.Where("id < ?", req.LastId)
	}
	keyword := strings.Trim(req.Keyword, "")
	if keyword != "" {
		query.Where("content LIKE ?", "%"+req.Keyword+"%")
	}
	query.Limit(req.Size).Order("id DESC").Find(&rMsgs)
	if len(rMsgs) < 1 {
		return result, nil
	}
	var msgIds []string
	msgMap := make(map[string]int64)
	for _, msg := range rMsgs {
		msgIds = append(msgIds, msg.MsgUuid)
		msgMap[msg.MsgUuid] = msg.Id
	}
	msg := &model.Message{
		Recipient: req.RoomUuid,
	}
	msg.SetTable().
		Where("uuid IN ?", msgIds).
		Order("id DESC").
		Find(&result)
	for i, record := range result {
		result[i]["id"] = msgMap[fmt.Sprint(record["uuid"])]
	}
	setSenderInfo(result)
	return result, nil
}

// FriendNotice 查询添加好友通知
func FriendNotice(pagination *services.Pagination, user *model.User) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	if pagination.Size == 0 {
		pagination.Size = 50
	}
	// 仅支持查询近3个月消息
	timeLimit := services.TimeLimit()
	table := model.HandleMessageTable
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
	setSenderInfo(result)
	return result, nil
}

func setSenderInfo(data []map[string]interface{}) {
	// 查询发送者信息
	var senders []string
	for _, item := range data {
		sender, ok := item["sender"]
		if ok {
			if val, ok := sender.(string); ok {
				senders = append(senders, val)
			}
		}
	}
	var userList []*model.User
	model.Find(&model.Condition{
		Table:  "users",
		Fields: []string{"avatar", "uuid", "name", "mail"},
		Where: map[string]interface{}{
			"uuid": senders,
		},
	}, &userList)
	userMap := make(map[string]*model.User)
	for _, u := range userList {
		userMap[u.Uuid] = u
	}
	for i, item := range data {
		if sender, ok := item["sender"].(string); ok {
			if val, ok := userMap[sender]; ok {
				data[i]["avatar"] = val.Avatar
				data[i]["name"] = val.Name
				data[i]["mail"] = val.Mail
			}
		}
	}
}

// ReadMark 标记已读消息
func ReadMark(req *ReadMarkRequest, user *model.User) error {
	table := model.GetTableName("receive_messages", user.Uuid)
	room := model.GetRoomBuyUuid(req.RoomUuid)
	if room.Id < 1 {
		return nil
	}
	sqlTx := model.DB().
		Table(table).
		Where("room = ?", req.RoomUuid).
		Where("recipient = ?", user.Uuid)
	if room.Type == model.IsPrivateRoom {
		sqlTx.Where("sender = ?", req.FriendUuid)
	}
	return sqlTx.Where("is_read = ?", model.MessageUnreadStatus).
		Update("is_read", model.MessageReadStatus).Error

}
