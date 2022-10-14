<template>
    <div class="chat-box">
        <message-list :roomData="roomData" ref="msgListRef"></message-list>
        <send-message :roomData="roomData" @sendMessage="sendMsg"></send-message>
    </div>
</template>

<script>
import MessageList from "./message-list";
import SendMessage from "./send-message";

export default {
    name: "chat-box",
    components: {SendMessage, MessageList},
    props: {
        roomData: {
            type: Object,
        }
    },
    mounted() {
        this.$refs.msgListRef.scrollToBottom()
    },
    methods: {
        sendMsg() {
            // 发送完消息将聊天窗口滑动到底部
            this.$refs.msgListRef.scrollToBottom()
        },
    }
}
</script>

<style scoped>
.chat-box {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    overflow: hidden;
}
</style>
