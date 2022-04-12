package message

const (
	StateSeeding  = iota //发送中
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
)