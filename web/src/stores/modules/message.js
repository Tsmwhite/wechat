import Vue from "vue"

export default {
    state: {
        VideoCall: null,
        VideoCallRtcCandidate: null,
        MessageMapList: {
            // 消息接受者（房间id）：消息列表
            0: [],
        }
    },
    mutations: {
        PUSH_MSG(state, message) {
            //console.log("PUSH_MSG", message)
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
        ASSIGN_HISORY(state, {recipient, messages}) {
            if (!state.MessageMapList[recipient]) {
                Vue.set(state.MessageMapList, recipient, messages)
            } else {
                let list = [...messages, ...state.MessageMapList[recipient]]
                Vue.set(state.MessageMapList, recipient, list)
            }
        },
        VIDEO_CALL_SET(state, message) {
            state.VideoCall = message
        },
        VIDEO_CALL_CLEAR(state) {
            state.VideoCall = null
        },
        VIDEO_CALL_RTC_CANDIDATE_SET(state, message) {
            state.VideoCallRtcCandidate = message
        },
        VIDEO_CALL_RTC_CANDIDATE_CLEAR(state) {
            state.VideoCallRtcCandidate = null
        },
    },
    actions: {
        push({commit}, message) {
            commit("PUSH_MSG", message)
        },
        history({commit}, {recipient, messages}) {
            commit("ASSIGN_HISORY", {recipient, messages})
        },
        videoCallSet({commit}, message) {
            commit("VIDEO_CALL_SET", message)
        },
        videoCallClear({commit}) {
            commit("VIDEO_CALL_CLEAR")
        },
        videoCallRtcCandidateSet({commit}, message) {
            commit("VIDEO_CALL_RTC_CANDIDATE_SET", message)
        },
        videoCallRtcCandidateClear({commit}) {
            commit("VIDEO_CALL_RTC_CANDIDATE_CLEAR")
        },
    },
    getters: {},
}

