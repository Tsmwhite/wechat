<template>
    <van-popup v-model="flag"
               style="width: 80%;padding: 16px;border-radius: 8px;">
        <h3 style="margin-top: 0;">添加好友</h3>
        <van-form @submit="onSubmit">
            <van-field
                v-model="formState.content"
                rows="3"
                autosize
                label="验证信息"
                type="textarea"
                maxlength="120"
                placeholder="请输入验证信息（必填）"
                show-word-limit
            />
            <van-field
                v-model="formState.remark"
                label="备注信息"
                placeholder="通过好友申请后自动备注"
                maxlength="12"
                show-word-limit
            />
            <div style="margin: 16px;">
                <van-button round block type="info" native-type="submit">提交</van-button>
            </div>
        </van-form>
    </van-popup>
</template>

<script>
import {sendAddFriendApply} from "../api/common";
import {copyObject} from "../utis/tools";

export default {
    name: "add-friend-modal",
    props: {
        userInfo: {
            type: Object,
            default: () => {
                return {}
            },
        }
    },
    data() {
        return {
            flag: false,
            formState: {
                type: 0,
                uuid: "",
                content: "",
                remark: "",
            },
            formStateDefault: {},
        }
    },
    mounted() {
        this.formStateDefault = copyObject(this.formState)
    },
    methods: {
        show() {
            this.reset()
            this.flag = true
        },
        reset() {
            this.formState = copyObject(this.formStateDefault)
        },
        onSubmit() {
            this.formState.content = this.formState.content.trim()
            this.formState.remark = this.formState.remark.trim()
            if (this.formState.content === "") {
                return this.$toast.fail("请输入验证信息")
            }
            sendAddFriendApply({
                ...this.formState,
                uuid: this.userInfo.uuid,
            }).then(res => {
                this.$toast.success(res.msg)
                this.flag = false
            })
        }
    }
}
</script>

<style scoped>

</style>
