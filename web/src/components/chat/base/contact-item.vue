<template>
    <div class="contact-option">
        <div class="left">
            <div :class="['avatar',{unread:contact.unread_count > 0},{'only-tip':contact.only_tip}]"
                 :data-unread="contact.unread_count > 99 ? 99 : contact.unread_count">
                <img v-if="contact.avatar" :src="contact.avatar">
                <img v-else src="@/assets/default-group-avatar.jpeg"/>
            </div>
        </div>
        <div class="right">
            <div class="info">
                <div class="nickname">{{ contact.name || contact.title }}</div>
                <div class="last-msg">{{ contact.last_msg }}</div>
            </div>
            <div class="extra">
                <slot name="extra">
                    <div class="last-time">{{ contact.last_time }}</div>
                </slot>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "contact-item",
    props: {
        contact: {
            type: Object,
            default:() => {
                return {
                    name: "小马",
                    avatar: "https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png",
                    unread_count: 0,
                    only_tip: false,
                    last_msg: "",
                    last_time: "",
                }
            }
        }
    },
    methods: {

    }
}
</script>

<style scoped lang="less">
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
</style>
