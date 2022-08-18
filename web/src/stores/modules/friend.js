import Vue from "vue"

export default {
    state: {
        FriendMap: {
            // 好友uuid：好友信息
            0: {},
        }
    },
    mutations: {
        UPDATE(state, userInfo) {
            Vue.set(state.FriendMap, userInfo.uuid, userInfo)
        },
    },
    actions: {
        update({commit}, userInfo) {
            //console.log("userInfo", userInfo)
            commit("UPDATE", userInfo)
        },
    },
    getters: {},
}

