import Vue from 'vue'
import App from './App.vue'
import router from "./route"

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import {WEB_SOCKET} from "@/library/socket/web-socket";

Vue.config.productionTip = false

Vue.use(ElementUI)
Vue.prototype.$WebSocket = WEB_SOCKET

new Vue({
    render: h => h(App),
    router
}).$mount('#app')

