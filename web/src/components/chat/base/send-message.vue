<template>
    <div class="send-message">
        <div class="top">
            <div class="left">
                <voice></voice>
            </div>
            <div class="center">
                <van-field rows="1"
                           v-model="message"
                           :autosize="{maxHeight:100}"
                           style="border-radius: 8px"
                           @blur="sendTextMessage"
                           type="textarea"/>
            </div>
            <div class="right">
                <emoji></emoji>
                <span @click="showExtra"><add></add></span>
            </div>
        </div>
        <div v-if="extraFlag"
             class="extra">
            <div class="extra-item" @click="sendVideoCall">
                <send-video></send-video>
                <span class="title">视频通话</span>
            </div>
        </div>

        <video-call :room-data="roomData" ref="videoCallRef"></video-call>
    </div>
</template>

<script>
import Voice from "@/components/icons/send-message-icons/voice";
import Emoji from "@/components/icons/send-message-icons/emoji";
import Add from "@/components/icons/send-message-icons/add";
import SendVideo from "@/components/icons/send-message-icons/send-video";
import VideoCall from "../extras/video-call";

export default {
    name: "send-message",
    components: {VideoCall, SendVideo, Add, Emoji, Voice},
    props: {
        roomData: {
            type: Object,
            required: true,
        }
    },
    data() {
        return {
            message: '',
            extraFlag: false,
        }
    },
    mounted() {
        this.$SetR(this.roomData.room)
        // let count = 1
        // const timer = setInterval(() => {
        //     this.message = "第"+count+"条消息"
        //     this.sendTextMessage()
        //     if (count > 100000) {
        //         clearInterval(timer)
        //     }
        //     count += 1
        // },100)
    },
    methods: {
        showExtra() {
            this.extraFlag = true
        },
        sendTextMessage() {
            this.message = this.message.trim()
            if (this.message === "") {
                return
            }
            this.$WebSocket.send(this.$Chat.text(this.message))
            this.$emit('sendMessage', this.message)
            this.message = ""
        },
        sendVideoCall() {
            this.$refs.videoCallRef.open()
        }
    }
}
</script>

<style scoped lang="less">
.send-message {
    width: 100%;

    .top {
        background: #f7f7f7;
        width: 100%;
        border-top: 1px solid #d7d8d9;
        display: flex;
        align-items: center;
        padding: 12px 5px;

        .center {
            flex: 2;
            margin: 0 4px;
        }

        .right {
            display: flex;
            align-items: center;
            margin-right: 12px;
        }
    }

    .extra {
        display: flex;
        margin: 12px;

        .extra-item {
            display: flex;
            flex-direction: column;
            align-items: center;
            padding: 6px;
            background: rgba(255, 255, 255, 0.65);;
            border-radius: 4px;

            .title {
                font-size: 12px;
                color: rgba(0, 0, 0, 0.65);
            }
        }
    }
}
</style>
