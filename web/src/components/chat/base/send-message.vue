<template>
    <div class="send-message">
        <div class="left">
            <voice></voice>
        </div>
        <div class="center">
            <van-field rows="1"
                       v-model="message"
                       :autosize="{maxHeight:100}"
                       style="border-radius: 8px"
                       type="textarea"/>
        </div>
        <div class="right">
            <emoji></emoji>
            <add></add>
        </div>
    </div>
</template>

<script>
import Voice from "@/components/icons/send-message-icons/voice";
import Emoji from "@/components/icons/send-message-icons/emoji";
import Add from "@/components/icons/send-message-icons/add";

export default {
    name: "send-message",
    components: {Add, Emoji, Voice},
    data() {
        return {
            message: ''
        }
    }
}
</script>

<style scoped lang="less">
.send-message {
    position: fixed;
    bottom: 0;
    left: 0;
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
</style>
