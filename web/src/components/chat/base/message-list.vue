<template>
    <div class="message-list">
        <van-pull-refresh v-model="refreshing" @refresh="loadHistory">
            <div v-for="(item,index) in data"
                 :key="index"
                 :class="['message-box',{right:item.Recipient !==0 }]">
                <div v-if="item.Recipient ===0" class="avatar">
                    <img v-lazy="'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'">
                </div>
                <!--消息前置提示@start消息状态-发送中、发送失败-->
                <div v-if="item.Recipient !==0" class="berfore-tip">
                    <van-loading v-if="false" class="loading-icon" size="14"/>
                    <van-icon v-if="true" name="warning" class="fail-icon"/>
                </div>
                <!--消息前置提示@end-->
                <!--消息体@start-->
                <template v-if="item.SecondType === $MsgType.Text">
                    <div class="text">{{ item.Content }}</div>
                </template>
                <template v-else-if="item.SecondType === $MsgType.Image">
                    <div class="image"><img :src="item.Content"></div>
                </template>
                <!--消息体@end-->
                <div v-if="item.Recipient !==0" class="avatar">
                    <img v-lazy="'https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png'">
                </div>
            </div>
        </van-pull-refresh>
    </div>
</template>

<script>
export default {
    name: "message-list",
    data() {
        return {
            loading: true,
            finished: false,
            refreshing: false,
            data: [],
        }
    },
    mounted() {
        this.init()
    },
    methods: {
        init() {
            setTimeout(() => {
                this.loadData()
                this.loading = false
                this.finished = true
            }, 1200)
        },
        loadHistory() {
            setTimeout(() => {
                for (let i = 0; i < 10; i++) {
                    this.data.unshift({
                        Recipient: i % 2 === 0 ? 0 : 1,
                        SendTime: '',
                        Content: "信息" + i,
                        Type: 200,
                        SecondType: 400,
                    })
                }
                this.refreshing = false
            })

        },
        loadData() {
            const contents = [
                "嗯哼",
                "怎么说",
                "明天一起去玩啊明天一起去玩啊明天一起去玩啊明天一起去玩啊明天一起去玩啊明天一起去玩啊明天一起去玩啊明天一起去玩啊",
                "????",
                "怎么样",
                "ok",
                "ok",
                "go go go",
                "lol 啊",
                "kkk",
                "kkk",
                "kkk",
                "kkk",
                "kkk",
                "kkk",
                "kkk",
                "go go go",
                "lol 啊",
                "kkk",
                "kkk",
                "kkk",
                "kkk",
                "go go go",
                "lol 啊",
                "kkk",
                "kkk",
                "kkk",
                "kkk",
            ]
            const message1 = {
                Recipient: 1,
                SendTime: '',
                Content: "https://lofter.lf127.net/1611802375723/xizhang3.png?imageView&type=jpg&quality=80&stripmeta=0&thumbnail=4000x4000",
                Type: 200,
                SecondType: 401,
            }
            const message2 = {
                Recipient: 0,
                SendTime: '',
                Content: "https://lofter.lf127.net/1611802375723/xizhang3.png?imageView&type=jpg&quality=80&stripmeta=0&thumbnail=4000x4000",
                Type: 200,
                SecondType: 401,
            }
            const message3 = {
                Recipient: 1,
                SendTime: '',
                Content: "嗯嗯嗯",
                Type: 200,
                SecondType: 400,
            }
            const message4 = {
                Recipient: 1,
                SendTime: '',
                Content: "嗯嗯嗯",
                Type: 200,
                SecondType: 400,
            }
            this.data.push(message1)
            this.data.push(message2)
            this.data.push(message3)
            this.data.push(message4)
            for (let i = 0; i < contents.length; i++) {
                this.data.push({
                    Recipient: i % 2 === 0 ? 0 : 1,
                    SendTime: '',
                    Content: contents[i],
                    Type: 200,
                    SecondType: 400,
                })
            }
        }
    }
}
</script>

<style scoped lang="less">
.message-list {
    margin-top: 60px;
    padding-bottom: 60px;
    height: 100%;
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
