package messagesrv

import (
	"errors"
	"strings"
	"time"
	"wechat/app/services"
	model "wechat/models"
)

type HistoryRequest struct {
	RoomUuid string
	Keyword  string `validate:"noRequired"`
	Page     int    `validate:"noRequired"`
	Size     int    `validate:"noRequired"`
	LastId   int    `validate:"noRequired"`
}

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
	timeLimit := time.Now().Unix() + 3*30*24*60*60
	query := msg.SetTable().
		Where("recipient = ?", req.RoomUuid).
		Where("is_del = ?", 0).
		Where("send_time < ?", timeLimit)
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
