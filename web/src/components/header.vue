<template>
    <div>
    <el-menu
            :default-active="activeIndex2"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff"
            style="padding-top: 12px"
            active-text-color="#ffd04b">
        <el-menu-item index="1">
            <el-badge :value="12"
                      :max="99"
                      class="item">
             消息中心
            </el-badge>
        </el-menu-item>
        <el-submenu index="2">
            <template slot="title">我的工作台</template>
            <el-menu-item index="2-1">选项1</el-menu-item>
            <el-menu-item index="2-2">选项2</el-menu-item>
            <el-menu-item index="2-3">选项3</el-menu-item>
            <el-submenu index="2-4">
                <template slot="title">选项4</template>
                <el-menu-item index="2-4-1">选项1</el-menu-item>
                <el-menu-item index="2-4-2">选项2</el-menu-item>
                <el-menu-item index="2-4-3">选项3</el-menu-item>
            </el-submenu>
        </el-submenu>
        <el-menu-item
                style="float: right"
                index="3">
            <div class="demo-type">
                <div>
                    <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                    admin
                </div>
            </div>
        </el-menu-item>
    </el-menu>
    </div>
</template>

<script>
    export default {
        name:"header",
        data() {
            return {
                activeIndex: '1',
                activeIndex2: '1'
            };
        },
        methods: {
            handleSelect(key, keyPath) {
                window.console.log(key, keyPath);
            }
        }
    }
</script>