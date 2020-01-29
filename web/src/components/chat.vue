<!--聊天组件-->
<template>
    <el-container style="max-height:600px;border: 1px solid #eee;position: relative">
        <el-aside  style="background-color: rgb(238, 241, 246);height: 100%;position:fixed;left: 0;width: 20%">
            <el-menu :default-openeds="['1', '3']" style="padding-left: 12px">
                <el-menu-item v-for="item in memberList" :key="item.id">
                    <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                    {{item.nickname}}
                </el-menu-item>
            </el-menu>
        </el-aside>

        <el-container style="position:fixed;right: 0%;width: 80%;">
            <el-header style="text-align: right; font-size: 12px;background: #eeeeee">
                <el-dropdown>
                    <i class="el-icon-setting" style="margin-right: 15px"></i>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item>查看</el-dropdown-item>
                        <el-dropdown-item>新增</el-dropdown-item>
                        <el-dropdown-item>删除</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                <span style="float: left">王小虎</span>
            </el-header>

            <el-main style="height: 420px;border-bottom: 1px solid #eeeeee">
                <div class="chatContainer">
                    <p v-for="item in msgList" :key="item.id">
                       <span>
                           {{item.content}}
                       </span>
                    </p>
                    <p class="right">
                       <span>
                           你好
                       </span>
                    </p>
                </div>
            </el-main>

            <inputReply></inputReply>

        </el-container>
    </el-container>
</template>

<style>
    .el-header {
        background-color: #B3C0D1;
        color: #333;
        line-height: 60px;
    }

    .el-aside {
        color: #333;
    }
    .chatContainer p span{
        background: #4fc08d;
        color: #eeeeee;
        padding: 5px 10px;
        margin: 15px;
        border-radius: 16px 16px;
    }

    .chatContainer p.right{
        float: right;
    }
    .chatContainer p.right span{
        background: #333333;
    }

</style>

<script>
    import inputReply from "./inputReply"
    export default {
        name:"chat",
        components:{
            inputReply
        },
        data() {
            const item = {
                date: '2016-05-02',
                name: '王小虎',
                address: '上海市普陀区金沙江路 1518 弄'
            };
            return {
                //好友列表
                memberList:[
                    {
                        nickname:"清月",
                        id:1,
                    },
                    {
                        nickname:"asunder",
                        id:2,
                    },
                    {
                        nickname:"summer",
                        id:3,
                    },
                    {
                        nickname:"芥末",
                        id:4,
                    },
                    {
                        nickname:"离恨",
                        id:5,
                    },
                    {
                        nickname:"哈登",
                        id:6,
                    },
                    {
                        nickname:"科比",
                        id:7,
                    }
                ],
                //消息列表
                msgList:[
                    {
                        isMe:0,
                        content:"你有什么事情吗？",
                        date:"2020-01-26 18:45",
                        id:4,
                    },
                    {
                        isMe:1,
                        content:"请问现在方便吗？",
                        date:"2020-01-26 18:45",
                        id:3,
                    },
                    {
                        isMe:0,
                        content:"很高兴认识你",
                        date:"2020-01-26 18:45",
                        id:2
                    },
                    {
                        isMe:1,
                        content:"您好，很高兴认识你",
                        date:"2020-01-26 18:44",
                        id:1
                    }
                ],
                tableData: Array(20).fill(item)
            }
        }
    };
</script>