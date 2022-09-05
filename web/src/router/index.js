import Vue from 'vue'
import Router from 'vue-router'
import {HasLogin} from "../utis/cache";

Vue.use(Router)

const RouterInstance = new Router({
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login.vue'),
        },
        {
            path: '/',
            name: 'home',
            component: () => import ('@/views/index'),
            meta: {
                auth: true,
                showMainTab: true,
            }
        },
        {
            path: '/index',
            name: 'index',
            component: () => import ('@/views/index'),
            meta: {
                auth: true,
                showMainTab: true,
            }
        },
        {
            path: '/chat',
            name: 'chat',
            component: () => import ('@/views/chat'),
            meta: {
                auth: true,
            }
        },
        {
            path: '/friend-notice',
            name: 'friendNotice',
            component: () => import ('@/views/friend-notice'),
            meta: {
                auth: true,
            }
        },
    ]
})

RouterInstance.beforeEach((to, form, next) => {
    let hasAuth = to.meta.auth || false
    console.log("hasAuth", hasAuth, HasLogin())
    if (hasAuth && !HasLogin()) {
        next("/login");
        return
    }
    next();
})
export default RouterInstance
