// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router/index'
import Vant, {Lazyload} from 'vant';
import 'vant/lib/index.css';
import WEB_SOCKET from "./library/socket/web-socket";
import {Chat, setRecipient} from "./library/message/chat"
import store from "./stores/index"
import {GetCurrentUuid, HasLogin} from "./utis/cache";

Vue.use(Lazyload);
Vue.use(Vant);
Vue.prototype.$WebSocket = WEB_SOCKET
Vue.prototype.$Chat = Chat
Vue.prototype.$SetR = setRecipient
Vue.prototype.$CurrentUuid = HasLogin() ? GetCurrentUuid() : 0
Vue.prototype.$Log = (...data) => {
    console.log("[system]:", ...data)
}
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    store,
    components: {App},
    template: '<App/>'
})
