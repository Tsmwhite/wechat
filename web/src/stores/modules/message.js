import Vue from "vue"

export default {
    state: {
        MessageMapList: {
            // 消息接受者（房间id）：消息列表
            0: [],
        }
    },
    mutations: {
        PUSH_MSG(state, message) {
            console.log("PUSH_MSG", message)
            if (!state.MessageMapList[message.recipient]) {
                Vue.set(state.MessageMapList, message.recipient, [message])
                // state.MessageMapList[message.recipient] = []
            } else {
                let list = [...state.MessageMapList[message.recipient], message]
                Vue.set(state.MessageMapList, message.recipient, list)
                //state.MessageMapList[message.recipient].push(message)
            }
        },
        READ_MSG(state, {recipient, messageIds}) {

        },
    },
    actions: {
        push({commit}, message) {
            commit("PUSH_MSG", message)
        }
    },
    getters: {},
}

