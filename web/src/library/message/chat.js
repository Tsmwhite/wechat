let Recipient = "" // 消息接收者
const time = () => {
    return (new Date()).getTime() / 1000
}
const setRecipient = (recipient) => {
    Recipient = recipient
}
const commonStruct = {
    Recipient: Recipient,
    SendTime: time()
}
export const Chat = {
    ping() {
        return {
            Content: "ping",
            Type: 0,
        }
    },
    text(content) {
        return {
            Content: content,
            Type: 200,
            SecondType: 400,
            ...commonStruct
        }
    },
    image(address) {
        return {
            Content: address,
            Type: 200,
            SecondType: 401,
            ...commonStruct
        }
    },
}

export default {
    setRecipient,
    Chat,
}
