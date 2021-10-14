<template>
  <div style="padding:30px 60px">
    <div style="display: flex;align-items: center">
      <el-input v-model="input"
                style="width: 600px;margin: 15px"
                @keypress.native.enter="sendBtnClick"></el-input>
      <el-button type="primary" plain @click="sendBtnClick">OK</el-button>
      <el-button type="danger" plain @click="closeBtnClick">CLOSE</el-button>
    </div>
    <el-alert v-for="(item,index) in messageList"
              :key="index"
              :title="item.Text"
              :closable="false"
              center
              style="margin: 15px"
              :type="item.type"></el-alert>
  </div>
</template>

<script>
let websocket
export default {
  name: "test",
  data() {
    return {
      wsUri: "ws://127.0.0.1:8011/ws",
      input: "",
      token: "",
      messageList: [],
    }
  },
  mounted() {
    let token = this.$route.query.token || ''
    this.$WebSocket.MessageHandleCallback = (e) => {
      console.log('msg', e)
      e.type = "success"
      this.messageList.push(e)
    }
    this.$WebSocket.Init(token)
  },
  methods: {
    doSend(message) {
      this.writeToScreen("SENT: " + message, 'info');
      this.$WebSocket.send({
        Text: message,
        Type: 200,
        SecondType: 400
      });
    },

    writeToScreen(message, type) {
      this.messageList.push({
        Text: message,
        type: type,
      })
    },
    sendBtnClick() {
      this.doSend(this.input, 'success');
      this.input = '';
    },

    closeBtnClick() {
      websocket.close();
    }
  }
}
</script>

<style scoped>

</style>