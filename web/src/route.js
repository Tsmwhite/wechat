import Vue from "vue";
import VueRouter from "vue-router"
Vue.use(VueRouter)

import index from "./pages/index"
import login from "./pages/login"
import register from "./pages/register"
import test from "./pages/test"

const routes = [
    { path: '/', component:test},
    { path: '/login', component:login},
    { path: '/register', component: register },
    { path: '/index', component: index },
]

const router = new VueRouter({
    //model:"history",
    routes // (缩写) 相当于 routes: routes
})

export default router