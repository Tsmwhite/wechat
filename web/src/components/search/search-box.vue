<template>
    <div class="search-box">
        <van-tabbar v-model="active" :fixed="false">
            <van-tabbar-item icon="friends-o">好友</van-tabbar-item>
            <van-tabbar-item icon="user-o">用户</van-tabbar-item>
            <van-tabbar-item icon="chat-o">群聊</van-tabbar-item>
        </van-tabbar>
        <div class="search-list">
            <contact-item></contact-item>
        </div>
    </div>
</template>

<script>
import ContactItem from "../chat/base/contact-item";
export default {
    name: "search-box",
    components: {ContactItem},
    data() {
        return {
            active: 0,
        }
    }
}
</script>

<style scoped lang="less">
.search-list {
    padding: 16px;
    background: #f7f7f7;
}
</style>
