import request from "./request";

export default {
    getContacts: () => {
        return request.post('/api/chat/contacts')
    }
}
