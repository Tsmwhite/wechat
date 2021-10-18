<template>
    <div class="chat-box">
        <chat-header></chat-header>
        <message-list></message-list>
    </div>
</template>

<script>
import ChatHeader from "../base/chat-header";
import MessageList from "./base/message-list";

export default {
    name: "chat-box",
    components: {MessageList, ChatHeader}
}
</script>

<style scoped>
    .chat-box {
        width: 100%;
        height: 100%;
    }
</style>
