const time = () => {
    return Math.ceil((new Date()).getTime() / 1000)
}

const commonStruct = {
    Recipient: "",
    SendTime: time()
}

export const setRecipient = (recipient) => {
    commonStruct.Recipient = recipient
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
