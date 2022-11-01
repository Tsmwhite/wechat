<template>
    <div id="messageList" class="message-list">
        <van-list id="messageListScroll"
                  v-model="loading"
                  :finished="finished"
                  finished-text="没有更多了"
                  @load="loadHistory"
                  :offset="80"
                  direction="up">
            <template v-for="(item,index) in messages">
                <div v-if="item.second_type === MessageTypes.videoCall"
                     :key="index">
                   <div style="font-size: 12px;color: rgba(0,0,0,0.65)"> {{item.sender === currentUuid() ? '我' : '对方'}}发起了视频通话【已结束】</div>
                </div>
                <div v-else
                     :key="index"
                     :class="['message-box',{right:item.sender === currentUuid() }]">
                    <div v-if="item.sender !== currentUuid()"
                         class="avatar">
                        <img :src="getAvatar(item)">
                    </div>
                    <!--消息前置提示@start消息状态-发送中、发送失败-->
                    <div v-if="item.sender === currentUuid()"
                         class="before-tip">
                        <van-loading v-if="false" class="loading-icon" size="14"/>
                        <van-icon v-if="false" name="warning" class="fail-icon"/>
                    </div>
                    <!--消息前置提示@end-->
                    <!--消息体@start-->
                    <template v-if="item.second_type === MessageTypes.text">
                        <div class="text" v-html="item.content"></div>
                    </template>
                    <template v-else-if="item.second_type === MessageTypes.image">
                        <div class="image"><img :src="item.content"></div>
                    </template>
                    <!--消息体@end-->
                    <div v-if="item.sender === currentUuid()"
                         class="avatar">
                        <img :src="meAvatar">
                    </div>
                </div>
            </template>
        </van-list>
    </div>
</template>

<script>
import {GetCurrentUuid, GetUserInfo} from "../../../utis/cache";
import {GetUserByUuid} from "../../../utis/cookie";
import {getHistory} from "../../../api/common";
import {MessageTypes} from "../../../library/message/const";

export default {
    name: "message-list",
    props: {
        roomData: {
            type: Object,
        }
    },
    data() {
        return {
            MessageTypes,
            loading: false,
            loadingController: false,
            finished: false,
            refreshing: false,
            loadLastId: 0,
            pagination: {
                current: 1,
                pageSize: 60,
                total: 0,
            },
        }
    },
    computed: {
        meAvatar() {
            return GetUserInfo().avatar || 'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'
        },
        messages() {
            let storeData = this.$store.state
            let roomId = this.roomData.room || 0
            let list = storeData.msg.MessageMapList[roomId]
            if (!list) {
                return []
            }
            return list
        },
        getAvatar() {
            return (item) => {
                let storeData = this.$store.state
                let user = storeData.friend.FriendMap[item.sender]
                if (!user) {
                    GetUserByUuid({user_id: item.sender})
                    return 'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'
                }
                return user.avatar || 'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'
            }
        },
    },
    mounted() {
        //this.init()
    },
    methods: {
        init() {
            this.finished = false
            let storeData = this.$store.state
            let roomId = this.roomData.room || 0
            let list = storeData.msg.MessageMapList[roomId]
            if ((!list || list.length < 1) && roomId > 0) {
                this.loadHistory()
            } else  {
                this.finished = true
            }
        },
        currentUuid() {
            return GetCurrentUuid()
        },
        loadHistory() {
            if (!this.roomData.room) {
                this.finished = true
            }
            // if (this.loadingController) {
            //     return
            // }
            // this.loadingController = true
            getHistory({
                room_uuid: this.roomData.room,
                //page: this.pagination.current,
                size: this.pagination.pageSize,
                last_id: this.loadLastId,
            }).then(res => {
                let len = res.data.length
                if (len < 1) {
                    this.finished = true
                    this.loading = false
                    return
                }
                const container = document.getElementById('messageList')
                let oldHeight = container.scrollHeight
                this.loadLastId = res.data[len - 1].id || 0
                this.$store.dispatch("history", {
                    recipient: this.roomData.room,
                    messages: res.data.reverse()
                })
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
.message-list {
    //margin: 60px 0;
    //height: calc(100% - 60px);
    margin-top: 50px;
    overflow-y: scroll;
    flex: 1;

    .message-box {
        max-width: 100%;
        margin: 15px 2px;
        text-align: left;
        position: relative;
        display: flex;
        //align-items: center;

        > div {
            display: inline-block;
        }

        .avatar {
            margin: 0 5px;
            width: 36px;
            height: 36px;
        }

        .avatar > img {
            border-radius: 6px;
            width: 100%;
            height: 100%;
        }

        .before-tip {
            margin-right: 5px;

            .fail-icon {
                color: red;
            }

            .loading-icon {
                display: inline-block
            }
        }

        .text {
            background: #FFFFFF;
            border-radius: 6px;
            padding: 8px 12px;
            font-size: 14px;
            max-width: 60%;
            text-align: left;
            overflow: hidden;
            white-space: break-spaces;
            word-break: break-all;
        }

        .image {
            max-width: 60%;
            padding: 5px;

            > img {
                border-radius: 8px;
                max-width: 100%;
            }
        }
    }

    .message-box.right {
        justify-content: flex-end;

        .text {
            background: #bae4ff;
        }
    }
}
</style>
