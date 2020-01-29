<template>
   <div>
       <Header></Header>
       <Chat></Chat>
   </div>
</template>

<script>
    import Header from "../components/header"
    import Chat   from "../components/chat"

    export default {
        name: 'index',
        components: {
            Header,
            Chat
        }
    }
</script>

<style scoped>

</style>