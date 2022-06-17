import store from "../../stores/index"

export default (messageObj) => {
    store.dispatch("push", messageObj)
}
