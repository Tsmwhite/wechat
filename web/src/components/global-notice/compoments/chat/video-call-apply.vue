<template>
    <div class="apply-notice-container" v-if="applyNoticeFlag">
        <div class="apply-notice-box">
            <img class="avatar" :src="friendInfo.avatar"/>
            <div class="title">
                <span style="color: #1989fa">{{ friendInfo.remark || friendInfo.name }}</span>
                邀请你视频通话
            </div>
            <div style="margin-top: 12px">
                <van-button type="warning"
                            size="small"
                            style="width: 120px" @click="reject">拒绝
                </van-button>
                <van-button type="info"
                            size="small"
                            style="width: 120px;margin-left: 24px;" @click="agree">同意
                </van-button>
            </div>
        </div>
    </div>
</template>

<script>
import {getFriendInfo} from "@/api/common";
import {VideoCallStatus} from "../../../../library/message/const";
import {getContactInfo} from "../../../../api/common";
import {CurrentContactCachetKey, SetLocalStorage} from "../../../../utis/cache";

export default {
    name: "video-call-apply",
    data() {
        return {
            applyNoticeFlag: false,
            friendInfo: {},
        }
    },
    computed: {
        videoCallMsg() {
            let storeData = this.$store.state
            return storeData.msg.VideoCall
        },
    },
    watch: {
        videoCallMsg() {
            if (this.videoCallMsg) {
                switch (Number(this.videoCallMsg.status)) {
                    case VideoCallStatus.wait:
                        this.show()
                        break
                }
            }
        },
    },
    methods: {
        show() {
            getFriendInfo({
                friend: this.videoCallMsg.sender
            }).then(res => {
                this.friendInfo = res.data
                this.applyNoticeFlag = true
            })
        },
        close() {
            this.applyNoticeFlag = false
        },
        agree() {
            this.$Log("agree", this.videoCallMsg)
            let message = this.$Chat.videoCall(VideoCallStatus.agree)
            message.parent = this.videoCallMsg.uuid
            message.recipient = this.videoCallMsg.recipient
            this.$WebSocket.send(message)
            this.createConnect()
        },
        reject() {
            this.$Log("reject", this.videoCallMsg)
            let message = this.$Chat.videoCall(VideoCallStatus.reject)
            message.parent = this.videoCallMsg.uuid
            message.recipient = this.videoCallMsg.recipient
            this.$WebSocket.send(message)
            this.close()
        },
        createConnect() {
            getContactInfo({
                room: this.videoCallMsg.recipient
            }).then(res => {
                let info = res.data
                this.$Log("chatRoom", info)
                this.close()
                SetLocalStorage(CurrentContactCachetKey, JSON.stringify(info))
                this.$router.replace({
                    path: "/chat",
                })
            })
        }
    }
}
</script>

<style scoped lang="less">
.apply-notice-container {
    position: absolute;
    z-index: 999;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background: rgba(0, 0, 0, 0.9);

    .apply-notice-box {
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background: #FFFFFF;
        padding: 16px;
        width: calc(100% - 60px);
        height: 180px;
        margin: auto;
        border-radius: 14px;

        .avatar {
            width: 88px;
            height: 88px;
            border-radius: 4px;
        }
    }
}
</style>
