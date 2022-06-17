import {GetCurrentUuid} from "../../utis/cache";

const time = () => {
    return Math.ceil((new Date()).getTime() / 1000)
}

const CurrentChatInfo = {
    Recipient: ""
}

const commonStruct = () => {
   return {
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
            type: 0,
        }
    },
    text(content) {
        return {
            content: content,
            type: 200,
            second_type: 400,
            ...commonStruct(),
        }
    },
    image(address) {
        return {
            content: address,
            type: 200,
            second_type: 401,
            ...commonStruct(),
        }
    },
}
