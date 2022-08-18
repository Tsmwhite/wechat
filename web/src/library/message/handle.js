import store from "../../stores/index"

export default (messageObj) => {
    switch (messageObj.type) {
        case 0:
            break
        case 200:
            store.dispatch("push", messageObj)
            break
    }
}
