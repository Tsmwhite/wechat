<template>
    <div id="app">
        <router-view/>
<!--        <keep-alive>-->
<!--            <router-view/>-->
<!--        </keep-alive>-->
        <!--        <main-tab></main-tab>-->
    </div>
</template>

<script>
import MainTab from "./components/base/main-tab";

export default {
    name: 'App',
    components: {MainTab}
}
</script>

<style>
html,body,#app {
    height: 100%;
}
.h100 {
    height: 100%;
}
body {
    background: #f7f7f7;
}

#app {
    font-family: PingFang SC, 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
}
</style>
