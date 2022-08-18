import Cookies from 'js-cookie'
import {getUserInfo} from "../api/common";
import store from "../stores/index"

const UserInfoKey = "Tsmwhite-user-info-"
const UserInfoLoadingKey = "Tsmwhite-user-info-loading:"

export const expiredTime = (minute) => {
    return new Date(new Date().getTime() + minute * 60 * 1000);
}

async function getUserInfoReq(uuid) {
    let res = await getUserInfo(uuid)
    return res.data || {
        avatar: "",
        name: "",
    }
}

export const GetUserByUuid = (uuid) => {
    let user = Cookies.get(UserInfoKey + uuid)
    let loading = Cookies.get(UserInfoLoadingKey + uuid)
    if (!user) {
        if (loading) {
            return {
                avatar: "",
                name: "",
            }
        }
        Cookies.set(UserInfoLoadingKey + uuid, 1, {expires: expiredTime(1)})
        getUserInfoReq(uuid).then(user => {
            Cookies.set(UserInfoKey + uuid, JSON.stringify(user), {expires: expiredTime(5)})
            store.dispatch("update", user)
        })
    } else {
        user = JSON.parse(user)
        store.dispatch("update", user)
    }
    return user
}
