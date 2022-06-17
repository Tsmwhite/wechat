import request from "./request";

export default {
    getContacts: () => {
        return request.post('/chat/contacts')
    }
}
