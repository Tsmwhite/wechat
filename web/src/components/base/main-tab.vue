<template>
    <div class="main-tab">
        <div>
            <van-icon class="icon" name="chat" />
        </div>
        <div>
            <van-icon class="icon" name="friends-o" />
        </div>
        <div>
            <van-icon class="icon" name="fire-o" />
        </div>
        <div>
            <van-icon class="icon" name="manager-o" />
        </div>
    </div>
</template>

<script>
export default {
    name: "main-tab"
}
</script>

<style scoped lang="less">
.main-tab {
    position: fixed;
    bottom: 0;
    left: 0;
    display: flex;
    width: 100%;
    justify-content: space-between;
    border-top: 1px solid #e1e6eb;
    background: #f7f7f7;
    > div {
        margin: 12px 24px;
        .van-icon {
            font-size: 22px;
        }
    }
}
</style>
