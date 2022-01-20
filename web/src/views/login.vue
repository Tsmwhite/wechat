<template>
    <div class="loginApp">
        <img class="bgImage" :src="Bg"/>
        <div class="form-box">
            <h3>尬聊啊！兄嘚！</h3>
            <van-form @submit="onSubmit">
                <van-field v-model="form.phone"
                           type="number"
                           label="手机号"
                           :clearable="true"
                           :clickable="true"
                           placeholder="请输入手机号"
                           :maxlength="11"/>
                <van-field v-model="form.verifyCode"
                           type="number"
                           :clearable="true"
                           :clickable="true"
                           :maxlength="6"
                           center
                           clearable
                           label="验证码"
                           placeholder="请输入验证码">
                    <template #button>
                        <van-button size="small" type="info">发送验证码</van-button>
                    </template>
                </van-field>
                <van-button round type="info" style="width: 80%;margin-top: 20px">登录</van-button>
                <van-checkbox v-model="checked"
                              icon-size="14px"
                              shape="square"
                              style="margin-top: 20px;font-size: 12px;display: flex;justify-content: center;">已阅读并同意使用协议
                </van-checkbox>
            </van-form>
        </div>
    </div>
</template>

<script>
import Bg from "@/assets/login-bg.png"

export default {
    name: "login",
    data() {
        return {
            Bg,
            checked: true,
            form: {
                phone: "",
                verifyCode: "",
            }
        }
    },
    methods: {
        onSubmit(values) {
            if (this.form.phone === "") {
                return this.$toast.fail("请输入手机号")
            }
            if (this.form.phone.length !== 11) {
                return this.$toast.fail("手机号码格式不正确")
            }
            if (this.form.verifyCode === "") {
                return this.$toast.fail("请输入验证码")
            }
        },
    }
}
</script>

<style scoped lang="less">
.loginApp {
    .bgImage {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }

    .form-box {
        text-align: center;
        opacity: 0.9;
        width: 90%;
        height: 280px;
        border-radius: 16px;
        padding: 5px;
        position: absolute;
        background: #FFFFFF;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        margin: auto;
    }
}
</style>
