import request from "./request";

export const getContacts = () => {
    return request.post('/chat/contacts')
}

export const getUserInfo = (data) => {
    return request.get('/getUserInfo', {
        params: data,
    })
}

export const getHistory = (data) => {
    return request.post('/chat/getHistory', data)
}

export default {
    getContacts,
    getUserInfo,
    getHistory,
}
