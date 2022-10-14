import store from "../../stores/index"
import {GetCurrentUuid} from "../../utis/cache";
import {MessageMainTypes, MessageTypes, VideoCallStatus} from "./const";

export default (messageObj) => {
    switch (messageObj.type) {
        case MessageMainTypes.ping:
            break
        case MessageMainTypes.chat:
            chatMessageHandle(messageObj)
            break
    }
}

const chatMessageHandle = (messageObj) => {
    switch (messageObj.second_type) {
        case MessageTypes.text:
        case MessageTypes.image:
            store.dispatch("push", messageObj)
            break
        case MessageTypes.videoCall:
            switch (messageObj.status) {
                case VideoCallStatus.wait:
                case VideoCallStatus.agree:
                case VideoCallStatus.createConn:
                case VideoCallStatus.createConnConfirm:
                    if (messageObj.sender !== GetCurrentUuid()) {
                        store.dispatch("videoCallSet", messageObj)
                    }
                    break
                case VideoCallStatus.candidate:
                    if (messageObj.sender !== GetCurrentUuid()) {
                        store.dispatch("videoCallRtcCandidateSet", messageObj)
                    }
                    break
                case VideoCallStatus.cancel:
                    break
                case VideoCallStatus.reject:
                    break
                case VideoCallStatus.timeouts:
                    break
            }
            break
    }
}

