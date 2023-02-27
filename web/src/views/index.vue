<template>
    <div class="h100">
        <div class="contact-list">
            <search-box></search-box>
            <van-list v-model="loading"
                      :finished="finished"
                      @load="loadData">
                <div class="contact-option" @click="openFriendNoticeBox">
                    <div class="left">
                        <div class="avatar">
                            <img src="@/assets/friend-notice-icon.png"/>
                        </div>
                    </div>
                    <div class="right">
                        <div class="info">
                            <div class="nickname">好友通知</div>
                            <div class="last-msg"></div>
                        </div>
                        <div class="extra">
                            <div class="last-time"></div>
                        </div>
                    </div>
                </div>
                <div class="contact-option"
                     v-for="(item,index) in data"
                     @click="openChat(item)"
                     :key="index">
                    <div class="left">
                        <div :class="['avatar',{unread:item.unreadCount > 0},{'only-tip':item.onlyTip}]"
                             :data-unread="item.unreadCount > 99 ? 99 : item.unreadCount">
                            <img v-if="item.avatar" v-lazy="item.avatar">
                            <img v-else v-lazy="groupDefaultAvatar">
                        </div>
                    </div>
                    <div class="right">
                        <div class="info">
                            <div class="nickname">{{ item.name }}</div>
                            <div class="last-msg">{{ item.lastMsg }}</div>
                        </div>
                        <div class="extra">
                            <div class="last-time">{{ item.lastTime }}</div>
                        </div>
                    </div>
                </div>
            </van-list>
        </div>
    </div>
</template>

<script>
import api from "../api/common"
import ChatBox from "../components/chat/base/chat-box";
import ChatHeader from "../components/chat/base/chat-header";
import SearchBox from "../components/search/search-box";
import groupDefaultAvatar from "../assets/default-group-avatar.jpeg"
import {CurrentContactCachetKey, SetLocalStorage} from "../utis/cache";
import {renderAvatar} from "../utis/tools"

export default {
    name: "index",
    components: {SearchBox, ChatHeader, ChatBox},
    data() {
        return {
            groupDefaultAvatar,
            renderAvatar,
            value: "",
            loading: true,
            finished: false,
            data: [],
            chat: {
                show: false,
                currentRoom: {},
            }
        }
    },
    mounted() {
        // let token = this.$route.query.token
        // this.$WebSocket.Init(token)
        // SetToken(token)
        this.init()
        this.$WebSocket.Detect()
    },
    methods: {
        init() {
            this.loadData()
        },
        loadData() {
            api.getContacts().then(res => {
                this.data = res.data
                this.loading = false
                this.finished = true
            })
        },
        openChat(item) {
            console.log("chatRoom", item)
            SetLocalStorage(CurrentContactCachetKey, JSON.stringify(item))
            this.$router.push({
                path: "/chat"
            })
        },
        openFriendNoticeBox() {
            this.$router.push({
                path: "/friend-notice"
            })
        },
    }
}
</script>

<style scoped lang="less">
.contact-list {
    padding-bottom: 60px;
    .contact-option {
        max-width: 100%;
        overflow: hidden;
        display: flex;
        align-items: center;
        padding: 6px 6px 0 6px;

        .left {
            display: flex;

            .avatar {
                position: relative;
            }

            .avatar.unread.only-tip::before {
                content: "";
                display: inline-block;
                width: 8px;
                height: 8px;
                border-radius: 4px;
                background: red;
                position: absolute;
                right: -2px;
                top: -2px;
            }

            .avatar.unread::before {
                content: attr(data-unread);
                display: inline-block;
                width: 18px;
                height: 18px;
                border-radius: 9px;
                background: red;
                position: absolute;
                right: -4px;
                top: -4px;
                font-size: 10px;
                font-weight: 500;
                color: #FFFFFF;
                text-align: center;
                line-height: 18px;
            }

            .avatar > img {
                width: 44px;
                height: 44px;
                border-radius: 5px;
            }
        }

        .right {
            display: flex;
            flex-grow: 1;
            width: 100%;
            margin-left: 16px;
            padding: 10px 0;
            border-bottom: 1px solid #e1e6eb;
            justify-content: space-between;

            .info {
                display: flex;
                align-items: start;
                justify-content: center;
                flex-direction: column;
                overflow: hidden;
                max-width: 200px;

                .nickname {
                    font-size: 14px;
                }

                .last-msg {
                    font-size: 12px;
                    color: #666666;
                    margin-top: 6px;
                    width: 100%;
                    white-space: nowrap;
                    overflow: hidden;
                    text-overflow: ellipsis;
                }
            }

            .extra {
                padding-top: 5px;
                padding-right: 5px;

                .last-time {
                    font-size: 12px;
                }
            }
        }
    }
}
</style>
