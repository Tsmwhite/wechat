package contactsrv

import (
	"gorm.io/gorm"
	"time"
	"wechat/core/encrypt"
	model "wechat/models"
)

// CreatePrivateRoomTx 建立用户单聊房间
func CreatePrivateRoomTx(tx *gorm.DB, u1, u2 *model.User)  error {
	room := &model.Room{}
	room.Uuid = CreateRoomKey(u1, u2)
	room.Type = model.IsPrivate
	// 成员uuid逗哈分割
	room.Members = u1.Uuid + "," + u2.Uuid
	room.CreateTime = time.Now().Unix()
	return tx.Create(room).Error
}

// CreateRoomKey 生成房间Uuid
// 房间成员uuid排序后md5 （单聊）
func CreateRoomKey(users ...*model.User) string {
	var uuids []string
	for _, user := range users {
		uuids = append(uuids, user.Uuid)
	}
	return encrypt.CreteUuidsKey(uuids)
}

func GetRoom(u1, u2 *model.User) *model.Room {
	roomUuid := CreateRoomKey(u1, u2)
	// 2、查询房间是否存在
	room := &model.Room{}
	model.Find(&model.Condition{
		Table: "rooms",
		Where: map[string]interface{}{
			"uuid": roomUuid,
		},
	}, room)
	return room
}
