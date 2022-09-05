<template>
    <div>
        <chat-header :room-data="currentRoom" :show-more-btn="false"></chat-header>
        <div id="messageList" class="message-list">
            <van-list id="messageListScroll"
                      v-model="loading"
                      :finished="finished"
                      finished-text="没有更多了"
                      @load="loadHistory"
                      :offset="80"
                      direction="up">
                <div v-for="(item,index) in messages"
                     :key="index"
                     class="message-box">
                    <div class="header">
                        <div class="left">
                            <div class="user-info">
                                <img class="avatar" :src="getUserInfo(item).avatar || defaultAvatar"/>
                                <span class="name">{{ getUserInfo(item).name }} </span>
                            </div>
                        </div>
                        <div class="right">
                            <span style="color:rgba(0, 0, 0, 0.45)">申请添加好友</span>
                        </div>
                    </div>
                    <div class="content">
                        <div>
                            <label>对方留言：</label>
                            <span>
                            {{ item.content }}
                        </span>
                        </div>
                    </div>
                    <template v-if="item.status == 0">
                        <div class="content">
                            <div>
                                <label>设置备注：</label>
                                <span>
                                <van-field v-model="item._remark"
                                           placeholder="请输入备注（可选）"
                                           :maxlength="12"
                                           show-word-limit/>
                            </span>
                            </div>
                        </div>
                        <div class="footer">
                            <van-button size="small" type="default" style="width: calc(50% - 12px);"
                                        @click="handleApply(item,2)">拒绝
                            </van-button>
                            <van-button size="small" type="info" style="width: calc(50% - 12px);"
                                        @click="handleApply(item,1)">同意
                            </van-button>
                        </div>
                    </template>
                    <template v-else-if="item.status == 1">
                        <div class="status">
                            <van-tag type="primary" size="large">已同意</van-tag>
                        </div>
                    </template>
                    <template v-else-if="item.status == 2">
                        <div class="status">
                            <van-tag type="primary" size="large" color="#ec6969">已拒绝</van-tag>
                        </div>
                    </template>
                    <template v-else>
                        <div class="status">
                            <van-tag type="primary" size="large" color="rgba(0,0,0,.25)">已过期</van-tag>
                        </div>
                    </template>
                </div>
            </van-list>
        </div>
    </div>
</template>

<script>
import ChatHeader from "../components/chat/base/chat-header";
import {getFriendNotice, handleAddFriendApply} from "../api/common";
import {GetUserByUuid} from "../utis/cookie";

export default {
    name: "friend-notice",
    components: {
        ChatHeader
    },
    data() {
        return {
            defaultAvatar: 'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png',
            currentRoom: {
                type: 0,
                name: '好友通知',
            },
            loading: false,
            finished: false,
            messages: [],
            pagination: {
                current: 1,
                pageSize: 10,
                total: 0,
                loadLastId: 0,
            },
        }
    },
    computed: {
        getUserInfo() {
            return (item) => {
                let storeData = this.$store.state
                let user = storeData.friend.FriendMap[item.sender]
                console.log("user",user)
                if (!user) {
                    GetUserByUuid({user_id: item.sender})
                    return {
                        avatar: this.defaultAvatar,
                        name: '...',
                    }
                }
                return user
            }
        },
    },
    methods: {
        loadHistory() {
            getFriendNotice({
                size: this.pagination.pageSize,
                last_id: this.pagination.loadLastId,
            }).then(res => {
                let len = res.data.length
                if (len < 1) {
                    this.finished = true
                    this.loading = false
                    return
                }
                const container = document.getElementById('messageList')
                let oldHeight = container.scrollHeight
                this.pagination.loadLastId = res.data[len - 1].id || 0
                res.data.map(item => {
                    item._remark = ""
                })
                this.messages.push(...res.data.reverse())
                this.$nextTick(() => {
                    if (this.pagination.current === 1) {
                        this.scrollToBottom()
                    } else {
                        container.scrollTop = (container.scrollHeight - oldHeight)
                    }
                    this.pagination.current++
                    setTimeout(() => {
                        this.loading = false
                    }, 3000)
                })

            }).catch(() => {
                this.loading = false
            })
        },
        handleApply(item, status) {
            handleAddFriendApply({
                uuid: item.uuid,
                status: status,
                remark: item._remark,
            }).then(res => {
                this.$toast.success(res.msg)
            })
        },
        scrollToBottom() {
            this.$nextTick(() => {
                const container = document.getElementById('messageList')
                container.scrollTop = container.scrollHeight
            })
        },
    }
}
</script>

<style scoped lang="less">
.message-box {
    margin: 24px;
    border-radius: 6px;
    border: 1px solid #e8e8e8;
    padding: 16px;
    background: #ffffff;

    .header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding-bottom: 16px;
        border-bottom: 1px solid #e8e8e8;
    }

    .content {
        margin-top: 16px;
        text-align: left;

        > label {
            color: rgba(0, 0, 0, 0.45);
        }
    }

    .footer {
        display: flex;
        justify-content: space-between;
        padding-top: 16px;
        margin-top: 16px;
    }

    .status {
        margin-top: 16px;
        text-align: right;
    }
}

.user-info {
    display: flex;
    align-items: center;

    .avatar {
        margin-right: 8px;
        width: 36px;
        height: 36px;
        border-radius: 6px;
    }

}
</style>
