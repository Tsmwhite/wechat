webpackJsonp([2],{"6Ch5":function(t,e){},"7paM":function(t,e,n){"use strict";var i=n("9GPK"),a=n("y0/c"),s=n("t5DY"),r=n("++XJ"),o=n("oqQY"),c=n.n(o),d={name:"message-list",props:{roomData:{type:Object}},data:function(){return{MessageTypes:r.b,loading:!1,loadingController:!1,finished:!1,refreshing:!1,loadLastId:0,pagination:{current:1,pageSize:60,total:0},lastSendMsgTime:null}},computed:{meAvatar:function(){return Object(i.e)().avatar||"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png"},messages:function(){var t=this.$store.state,e=this.roomData.room||0,n=t.msg.MessageMapList[e];return n||[]},getAvatar:function(){var t=this;return function(e){var n=t.$store.state.friend.FriendMap[e.sender];return n?n.avatar||"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png":(Object(a.a)({user_id:e.sender}),"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png")}}},mounted:function(){},methods:{init:function(){this.finished=!1;var t=this.$store.state,e=this.roomData.room||0,n=t.msg.MessageMapList[e];(!n||n.length<1)&&e>0?this.loadHistory():this.finished=!0},currentUuid:function(){return Object(i.b)()},loadHistory:function(){var t=this;this.roomData.room||(this.finished=!0),Object(s.e)({room_uuid:this.roomData.room,size:this.pagination.pageSize,last_id:this.loadLastId}).then(function(e){var n=e.data.length;if(n<1)return t.finished=!0,void(t.loading=!1);var i=document.getElementById("messageList"),a=i.scrollHeight;t.loadLastId=e.data[n-1].id||0,t.$store.dispatch("history",{recipient:t.roomData.room,messages:e.data.reverse()}),t.$nextTick(function(){1===t.pagination.current?t.scrollToBottom():i.scrollTop=i.scrollHeight-a,t.pagination.current++,setTimeout(function(){t.loading=!1},3e3)})}).catch(function(){t.loading=!1})},scrollToBottom:function(){this.$nextTick(function(){var t=document.getElementById("messageList");t.scrollTop=t.scrollHeight})},renderTime:function(t,e){var n=c()().unix();if(0===e)n-t.send_time;else{t.send_time-this.messages[e-1].send_time;var i=c()(1e3*t.send_time).format("YYYY-MM-DD HH:mm");if(c()(1e3*this.messages[e-1].send_time).format("YYYY-MM-DD HH:mm")===i)return""}return'<div class="send-msg-time">'+(n-t.send_time>86400?c()(1e3*t.send_time).format("YYYY-MM-DD HH:mm"):c()(1e3*t.send_time).format("HH:mm"))+"</div>"}}},l={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"message-list",attrs:{id:"messageList"}},[n("van-list",{attrs:{id:"messageListScroll",finished:t.finished,"finished-text":"没有更多了",offset:80,direction:"up"},on:{load:t.loadHistory},model:{value:t.loading,callback:function(e){t.loading=e},expression:"loading"}},[t._l(t.messages,function(e,i){return[e.second_type===t.MessageTypes.videoCall?n("div",{key:i,staticClass:"video-call-box"},[n("div",{staticStyle:{"font-size":"12px",color:"rgba(0,0,0,0.65)"}},[t._v("\n                    "+t._s(e.sender===t.currentUuid()?"我":"对方")+"发起了视频通话【已结束】\n                ")])]):[n("div",{domProps:{innerHTML:t._s(t.renderTime(e,i))}}),t._v(" "),n("div",{key:i,class:["message-box",{right:e.sender===t.currentUuid()}]},[e.sender!==t.currentUuid()?n("div",{staticClass:"avatar"},[n("img",{attrs:{src:t.getAvatar(e)}})]):t._e(),t._v(" "),e.sender===t.currentUuid()?n("div",{staticClass:"before-tip"},[t._e(),t._v(" "),t._e()],1):t._e(),t._v(" "),e.second_type===t.MessageTypes.text?[n("div",{staticClass:"text",domProps:{innerHTML:t._s(e.content)}})]:e.second_type===t.MessageTypes.image?[n("div",{staticClass:"image"},[n("img",{attrs:{src:e.content}})])]:t._e(),t._v(" "),e.sender===t.currentUuid()?n("div",{staticClass:"avatar"},[n("img",{attrs:{src:t.meAvatar}})]):t._e()],2)]]})],2)],1)},staticRenderFns:[]};var u=n("VU/8")(d,l,!1,function(t){n("n5u8"),n("6Ch5")},"data-v-3c76b01e",null).exports,h={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1647771122497",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"16033",width:"32",height:"32"}},[e("path",{attrs:{d:"M410.208 319.776a96 96 0 0 1 192 0v192.448a96 96 0 0 1-96 95.744 96 96 0 0 1-96-95.744V319.776z m96 352.192c88.192 0 160-71.648 160-159.744V319.776a160.096 160.096 0 0 0-160-159.776c-88.224 0-160 71.68-160 159.776v192.448a160.032 160.032 0 0 0 160 159.744zM742.4 448a32 32 0 0 0-32 32v24.672c0 116.192-94.72 210.688-211.2 210.688-116.448 0-211.2-94.496-211.2-210.688V480a32 32 0 0 0-64 0v24.672c0 140.16 105.76 255.904 241.76 272.448V864a32 32 0 0 0 64 0v-86.432C667.168 762.304 774.4 645.824 774.4 504.672V480a32 32 0 0 0-32-32","p-id":"16034"}})])])},staticRenderFns:[]};var f=n("VU/8")({name:"voice"},h,!1,function(t){n("XPRr")},"data-v-29829874",null).exports,m={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1647770199888",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"4299",width:"32",height:"32"}},[e("path",{attrs:{d:"M510.944 960c-247.04 0-448-200.96-448-448s200.992-448 448-448 448 200.96 448 448-200.96 448-448 448z m0-832c-211.744 0-384 172.256-384 384s172.256 384 384 384 384-172.256 384-384-172.256-384-384-384z",fill:"","p-id":"4300"}}),e("path",{attrs:{d:"M512 773.344c-89.184 0-171.904-40.32-226.912-110.624-10.88-13.92-8.448-34.016 5.472-44.896 13.888-10.912 34.016-8.48 44.928 5.472 42.784 54.688 107.136 86.048 176.512 86.048 70.112 0 134.88-31.904 177.664-87.552 10.784-14.016 30.848-16.672 44.864-5.888 14.016 10.784 16.672 30.88 5.888 44.864C685.408 732.32 602.144 773.344 512 773.344zM368 515.2c-26.528 0-48-21.472-48-48v-64c0-26.528 21.472-48 48-48s48 21.472 48 48v64c0 26.496-21.504 48-48 48zM656 515.2c-26.496 0-48-21.472-48-48v-64c0-26.528 21.504-48 48-48s48 21.472 48 48v64c0 26.496-21.504 48-48 48z",fill:"","p-id":"4301"}})])])},staticRenderFns:[]};var v=n("VU/8")({name:"emoji"},m,!1,function(t){n("CczY")},"data-v-7ff93387",null).exports,g={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1647770134907",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2981",width:"32",height:"32"}},[e("path",{attrs:{d:"M512 958.016c-119.648 0-232.128-46.368-316.736-130.56C110.624 743.2 64 631.2 64 512c0-119.168 46.624-231.2 131.232-315.424 84.608-84.192 197.088-130.56 316.736-130.56s232.128 46.368 316.704 130.56c84.672 84.224 131.264 196.256 131.264 315.392 0.032 119.2-46.592 231.232-131.264 315.456C744.128 911.616 631.648 958.016 512 958.016zM512 129.984c-102.624 0-199.072 39.744-271.584 111.936C167.936 314.048 128 409.984 128 512c0 102.016 39.904 197.952 112.384 270.048 72.512 72.192 168.96 111.936 271.584 111.936 102.592 0 199.072-39.744 271.584-111.936 72.48-72.16 112.416-168.064 112.384-270.08 0-102.016-39.904-197.92-112.384-270.016C711.072 169.76 614.592 129.984 512 129.984z","p-id":"2982"}}),e("path",{attrs:{d:"M736 480l-192 0L544 288c0-17.664-14.336-32-32-32s-32 14.336-32 32l0 192L288 480c-17.664 0-32 14.336-32 32s14.336 32 32 32l192 0 0 192c0 17.696 14.336 32 32 32s32-14.304 32-32l0-192 192 0c17.696 0 32-14.336 32-32S753.696 480 736 480z","p-id":"2983"}})])])},staticRenderFns:[]};var p=n("VU/8")({name:"add"},g,!1,function(t){n("Afxe")},"data-v-ff2a4b78",null).exports,C={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1663664322351",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2111",width:"32",height:"32"}},[e("path",{attrs:{d:"M658.069333 256a64 64 0 0 1 64 64l-0.021333 33.664 49.28-38.4A64 64 0 0 1 874.666667 365.781333v338.368a64 64 0 0 1-103.338667 50.474667l-49.28-38.4v26.496a64 64 0 0 1-64 64H213.333333a64 64 0 0 1-64-64V320a64 64 0 0 1 64-64h444.736z","p-id":"2112"}}),e("path",{attrs:{d:"M376.106667 638.933333l132.224-88.981333a21.333333 21.333333 0 0 0-0.170667-35.498667l-132.202667-87.274666a21.333333 21.333333 0 0 0-33.088 17.792v176.277333a21.333333 21.333333 0 0 0 33.237334 17.706667z","p-id":"2113"}})])])},staticRenderFns:[]};var $=n("VU/8")({name:"send-video"},C,!1,function(t){n("CTDB")},"data-v-6ae3d9c0",null).exports,M=n("mvHQ"),_=n.n(M),w={offerToReceiveVideo:!0,offerToReceiveAudio:!0},y=function(){var t,e,n,i=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{key:"tsm:"+Math.random().toString().substr(2,16),isSelf:!1,handleRemoteStream:null,handleIcecandidate:null},a={RTC:(t=function(t){"function"==typeof i.handleRemoteStream&&i.handleRemoteStream(t,a)},e=function(t){"function"==typeof i.handleIcecandidate&&i.handleIcecandidate(t,a)},n=new RTCPeerConnection({iceServers:[{urls:"stun:stun.l.google.com:19302"},{urls:"stun:stun1.l.google.com:19302"}]}),n.addEventListener("icecandidate",function(t){console.log("icecandidate event:",t),t.candidate?e({type:"candidate",label:t.candidate.sdpMLineIndex,id:t.candidate.sdpMid,candidate:t.candidate.candidate}):console.log("End of candidates.")}),n.addEventListener("track",function(e){"function"==typeof t&&t(e)}),n.setLocalDesc=function(t,e){return n.setLocalDescription(t).then(function(){e(t)}).catch(function(t){return console.log("setLocalDesc error:",t)})},n),key:i.key,isSelf:i.isSelf||!1,CreateOffer:function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:w;console.log("this",a),a.RTC.createOffer(e).then(function(e){a.RTC.setLocalDesc(e,t).then(function(){console.log("createOffer setLocalDesc OK")})}).catch(function(t){return console.log("createOffer error:",t)})},CreateAnswer:function(t){a.RTC.createAnswer().then(function(e){a.RTC.setLocalDesc(e,t).then(function(){console.log("CreateAnswer setLocalDesc OK")})}).catch(function(t){return console.log("createAnswer error:",t)})},HandleOffer:function(t,e){var n=new RTCSessionDescription(t);a.RTC.setRemoteDescription(n).then(function(){a.CreateAnswer(e)}).catch(function(t){return console.log("HandleOffer setRemoteDescription error:",t)})},HandleAnswer:function(t,e){var n=new RTCSessionDescription(t);a.RTC.setRemoteDescription(n).then(function(){e()}).catch(function(t){return console.log("HandleAnswer setRemoteDescription error:",t)})}};return a},D={name:"video-call",props:{roomData:{type:Object,required:!0}},data:function(){return{flag:!1,peerMap:{},memberKeys:[],localStream:null,localVideo:null,remoteVideo:null,localMediaConstraints:{video:!0,audio:!0}}},mounted:function(){console.log("open chat video call"),this.init()},computed:{videoCallRtcCandidate:function(){return this.$store.state.msg.VideoCallRtcCandidate},videoCallMsg:function(){return this.$store.state.msg.VideoCall}},watch:{videoCallRtcCandidate:function(){this.videoCallRtcCandidate&&this.videoCallRtcCandidate.status===r.c.candidate&&this.handleReceivedCandidate()},videoCallMsg:function(){if(this.videoCallMsg)switch(Number(this.videoCallMsg.status)){case r.c.agree:this.$Log("对方同意视频通话",this.videoCallMsg),this.createVideoCall();break;case r.c.createConn:this.$Log("createConn",this.videoCallMsg),this.flag||this.videoCallOfferHandle();break;case r.c.createConnConfirm:this.$Log("收到对方应答",this.videoCallMsg),this.videoCallAnswerHandle()}}},methods:{init:function(){this.localVideo=this.$refs.localVideoRef,this.memberKeys.push(this.roomData.object)},open:function(){var t=this;return!this.videoCallMsg&&(this.getLocalMediaData(function(){t.flag=!0;var e=t.$Chat.videoCall();t.$WebSocket.send(e)}),!0)},close:function(){this.flag=!1},handleRemoteStream:function(t,e){console.log("handleRemoteStream",t,e);var n=this.$refs["remote"+e.key+"VideoRef"][0];this.memberKeys.includes(e.key)&&n&&(n.srcObject=t.streams[0],n.addEventListener("loadedmetadata",function(){n.play()}))},handleReceivedCandidate:function(){if(this.videoCallRtcCandidate&&this.videoCallRtcCandidate.status===r.c.candidate){var t=JSON.parse(this.videoCallRtcCandidate.content),e=new RTCIceCandidate({sdpMLineIndex:t.label,candidate:t.candidate});this.peerMap[this.roomData.object].RTC.addIceCandidate(e).catch(function(t){console.log("addIceCandidate-error",t)})}},handleIcecandidate:function(t,e){var n=this.$Chat.videoCall(r.c.candidate);n.content=_()(t),this.send(n)},createPeerConn:function(t,e){var n=this;if(!this.peerMap[t]){var i=y({key:t,isSelf:e,handleRemoteStream:this.handleRemoteStream,handleIcecandidate:this.handleIcecandidate});e&&(this.localStream.getTracks().map(function(t){i.RTC.addTrack(t,n.localStream)}),i.CreateOffer(function(t){n.$Log("createVideoCall -> createOffer -> description:",t);var e=n.$Chat.videoCall(r.c.createConn);e.content=_()(t),n.send(e)})),this.peerMap[t]=i}return this.peerMap[t]},createVideoCall:function(){var t=this.createPeerConn(this.roomData.object,!0);this.peerMap[this.roomData.object]=t},videoCallAnswerHandle:function(){var t=this.videoCallMsg.sender,e=JSON.parse(this.videoCallMsg.content);this.peerMap[t].HandleAnswer(e,function(){})},getLocalMediaData:function(){var t=this,e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:null;navigator.mediaDevices.getUserMedia(this.localMediaConstraints).then(function(n){t.localStream=n,t.localVideo.srcObject=n,t.localVideo.play(),"function"==typeof e&&e(n)}).catch(function(e){t.$Log("本地视频采集错误「"+e.toString()+"」"),t.$toast.fail("视频通话暂不可用")})},videoCallOfferHandle:function(){var t=this;this.getLocalMediaData(function(e){if(t.$Log("signHandle videoCallMsg:",t.videoCallMsg),t.videoCallMsg&&t.videoCallMsg.status===r.c.createConn){var n=JSON.parse(t.videoCallMsg.content);t.flag=!0;var i=t.createPeerConn(t.roomData.object,!1);t.peerMap[t.roomData.object]=i,t.localStream.getTracks().map(function(e){i.RTC.addTrack(e,t.localStream)}),t.peerMap[t.roomData.object].HandleOffer(n,function(e){t.$Log("videoCallOfferHandle -> createAnswer -> description:",e);var n=t.$Chat.videoCall(r.c.createConnConfirm);n.content=_()(e),t.send(n)})}})},send:function(t){t.parent=this.videoCallMsg.parent||this.videoCallMsg.uuid,t.recipient=this.videoCallMsg.recipient,this.$WebSocket.send(t)}}},S={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{directives:[{name:"show",rawName:"v-show",value:this.flag,expression:"flag"}],staticClass:"video-call-box"},[e("div",{on:{click:this.close}},[this._v("关闭")]),this._v(" "),e("video",{ref:"localVideoRef",staticStyle:{width:"400px"},attrs:{autoplay:"",playsinline:"",controls:"",id:"localVideo"}}),this._v(" "),this._l(this.memberKeys,function(t){return[e("video",{key:t,ref:"remote"+t+"VideoRef",refInFor:!0,staticStyle:{width:"400px"},attrs:{autoplay:"",playsinline:"",controls:"",id:"remoteVideo"}})]})],2)},staticRenderFns:[]};var x={name:"send-message",components:{VideoCall:n("VU/8")(D,S,!1,function(t){n("UaVL")},"data-v-d3c344ac",null).exports,SendVideo:$,Add:p,Emoji:v,Voice:f},props:{roomData:{type:Object,required:!0}},data:function(){return{message:"",extraFlag:!1}},mounted:function(){this.$SetR(this.roomData.room)},methods:{showExtra:function(){this.extraFlag=!0},sendTextMessage:function(){this.message=this.message.trim(),""!==this.message&&(this.$WebSocket.send(this.$Chat.text(this.message)),this.$emit("sendMessage",this.message),this.message="")},sendVideoCall:function(){this.$refs.videoCallRef.open()}}},b={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"send-message"},[n("div",{staticClass:"top"},[n("div",{staticClass:"left"},[n("voice")],1),t._v(" "),n("div",{staticClass:"center"},[n("van-field",{staticStyle:{"border-radius":"8px"},attrs:{rows:"1",autosize:{maxHeight:100},type:"textarea"},on:{blur:t.sendTextMessage},model:{value:t.message,callback:function(e){t.message=e},expression:"message"}})],1),t._v(" "),n("div",{staticClass:"right"},[n("emoji"),t._v(" "),n("span",{on:{click:t.showExtra}},[n("add")],1)],1)]),t._v(" "),t.extraFlag?n("div",{staticClass:"extra"},[n("div",{staticClass:"extra-item",on:{click:t.sendVideoCall}},[n("send-video"),t._v(" "),n("span",{staticClass:"title"},[t._v("视频通话")])],1)]):t._e(),t._v(" "),n("video-call",{ref:"videoCallRef",attrs:{"room-data":t.roomData}})],1)},staticRenderFns:[]};var R={name:"chat-box",components:{SendMessage:n("VU/8")(x,b,!1,function(t){n("AIkw")},"data-v-67cc46db",null).exports,MessageList:u},props:{roomData:{type:Object}},mounted:function(){this.$refs.msgListRef.scrollToBottom()},methods:{sendMsg:function(){this.$refs.msgListRef.scrollToBottom()}}},T={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"chat-box"},[e("message-list",{ref:"msgListRef",attrs:{roomData:this.roomData}}),this._v(" "),e("send-message",{attrs:{roomData:this.roomData},on:{sendMessage:this.sendMsg}})],1)},staticRenderFns:[]};var O=n("VU/8")(R,T,!1,function(t){n("Ktc2")},"data-v-3d0fba8a",null);e.a=O.exports},AIkw:function(t,e){},AOzQ:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var i=n("9GPK"),a=n("ouP4"),s={name:"chat",components:{ChatBox:n("7paM").a,ChatHeader:a.a},data:function(){return{currentRoom:{}}},created:function(){var t=Object(i.c)(i.a);t?this.currentRoom=JSON.parse(t):this.$toast.fail("系统错误")},methods:{goBack:function(){this.$router.push({path:"/index"})}}},r={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"h100"},[e("chat-header",{attrs:{"room-data":this.currentRoom,"back-func":this.goBack}}),this._v(" "),e("chat-box",{attrs:{"room-data":this.currentRoom}})],1)},staticRenderFns:[]};var o=n("VU/8")(s,r,!1,function(t){n("FnSl")},"data-v-d4eb87cc",null);e.default=o.exports},Afxe:function(t,e){},CTDB:function(t,e){},CczY:function(t,e){},FnSl:function(t,e){},Ktc2:function(t,e){},UaVL:function(t,e){},XPRr:function(t,e){},n5u8:function(t,e){},oqQY:function(t,e,n){var i;i=function(){"use strict";var t="millisecond",e="second",n="minute",i="hour",a="day",s="week",r="month",o="quarter",c="year",d="date",l="Invalid Date",u=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,h=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,f={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_")},m=function(t,e,n){var i=String(t);return!i||i.length>=e?t:""+Array(e+1-i.length).join(n)+t},v={s:m,z:function(t){var e=-t.utcOffset(),n=Math.abs(e),i=Math.floor(n/60),a=n%60;return(e<=0?"+":"-")+m(i,2,"0")+":"+m(a,2,"0")},m:function t(e,n){if(e.date()<n.date())return-t(n,e);var i=12*(n.year()-e.year())+(n.month()-e.month()),a=e.clone().add(i,r),s=n-a<0,o=e.clone().add(i+(s?-1:1),r);return+(-(i+(n-a)/(s?a-o:o-a))||0)},a:function(t){return t<0?Math.ceil(t)||0:Math.floor(t)},p:function(l){return{M:r,y:c,w:s,d:a,D:d,h:i,m:n,s:e,ms:t,Q:o}[l]||String(l||"").toLowerCase().replace(/s$/,"")},u:function(t){return void 0===t}},g="en",p={};p[g]=f;var C=function(t){return t instanceof w},$=function t(e,n,i){var a;if(!e)return g;if("string"==typeof e){var s=e.toLowerCase();p[s]&&(a=s),n&&(p[s]=n,a=s);var r=e.split("-");if(!a&&r.length>1)return t(r[0])}else{var o=e.name;p[o]=e,a=o}return!i&&a&&(g=a),a||!i&&g},M=function(t,e){if(C(t))return t.clone();var n="object"==typeof e?e:{};return n.date=t,n.args=arguments,new w(n)},_=v;_.l=$,_.i=C,_.w=function(t,e){return M(t,{locale:e.$L,utc:e.$u,x:e.$x,$offset:e.$offset})};var w=function(){function f(t){this.$L=$(t.locale,null,!0),this.parse(t)}var m=f.prototype;return m.parse=function(t){this.$d=function(t){var e=t.date,n=t.utc;if(null===e)return new Date(NaN);if(_.u(e))return new Date;if(e instanceof Date)return new Date(e);if("string"==typeof e&&!/Z$/i.test(e)){var i=e.match(u);if(i){var a=i[2]-1||0,s=(i[7]||"0").substring(0,3);return n?new Date(Date.UTC(i[1],a,i[3]||1,i[4]||0,i[5]||0,i[6]||0,s)):new Date(i[1],a,i[3]||1,i[4]||0,i[5]||0,i[6]||0,s)}}return new Date(e)}(t),this.$x=t.x||{},this.init()},m.init=function(){var t=this.$d;this.$y=t.getFullYear(),this.$M=t.getMonth(),this.$D=t.getDate(),this.$W=t.getDay(),this.$H=t.getHours(),this.$m=t.getMinutes(),this.$s=t.getSeconds(),this.$ms=t.getMilliseconds()},m.$utils=function(){return _},m.isValid=function(){return!(this.$d.toString()===l)},m.isSame=function(t,e){var n=M(t);return this.startOf(e)<=n&&n<=this.endOf(e)},m.isAfter=function(t,e){return M(t)<this.startOf(e)},m.isBefore=function(t,e){return this.endOf(e)<M(t)},m.$g=function(t,e,n){return _.u(t)?this[e]:this.set(n,t)},m.unix=function(){return Math.floor(this.valueOf()/1e3)},m.valueOf=function(){return this.$d.getTime()},m.startOf=function(t,o){var l=this,u=!!_.u(o)||o,h=_.p(t),f=function(t,e){var n=_.w(l.$u?Date.UTC(l.$y,e,t):new Date(l.$y,e,t),l);return u?n:n.endOf(a)},m=function(t,e){return _.w(l.toDate()[t].apply(l.toDate("s"),(u?[0,0,0,0]:[23,59,59,999]).slice(e)),l)},v=this.$W,g=this.$M,p=this.$D,C="set"+(this.$u?"UTC":"");switch(h){case c:return u?f(1,0):f(31,11);case r:return u?f(1,g):f(0,g+1);case s:var $=this.$locale().weekStart||0,M=(v<$?v+7:v)-$;return f(u?p-M:p+(6-M),g);case a:case d:return m(C+"Hours",0);case i:return m(C+"Minutes",1);case n:return m(C+"Seconds",2);case e:return m(C+"Milliseconds",3);default:return this.clone()}},m.endOf=function(t){return this.startOf(t,!1)},m.$set=function(s,o){var l,u=_.p(s),h="set"+(this.$u?"UTC":""),f=(l={},l[a]=h+"Date",l[d]=h+"Date",l[r]=h+"Month",l[c]=h+"FullYear",l[i]=h+"Hours",l[n]=h+"Minutes",l[e]=h+"Seconds",l[t]=h+"Milliseconds",l)[u],m=u===a?this.$D+(o-this.$W):o;if(u===r||u===c){var v=this.clone().set(d,1);v.$d[f](m),v.init(),this.$d=v.set(d,Math.min(this.$D,v.daysInMonth())).$d}else f&&this.$d[f](m);return this.init(),this},m.set=function(t,e){return this.clone().$set(t,e)},m.get=function(t){return this[_.p(t)]()},m.add=function(t,o){var d,l=this;t=Number(t);var u=_.p(o),h=function(e){var n=M(l);return _.w(n.date(n.date()+Math.round(e*t)),l)};if(u===r)return this.set(r,this.$M+t);if(u===c)return this.set(c,this.$y+t);if(u===a)return h(1);if(u===s)return h(7);var f=(d={},d[n]=6e4,d[i]=36e5,d[e]=1e3,d)[u]||1,m=this.$d.getTime()+t*f;return _.w(m,this)},m.subtract=function(t,e){return this.add(-1*t,e)},m.format=function(t){var e=this,n=this.$locale();if(!this.isValid())return n.invalidDate||l;var i=t||"YYYY-MM-DDTHH:mm:ssZ",a=_.z(this),s=this.$H,r=this.$m,o=this.$M,c=n.weekdays,d=n.months,u=function(t,n,a,s){return t&&(t[n]||t(e,i))||a[n].slice(0,s)},f=function(t){return _.s(s%12||12,t,"0")},m=n.meridiem||function(t,e,n){var i=t<12?"AM":"PM";return n?i.toLowerCase():i},v={YY:String(this.$y).slice(-2),YYYY:this.$y,M:o+1,MM:_.s(o+1,2,"0"),MMM:u(n.monthsShort,o,d,3),MMMM:u(d,o),D:this.$D,DD:_.s(this.$D,2,"0"),d:String(this.$W),dd:u(n.weekdaysMin,this.$W,c,2),ddd:u(n.weekdaysShort,this.$W,c,3),dddd:c[this.$W],H:String(s),HH:_.s(s,2,"0"),h:f(1),hh:f(2),a:m(s,r,!0),A:m(s,r,!1),m:String(r),mm:_.s(r,2,"0"),s:String(this.$s),ss:_.s(this.$s,2,"0"),SSS:_.s(this.$ms,3,"0"),Z:a};return i.replace(h,function(t,e){return e||v[t]||a.replace(":","")})},m.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},m.diff=function(t,d,l){var u,h=_.p(d),f=M(t),m=6e4*(f.utcOffset()-this.utcOffset()),v=this-f,g=_.m(this,f);return g=(u={},u[c]=g/12,u[r]=g,u[o]=g/3,u[s]=(v-m)/6048e5,u[a]=(v-m)/864e5,u[i]=v/36e5,u[n]=v/6e4,u[e]=v/1e3,u)[h]||v,l?g:_.a(g)},m.daysInMonth=function(){return this.endOf(r).$D},m.$locale=function(){return p[this.$L]},m.locale=function(t,e){if(!t)return this.$L;var n=this.clone(),i=$(t,e,!0);return i&&(n.$L=i),n},m.clone=function(){return _.w(this.$d,this)},m.toDate=function(){return new Date(this.valueOf())},m.toJSON=function(){return this.isValid()?this.toISOString():null},m.toISOString=function(){return this.$d.toISOString()},m.toString=function(){return this.$d.toUTCString()},f}(),y=w.prototype;return M.prototype=y,[["$ms",t],["$s",e],["$m",n],["$H",i],["$W",a],["$M",r],["$y",c],["$D",d]].forEach(function(t){y[t[1]]=function(e){return this.$g(e,t[0],t[1])}}),M.extend=function(t,e){return t.$i||(t(e,w,M),t.$i=!0),M},M.locale=$,M.isDayjs=C,M.unix=function(t){return M(1e3*t)},M.en=p[g],M.Ls=p,M.p={},M},t.exports=i()}});
//# sourceMappingURL=2.3456505c3e949ae2bbb9.js.map