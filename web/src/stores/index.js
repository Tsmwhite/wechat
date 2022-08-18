import Vue from "vue"
import Vuex from "vuex"
import MsgStore from "./modules/message"
import UserStore from "./modules/user"
import FriendStore from "./modules/friend"

Vue.use(Vuex)
const storeInstance = new Vuex.Store({
    modules: {
        msg: MsgStore,
        user: UserStore,
        friend: FriendStore,
    }
})

export default storeInstance

export const LoginAfterStore = (userInfo) => {
    storeInstance.commit("SET", userInfo)
}

export const LogoutAfterStore = () => {
    storeInstance.commit("CLEAR")
}
