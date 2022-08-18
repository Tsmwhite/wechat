package message

const (
	TypeHeartbeat          = iota // 心跳
	TypeCommandRunInApp           // app运行指令
	TypeCommandRunInSystem        // 系统后运行指令
)

const (
	TypeSystemNoticeFull   = iota + 100 // 系统全屏通知
	TypeSystemNoticeTop                 // 系统顶部通知
	TypeSystemNoticeCenter              // 系统中部提示
)

const (
	TypeChatDefault    = iota + 200 // 默认聊天消息
	TypeAddFriendReq				// 添加好友请求
	TypeChatTip                     // 提示性消息
	TypeChatGroup					// 群消息
	TypeGroupNotice                 // 群通知
	TypeGroupAnonymous              // 群匿名消息
)

//  聊天消息子类
const (
	TypeText      = iota + 400 // 文本
	TypeImage                  // 图片
	TypeVoice                  // 语音
	TypeVideo                  // 视频
	TypeImageText              // 图文
	TypeFile                   // 文件
)
