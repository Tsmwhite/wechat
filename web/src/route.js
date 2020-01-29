import Vue from "vue";
import VueRouter from "vue-router"
Vue.use(VueRouter)

import index from "./pages/index"
import login from "./pages/login"
import register from "./pages/register"

const routes = [
    { path: '/', component:index},
    { path: '/login', component:login},
    { path: '/register', component: register },
    { path: '/index', component: index },
]

const router = new VueRouter({
    //mode:"history",
    routes // (缩写) 相当于 routes: routes
})

export default router