<template>
    <el-container>
        <el-aside style="width: 60%;">
            <img src="../assets/images/login_bg.jpg" style="position: fixed;width: 100%;height: 100%;z-index: -1"/>
        </el-aside>
        <el-container style="padding-top: 120px;">
            <el-form ref="form" label-width="80px" style="background: rgba(0,0,0,0.2);padding-right: 46px">
                <p style="text-align: right;color: #e4b949;">账户登录</p>
                <el-form-item label="账户">
                    <el-input v-model="account" placeholder="请输入账户"></el-input>
                </el-form-item>

                <el-form-item label="密码">
                    <el-input v-model="password" placeholder="请输入密码" show-password></el-input>
                </el-form-item>

                <el-form-item label="记住密码">
                    <el-switch v-model="rem_password"></el-switch>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">登录</el-button>
                    <el-button>没有账户去注册</el-button>
                </el-form-item>
            </el-form>
        </el-container>
    </el-container>
</template>

<script>
    export default {
        name:"login",
        data() {
            return {
                account:"",
                password:"",
                rem_password:"",
            }
        },
        methods: {
            onSubmit() {
                window.console.log('submit!');
            }
        }
    }
</script>

<style scoped>

</style>