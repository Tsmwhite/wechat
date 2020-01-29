<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
  import Ws from "./socket";
  export default {
    name: 'app',
    components: {

    },
    created() {
      Ws.Init("theSmallWhiteMe")
    }

  }
</script>

<style>
#app {

}
</style>
