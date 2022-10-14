import {GetCurrentUuid} from "../../utis/cache";
import {MessageMainTypes, MessageTypes, VideoCallStatus} from "./const";

const time = () => {
    return Math.ceil((new Date()).getTime() / 1000)
}

const CurrentChatInfo = {
    Recipient: ""
}

const commonStruct = () => {
    return {
        type: MessageMainTypes.chat,
        recipient: CurrentChatInfo.Recipient,
        sender: GetCurrentUuid(),
        send_time: time()
    }
}

export const setRecipient = (recipient) => {
    CurrentChatInfo.Recipient = recipient
}

export const Chat = {
    ping() {
        return {
            content: "ping",
            type: MessageMainTypes.ping,
        }
    },
    text(content) {
        return {
            content: content,
            second_type: MessageTypes.text,
            ...commonStruct(),
        }
    },
    image(address) {
        return {
            content: address,
            second_type: MessageTypes.image,
            ...commonStruct(),
        }
    },
    videoCall(status = VideoCallStatus.wait, memberUuids = []) {
        // status
        // 0 发起视频请求（等待同意）
        // 1 取消视频请求
        // 2 对方同意视频请求
        // 3 对方拒绝视频请求
        // 4 等待超时
        // memberUuids 群聊视频时需要指定视频通话成员（有人要数量限制）
        return {
            content: memberUuids.length ? memberUuids.join(",") : GetCurrentUuid(),
            second_type: MessageTypes.videoCall,
            status: Number(status),
            ...commonStruct(),
        }
    },
}
