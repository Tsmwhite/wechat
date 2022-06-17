import Vue from 'vue'
import Router from 'vue-router'
import {HasLogin} from "../utis/cache";

Vue.use(Router)

const RouterInstance = new Router({
    routes: [
        {
            path: '/',
            name: 'index',
            component: () => import ('@/views/Index'),
            meta: {
                auth: true,
            }
        },
        {
            path: '/index',
            name: 'index',
            component: () => import ('@/views/Index'),
            meta: {
                auth: true,
            }
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login.vue'),
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
