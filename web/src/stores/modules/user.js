const keys = [
    "token",
    "uuid",
    "name",
    "mobile",
    "mail",
    "avatar",
    "role_id",
]
export default {
    state: {
        token: "",
        uuid: "",
        name: "",
        mobile: "",
        mail: "",
        avatar: "",
        role_id: "",
    },
    mutations: {
        SET(state, userInfo) {
            for (let key of keys) {
                state[key] = userInfo[key]
            }
        },
        CLEAR(state) {
            for (let key of keys) {
                state[key] = ""
            }
        },
    },
    actions: {},
    getter: {},
}
