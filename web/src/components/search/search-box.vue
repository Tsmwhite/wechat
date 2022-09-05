<template>
    <div class="search-box">
        <van-search :show-action="showSearchFlag"
                    placeholder="搜索"
                    @focus="showSearchFlag = true"
                    @cancel="showSearchFlag = false"/>
        <van-popup v-model="showSearchFlag"
                   position="top"
                   style="width: 100%;height: 60%;">
            <van-search v-model="keyword"
                        @search="search"
                        placeholder="搜索"/>
            <div class="content">
                <van-tabbar v-model="active"
                            @change="tabChange"
                            :fixed="false">
                    <van-tabbar-item icon="friends-o">好友</van-tabbar-item>
                    <van-tabbar-item icon="user-o">用户</van-tabbar-item>
                    <van-tabbar-item icon="chat-o">群聊</van-tabbar-item>
                </van-tabbar>
                <div v-if="list.length > 0"
                     class="search-list">
                    <template v-for="(item,index) in list">
                        <contact-item :contact="item" :key="index">
                            <template #extra>
                                <van-button v-if="active === 1" type="info" size="small" @click="addFriend(item)">加好友</van-button>
                                <van-button v-else-if="active === 2" type="info" size="small" @click="addGroup(item)">加群</van-button>
                                <van-button v-else type="info" size="small" @click="linkChat">发消息</van-button>
                            </template>
                        </contact-item>
                    </template>
                </div>
                <van-empty v-else image="search" description="暂无数据"/>
            </div>
        </van-popup>
        <add-friend-modal :user-info="addUserInfo" ref="addFriendRef"></add-friend-modal>
    </div>
</template>

<script>
import ContactItem from "../chat/base/contact-item";
import {searchFriends, searchRoom, searchUser} from "../../api/common";
import AddFriendModal from "../add-friend-modal";

export default {
    name: "search-box",
    components: {AddFriendModal, ContactItem},
    data() {
        return {
            showSearchFlag: false,
            active: 0,
            keyword: "",
            friends: [],
            users: [],
            rooms: [],
            existSearchMap: {
                0: "",
                1: "",
                2: "",
            },
            addUserInfo: {},
        }
    },
    watch: {
        keyword() {
            this.keyword = this.keyword.trim()
        }
    },
    computed: {
        list() {
            switch (this.active) {
                case 0:
                    return this.friends
                case 1:
                    return this.users
                case 2:
                    return this.rooms
            }
            return []
        }
    },
    methods: {
        tabChange() {
            if (this.keyword !== "" &&
                this.existSearchMap[this.active] !== this.keyword) {
                this.search()
            }
        },
        search() {
            if (this.keyword === "") {
                return
            }
            switch (this.active) {
                case 0:
                    this.searchFriends()
                    break
                case 1:
                    this.searchUser()
                    break
                case 2:
                    this.searchRoom()
                    break
            }
        },
        searchFriends() {
            searchFriends({
                keyword: this.keyword
            }).then(res => {
                this.friends = res.data || []
                this.existSearchMap[0] = this.keyword
            })
        },
        searchUser() {
            searchUser({
                keyword: this.keyword
            }).then(res => {
                this.users = res.data || []
                this.existSearchMap[1] = this.keyword
            })
        },
        searchRoom() {
            searchRoom({
                keyword: this.keyword
            }).then(res => {
                this.rooms = res.data || []
                this.existSearchMap[2] = this.keyword
            })
        },
        addFriend(user) {
            this.addUserInfo = user
            this.$refs.addFriendRef.show()
        },
        addGroup() {

        },
        linkChat() {

        },
    }
}
</script>

<style scoped lang="less">
.content {
    width: 100%;
}

.search-list {
    padding: 16px;
    background: #f7f7f7;
}
</style>
