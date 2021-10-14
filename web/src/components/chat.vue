<!--聊天组件-->
<template>
    <el-container style="max-height:600px;border: 1px solid #eee;position: relative">
        <!--好友列表@start-->
        <el-aside  style="background-color: rgb(238, 241, 246);height: 100%;position:fixed;left: 0;width: 20%">
            <el-menu :default-openeds="['1', '3']">
                <el-menu-item v-for="(item,key) in memberList"
                              @click="SelectChat(key)"
                              :key="item.id"
                              :class="item.id == newCaht.id ? 'active' : ''">
                    <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                    {{item.nickname}}
                </el-menu-item>
            </el-menu>
        </el-aside>
        <!--好友列表@end-->

        <!--聊天框@start-->
        <el-container style="position:fixed;right: 0%;width: 80%;" v-if="newCaht != null">
            <!--好友信息@start-->
            <el-header style="text-align: right; font-size: 12px;background: #eeeeee">
                <el-dropdown>
                    <i class="el-icon-setting" style="margin-right: 15px"></i>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item>查看</el-dropdown-item>
                        <el-dropdown-item>新增</el-dropdown-item>
                        <el-dropdown-item>删除</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                <span style="float: left">{{newCaht.nickname}}</span>
            </el-header>
            <!--好友信息@end-->

            <!--聊天信息@start-->
            <el-main style="height: 420px;border-bottom: 1px solid #eeeeee">
                <div class="chatContainer">
                    <p v-for="item in msgList"
                       :class="item.isMe == 1 ? 'right':'left'"
                       :key="item.id">
                        <span v-text="item.content"></span>
                    </p>
                </div>
            </el-main>
            <!--好友信息@end-->

            <inputReply></inputReply>

        </el-container>
        <!--聊天框@end-->

        <el-container style="position:fixed;right: 0%;width: 80%;" v-else>
            <img src="../assets/images/chatNull.jpg" style="width: 100%;margin: auto;"/>
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

    .chatContainer p{
        margin: 7px 20px !important;
        display: block;
        position: relative;
    }

    .chatContainer p span{
        background: #4fc08d;
        color: #eeeeee;
        padding: 5px 10px;
        /*margin: 15px;*/
        border-radius: 16px 16px;
        display: inline-block;
    }

    .chatContainer p.right{
       text-align: right;
       padding-left: 120px;
    }

    .chatContainer p.left{
        text-align: left;
        padding-right: 120px;
    }

    .chatContainer p.right span{
        background: #333333;
        text-align: left;
    }

    .el-menu-item.active {
        background: #c8e1f3;
    }
</style>

<script>
    import inputReply from "./inputReply"
    export default {
        name:"chat",
        components:{
            inputReply
        },
        methods:{
            SelectChat:function(index)  {
                //window.console.log(index)
                this.newCaht = this.memberList[index]
            },
        },
        data() {
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
                        content:"全功能ORM（几乎）\n" +
                            "关联（包含一个，包含多个，属于，多对多，多种包含）\n" +
                            "Callbacks（创建/保存/更新/删除/查找之前/之后）\n" +
                            "预加载（急加载）\n" +
                            "事务\n" +
                            "复合主键\n" +
                            "SQL Builder\n" +
                            "自动迁移\n" +
                            "日志\n" +
                            "可扩展，编写基于GORM回调的插件\n" +
                            "每个功能都有测试\n" +
                            "开发人员友好？",
                        date:"2020-01-26 18:45",
                        id:5,
                    },
                    {
                        isMe:1,
                        content:"package main\n" +
                            "\n" +
                            "import (\n" +
                            "    \"github.com/jinzhu/gorm\"\n" +
                            "    _ \"github.com/jinzhu/gorm/dialects/sqlite\"\n" +
                            ")\n" +
                            "\n" +
                            "type Product struct {\n" +
                            "  gorm.Model\n" +
                            "  Code string\n" +
                            "  Price uint\n" +
                            "}\n" +
                            "\n" +
                            "func main() {\n" +
                            "  db, err := gorm.Open(\"sqlite3\", \"test.db\")\n" +
                            "  if err != nil {\n" +
                            "    panic(\"连接数据库失败\")\n" +
                            "  }\n" +
                            "  defer db.Close()\n" +
                            "\n" +
                            "  // 自动迁移模式\n" +
                            "  db.AutoMigrate(&Product{？",
                        date:"2020-01-26 18:45",
                        id:4,
                    },
                    {
                        isMe:0,
                        content:"很高兴认识你",
                        date:"2020-01-26 18:45",
                        id:3
                    },
                    {
                        isMe:1,
                        content:"您好，很高兴认识你",
                        date:"2020-01-26 18:44",
                        id:2
                    },
                    {
                        isMe:1,
                        content:"您好，很高兴认识你",
                        date:"2020-01-26 18:44",
                        id:1
                    }
                ],
                //当前聊天好友
                newCaht: {
                    nickname: "清月",
                    id: 1,
                }
            }
        }
    };
</script>