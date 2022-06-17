import Vue from "vue"
import Vuex from "vuex"
import MsgStore from "./modules/message"
import UserStore from "./modules/user"

Vue.use(Vuex)
const storeInstance = new Vuex.Store({
    modules: {
        msg: MsgStore,
        user: UserStore,
    }
})

export default storeInstance

export const LoginAfterStore = (userInfo) => {
    storeInstance.commit("SET", userInfo)
}

export const LogoutAfterStore = () => {
    storeInstance.commit("CLEAR")
}
