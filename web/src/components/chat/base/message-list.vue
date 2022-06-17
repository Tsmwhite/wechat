<template>
    <div id="messageList" class="message-list">
        <van-list id="messageListScroll"
                  v-model="loading"
                  :finished="finished"
                  finished-text="没有更多了"
                  @load="loadHistory"
                  direction="up">
            <div v-for="(item,index) in messages"
                 :key="index"
                 :class="['message-box',{right:item.sender === currentUuid() }]">
                <div v-if="item.sender !== currentUuid()" class="avatar">
                    <img :src="getAvatar(item)">
                </div>
                <!--消息前置提示@start消息状态-发送中、发送失败-->
                <div v-if="item.sender === currentUuid()" class="berfore-tip">
                    <van-loading v-if="false" class="loading-icon" size="14"/>
                    <van-icon v-if="false" name="warning" class="fail-icon"/>
                </div>
                <!--消息前置提示@end-->
                <!--消息体@start-->
                <template v-if="item.second_type === $MsgType.Text">
                    <div class="text">{{ item.content }}</div>
                </template>
                <template v-else-if="item.second_type === $MsgType.Image">
                    <div class="image"><img :src="item.content"></div>
                </template>
                <!--消息体@end-->
                <div v-if="item.sender === currentUuid()" class="avatar">
                    <img :src="meAvatar">
                </div>
            </div>
        </van-list>
    </div>
</template>

<script>
import {GetCurrentUuid, GetUserInfo} from "../../../utis/cache";

export default {
    name: "message-list",
    props: {
        roomData: {
            type: Object,
        }
    },
    data() {
        return {
            loading: false,
            finished: true,
            refreshing: false,
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
            this.finished = true
            return list
        },
    },
    mounted() {
        this.init()
    },
    methods: {
        init() {
            this.loadData()
        },
        getAvatar() {
            if (this.roomData["is_private"] === 0) {
                return this.roomData.avatar ||  'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'
            }
            return  'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'
        },
        currentUuid() {
            return GetCurrentUuid()
        },
        loadHistory() {

        },
        loadData() {

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
    margin: 60px 0;
    height: calc(100% - 120px);
    overflow-y: scroll;

    .message-box {
        max-width: 100%;
        margin: 15px 2px;
        text-align: left;
        position: relative;
        display: flex;
        align-items: center;

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

        .berfore-tip {
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
