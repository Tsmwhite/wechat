const PREFIX = "Tsmwhite:"
export const SetSessionStorage = (key, value) => {
    window.sessionStorage.setItem(PREFIX + key, value);
};

export const GetSessionStorage = (key) => {
    return window.sessionStorage.getItem(PREFIX + key);
};

export const SetLocalStorage = (key, value) => {
    window.localStorage.setItem(PREFIX + key, value);
};

export const GetLocalStorage = (key) => {
    return window.localStorage.getItem(PREFIX + key);
};

export const RemoveLocalStorage = (key) => {
    return window.localStorage.removeItem(PREFIX + key);
};

export const CurrentContactCachetKey = "current-contact"
const localStorageKeys = [
    "token",
    "user-info",
    CurrentContactCachetKey,
]

const Cache = {}

export const GetToken = () => {
    if (!Cache.token) {
        Cache.token = GetLocalStorage("token")
    }
    return Cache.token
}

export const GetCurrentUuid = () => {
    if (!Cache.uuid) {
        let userInfo = GetLocalStorage("user-info")
        userInfo = JSON.parse(userInfo)
        Cache.uuid = userInfo.uuid
    }
    return Cache.uuid
}

export const GetUserInfo = () => {
    if (!Cache.userInfo) {
        let userInfo = GetLocalStorage("user-info")
        Cache.userInfo = JSON.parse(userInfo)
    }
    return Cache.userInfo || {}
}

export const LogoutAfterCache = () => {
    for (let key of localStorageKeys) {
        RemoveLocalStorage(key)
    }
}

export const LoginAfterCache = (userInfo) => {
    SetLocalStorage("token", userInfo.token)
    SetLocalStorage("user-info", JSON.stringify(userInfo))
}

export const HasLogin = () => {
    let token = GetLocalStorage("token")
    return (token !== null && token !== "")
}
