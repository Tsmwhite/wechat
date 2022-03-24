// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vant from 'vant';
import 'vant/lib/index.css';
import {WEB_SOCKET} from "@/library/socket/web-socket";
import {Message} from "./library/message/type";
import {Lazyload} from 'vant';

Vue.use(Lazyload);
Vue.use(Vant);
Vue.prototype.$WebSocket = WEB_SOCKET
Vue.prototype.$MsgType = Message
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: {App},
    template: '<App/>'
})
