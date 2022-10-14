const MessageMainTypesPing = 0
const MessageMainTypesChat = 200
export const MessageMainTypes = {
    ping: MessageMainTypesPing,
    chat: MessageMainTypesChat,
}

export const MessageTypes = {
    text: 400,
    image: 401,
    videoCall: 406
}

// 视频通话
// 0 发起视频请求
// 1 取消视频请求
// 2 对方同意视频请求
// 3 对方拒绝视频请求
// 4 等待超时

// 10 建立通话（webrtc description传输）
// 11 建立通话确认（webrtc description回传）
export const VideoCallStatus = {
    wait: 0,
    cancel: 1,
    agree: 2,
    reject: 3,
    timeouts: 4,
    candidate: 5,
    createConn: 10,
    createConnConfirm: 11,
}
