// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router/index'
import Vant from 'vant';
import 'vant/lib/index.css';
import WEB_SOCKET from "@/library/socket/web-socket";
import {Message} from "./library/message/type";
import {Chat,setRecipient} from "./library/message/chat"
import {Lazyload} from 'vant';
import store from "./stores/index"

Vue.use(Lazyload);
Vue.use(Vant);
Vue.prototype.$WebSocket = WEB_SOCKET
Vue.prototype.$MsgType = Message
Vue.prototype.$Chat = Chat
Vue.prototype.$SetR = setRecipient
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    store,
    components: {App},
    template: '<App/>'
})
