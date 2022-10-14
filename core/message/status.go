package message

const (
	StateSeeding  = iota + 4000//发送中
	StateFail            // 发送失败
	StateSuccess         // 发送完成
	StateRevoke          // 消息撤回
	StateHaveRead        // 消息已读
	StateDelete          // 消息删除
)

// 添加好友申请状态
const (
	AddFriendStatusNormal = iota
	AddFriendStatusAgree
	AddFriendStatusReject
	AddFriendStatusLose
)

// 视频通话状态
// 0 发起视频请求（等待同意）
// 1 取消视频请求
// 2 对方同意视频请求
// 3 对方拒绝视频请求
// 4 等待超时

// 10 建立通话（传输webrtc description传输）
const (
	VideoCallStatusWait = iota
	VideoCallStatusCancel = iota
	VideoCallStatusAgree
	VideoCallStatusReject
	VideoCallStatusTimeouts
)

const (
	VideoCallStatusCreateConn = iota + 10
)