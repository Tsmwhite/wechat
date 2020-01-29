<!--消息输入组件-->
<template>
    <div>
        <el-main>
            <div class="replyContainer">
                <div>
                    <i class="el-icon-picture" title="添加图片"></i>
                    <i class="el-icon-camera-solid" title="拍摄照片"></i>
                    <i class="el-icon-folder-add"  title="添加文件"></i>
                    <i class="el-icon-microphone" title="发送语言"></i>
                    <i class="el-icon-video-camera" title="发送视频"></i>
                    <i class="el-icon-delete-solid" title="发送视频"></i>
                    <i class="el-icon-s-tools" title="发送视频"></i>
                </div>
                <textarea class="msgSendInput" maxlength="1200"></textarea>
            </div>
        </el-main>
    </div>
</template>

<script>
    export default {
        name: "inputReply",
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    .replyContainer i {
        margin: 12px;
        font-size: 22px
    }

    .msgSendInput{
        width: 90%;
        height: 140px;
        border: none;
    }
</style>