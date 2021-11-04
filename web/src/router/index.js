import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'index',
            component: () => import ('@/views/Index'),
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login.vue'),
        },
        {
            path: '/chat',
            name: 'chat',
            component: () => import('@/components/chat/chat-box'),
        }
    ]
})
