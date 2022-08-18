<template>
    <div class="h100">
        <chat-header :room-data="currentRoom"
                     :back-func="goBack"></chat-header>
        <chat-box :room-data="currentRoom"></chat-box>
    </div>
</template>

<script>
import {CurrentContactCachetKey, GetLocalStorage} from "../utis/cache";
import ChatHeader from "../components/chat/base/chat-header";
import ChatBox from "../components/chat/base/chat-box";

export default {
    name: "chat",
    components: {ChatBox, ChatHeader},
    data() {
        return {
            currentRoom: {},
        }
    },
    created() {
        const currentRoom = GetLocalStorage(CurrentContactCachetKey)
        if (!currentRoom) {
            this.$toast.fail("系统错误")
            return
        }
        this.currentRoom = JSON.parse(currentRoom)
    },
    methods: {
        goBack() {
            this.$router.push({
                path: "/index"
            })
        }
    }
}
</script>

<style scoped>

</style>
