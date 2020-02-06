<template>
    <el-container>
        <el-aside style="width: 60%;">
            <img src="../assets/images/login_bg.jpg" style="position: fixed;width: 100%;height: 100%;z-index: -1"/>
        </el-aside>
        <el-container style="padding-top: 120px;">
            <el-form ref="form" label-width="80px" style="background: rgba(0,0,0,0.2);padding-right: 46px">
                <p style="text-align: right;color: #e4b949;">账户注册</p>
                <el-form-item label="账户">
                    <el-input v-model="account" placeholder="请输入邮箱"></el-input>
                </el-form-item>

                <el-form-item label="验证码" style="position:relative;">
                    <el-input v-model="code" placeholder="请输入验证码"></el-input>
                    <span style="position: absolute;left:120px;color: #888888;cursor: pointer">| 发送验证码</span>
                </el-form-item>

                <el-form-item label="密码">
                    <el-input v-model="password" placeholder="请输入密码" show-password></el-input>
                </el-form-item>

                <el-form-item label="确认密码">
                    <el-input v-model="password_" placeholder="请确认密码" show-password></el-input>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">注册</el-button>
                    <el-button>
                        <router-link to="/login" style="color: #888888;text-decoration:none">已有账户去登录</router-link>
                    </el-button>
                </el-form-item>
            </el-form>
        </el-container>
    </el-container>
</template>

<script>
    export default {
        name:"register",
        data() {
            return {
                account:"",
                password:"",
                password_:"",
                code:"",
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