<template>
   <div>
       <HeaderAssembly></HeaderAssembly>
       <Chat></Chat>
   </div>
</template>

<script>
    import HeaderAssembly from "../components/headerassembly"
    import Chat   from "../components/chat"

    export default {
        name: 'index',
        components: {
            HeaderAssembly,
            Chat
        }
    }
</script>

<style scoped>

</style>