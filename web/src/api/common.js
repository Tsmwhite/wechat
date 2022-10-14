import request from "./request";

export const getContacts = () => {
    return request.post('/chat/contacts')
}

export const getContactInfo = (data) => {
    return request.post('/chat/getContactInfo', data)
}

export const getUserInfo = (data) => {
    return request.get('/getUserInfo', {
        params: data,
    })
}

export const getHistory = (data) => {
    return request.post('/chat/getHistory', data)
}

export const getFriendInfo = (data) => {
    return request.post('/chat/getFriendInfo', data)
}

export const getFriendNotice = (data) => {
    return request.post('/chat/getFriendNotice', data)
}

export const searchFriends = (data) => {
    return request.post('/chat/searchFriend', data)
}

export const searchUser = (data) => {
    return request.post('/chat/searchUser', data)
}

export const searchRoom = (data) => {
    return request.post('/chat/searchRoom', data)
}

// 发送添加好友申请
export const sendAddFriendApply = (data) => {
    return request.post('/chat/addFriends', data)
}

// 处理添加好友申请
export const handleAddFriendApply = (data) => {
    return request.post('/chat/addFriendsHandle', data)
}

export default {
    getContacts,
    getUserInfo,
    getHistory,
    getFriendNotice,
    handleAddFriendApply,
}
