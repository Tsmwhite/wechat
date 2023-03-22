<template>
  <div class="chat-header">
    <div class="left">
      <van-icon class="icon" name="arrow-left" @click="back"/>
    </div>
    <div class="center">
      <div class="title" v-if="roomData.type === 1">{{ roomData.name }}（{{ roomData.member_num }}）</div>
      <div class="title" v-else>{{ roomData.remark || roomData.name }}</div>
    </div>
    <div class="right">
      <van-icon v-if="showMoreBtn" class="icon" name="ellipsis"/>
    </div>
  </div>
</template>

<script>
export default {
  name: "chat-header",
  props: {
    backFunc: {
      type: Function,
    },
    roomData: {
      type: Object,
    },
    showMoreBtn: {
      type: Boolean,
      default: true,
    }
  },
  data() {
    return {}
  },
  methods: {
    back() {
      if (typeof this.backFunc === "function") {
        this.backFunc()
      } else {
        this.$router.go(-1)
      }
    }
  }
}
</script>

<style scoped lang="less">
.chat-header {
  background: #f7f7f7;
  position: fixed;
  z-index: 999;
  opacity: 1;
  width: 100%;
  top: 0;
  left: 0;
  display: flex;
  justify-content: space-between;
  padding: 12px 0;
  border-bottom: 1px solid #d7d8d9;

  .icon {
    font-weight: bold;
    font-size: 16px;
    color: #666666
  }

  .left {
    margin-left: 12px;
  }

  .right {
    margin-right: 12px;
  }

  .center {
    .title {
      font-weight: 500;
    }
  }
}
</style>
