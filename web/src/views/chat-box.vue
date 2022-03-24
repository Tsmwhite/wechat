<template>
    <div class="chat-box">
        <chat-header></chat-header>
        <message-list></message-list>
        <send-message></send-message>
    </div>
</template>

<script>
import ChatHeader from "@/components/base/chat-header";
import MessageList from "@/components/chat/base/message-list";
import SendMessage from "@/components/chat/base/send-message";

export default {
    name: "chat-box",
    components: {SendMessage, MessageList, ChatHeader}
}
</script>

<style scoped>
    .chat-box {
        width: 100%;
        height: 100%;
    }
</style>
