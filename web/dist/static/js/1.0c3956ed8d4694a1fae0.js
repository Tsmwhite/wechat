webpackJsonp([1],{"/i1X":function(t,e){},"2jWG":function(t,e){},"7paM":function(t,e,a){"use strict";var n=a("9GPK"),i=a("y0/c"),s=a("t5DY"),o=a("++XJ"),r={name:"message-list",props:{roomData:{type:Object}},data:function(){return{MessageTypes:o.b,loading:!1,loadingController:!1,finished:!1,refreshing:!1,loadLastId:0,pagination:{current:1,pageSize:60,total:0}}},computed:{meAvatar:function(){return Object(n.e)().avatar||"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png"},messages:function(){var t=this.$store.state,e=this.roomData.room||0,a=t.msg.MessageMapList[e];return a||[]},getAvatar:function(){var t=this;return function(e){var a=t.$store.state.friend.FriendMap[e.sender];return a?a.avatar||"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png":(Object(i.a)({user_id:e.sender}),"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png")}}},mounted:function(){},methods:{init:function(){this.finished=!1;var t=this.$store.state,e=this.roomData.room||0,a=t.msg.MessageMapList[e];(!a||a.length<1)&&e>0?this.loadHistory():this.finished=!0},currentUuid:function(){return Object(n.b)()},loadHistory:function(){var t=this;this.roomData.room||(this.finished=!0),Object(s.e)({room_uuid:this.roomData.room,size:this.pagination.pageSize,last_id:this.loadLastId}).then(function(e){var a=e.data.length;if(a<1)return t.finished=!0,void(t.loading=!1);var n=document.getElementById("messageList"),i=n.scrollHeight;t.loadLastId=e.data[a-1].id||0,t.$store.dispatch("history",{recipient:t.roomData.room,messages:e.data.reverse()}),t.$nextTick(function(){1===t.pagination.current?t.scrollToBottom():n.scrollTop=n.scrollHeight-i,t.pagination.current++,setTimeout(function(){t.loading=!1},3e3)})}).catch(function(){t.loading=!1})},scrollToBottom:function(){this.$nextTick(function(){var t=document.getElementById("messageList");t.scrollTop=t.scrollHeight})}}},c={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"message-list",attrs:{id:"messageList"}},[a("van-list",{attrs:{id:"messageListScroll",finished:t.finished,"finished-text":"没有更多了",offset:80,direction:"up"},on:{load:t.loadHistory},model:{value:t.loading,callback:function(e){t.loading=e},expression:"loading"}},[t._l(t.messages,function(e,n){return[e.second_type===t.MessageTypes.videoCall?a("div",{key:n},[a("div",{staticStyle:{"font-size":"12px",color:"rgba(0,0,0,0.65)"}},[t._v(" "+t._s(e.sender===t.currentUuid()?"我":"对方")+"发起了视频通话【已结束】")])]):a("div",{key:n,class:["message-box",{right:e.sender===t.currentUuid()}]},[e.sender!==t.currentUuid()?a("div",{staticClass:"avatar"},[a("img",{attrs:{src:t.getAvatar(e)}})]):t._e(),t._v(" "),e.sender===t.currentUuid()?a("div",{staticClass:"before-tip"},[t._e(),t._v(" "),t._e()],1):t._e(),t._v(" "),e.second_type===t.MessageTypes.text?[a("div",{staticClass:"text",domProps:{innerHTML:t._s(e.content)}})]:e.second_type===t.MessageTypes.image?[a("div",{staticClass:"image"},[a("img",{attrs:{src:e.content}})])]:t._e(),t._v(" "),e.sender===t.currentUuid()?a("div",{staticClass:"avatar"},[a("img",{attrs:{src:t.meAvatar}})]):t._e()],2)]})],2)],1)},staticRenderFns:[]};var l=a("VU/8")(r,c,!1,function(t){a("/i1X")},"data-v-324daf56",null).exports,d={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1647771122497",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"16033",width:"32",height:"32"}},[e("path",{attrs:{d:"M410.208 319.776a96 96 0 0 1 192 0v192.448a96 96 0 0 1-96 95.744 96 96 0 0 1-96-95.744V319.776z m96 352.192c88.192 0 160-71.648 160-159.744V319.776a160.096 160.096 0 0 0-160-159.776c-88.224 0-160 71.68-160 159.776v192.448a160.032 160.032 0 0 0 160 159.744zM742.4 448a32 32 0 0 0-32 32v24.672c0 116.192-94.72 210.688-211.2 210.688-116.448 0-211.2-94.496-211.2-210.688V480a32 32 0 0 0-64 0v24.672c0 140.16 105.76 255.904 241.76 272.448V864a32 32 0 0 0 64 0v-86.432C667.168 762.304 774.4 645.824 774.4 504.672V480a32 32 0 0 0-32-32","p-id":"16034"}})])])},staticRenderFns:[]};var u=a("VU/8")({name:"voice"},d,!1,function(t){a("XPRr")},"data-v-29829874",null).exports,f={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1647770199888",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"4299",width:"32",height:"32"}},[e("path",{attrs:{d:"M510.944 960c-247.04 0-448-200.96-448-448s200.992-448 448-448 448 200.96 448 448-200.96 448-448 448z m0-832c-211.744 0-384 172.256-384 384s172.256 384 384 384 384-172.256 384-384-172.256-384-384-384z",fill:"","p-id":"4300"}}),e("path",{attrs:{d:"M512 773.344c-89.184 0-171.904-40.32-226.912-110.624-10.88-13.92-8.448-34.016 5.472-44.896 13.888-10.912 34.016-8.48 44.928 5.472 42.784 54.688 107.136 86.048 176.512 86.048 70.112 0 134.88-31.904 177.664-87.552 10.784-14.016 30.848-16.672 44.864-5.888 14.016 10.784 16.672 30.88 5.888 44.864C685.408 732.32 602.144 773.344 512 773.344zM368 515.2c-26.528 0-48-21.472-48-48v-64c0-26.528 21.472-48 48-48s48 21.472 48 48v64c0 26.496-21.504 48-48 48zM656 515.2c-26.496 0-48-21.472-48-48v-64c0-26.528 21.504-48 48-48s48 21.472 48 48v64c0 26.496-21.504 48-48 48z",fill:"","p-id":"4301"}})])])},staticRenderFns:[]};var v=a("VU/8")({name:"emoji"},f,!1,function(t){a("CczY")},"data-v-7ff93387",null).exports,h={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1647770134907",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2981",width:"32",height:"32"}},[e("path",{attrs:{d:"M512 958.016c-119.648 0-232.128-46.368-316.736-130.56C110.624 743.2 64 631.2 64 512c0-119.168 46.624-231.2 131.232-315.424 84.608-84.192 197.088-130.56 316.736-130.56s232.128 46.368 316.704 130.56c84.672 84.224 131.264 196.256 131.264 315.392 0.032 119.2-46.592 231.232-131.264 315.456C744.128 911.616 631.648 958.016 512 958.016zM512 129.984c-102.624 0-199.072 39.744-271.584 111.936C167.936 314.048 128 409.984 128 512c0 102.016 39.904 197.952 112.384 270.048 72.512 72.192 168.96 111.936 271.584 111.936 102.592 0 199.072-39.744 271.584-111.936 72.48-72.16 112.416-168.064 112.384-270.08 0-102.016-39.904-197.92-112.384-270.016C711.072 169.76 614.592 129.984 512 129.984z","p-id":"2982"}}),e("path",{attrs:{d:"M736 480l-192 0L544 288c0-17.664-14.336-32-32-32s-32 14.336-32 32l0 192L288 480c-17.664 0-32 14.336-32 32s14.336 32 32 32l192 0 0 192c0 17.696 14.336 32 32 32s32-14.304 32-32l0-192 192 0c17.696 0 32-14.336 32-32S753.696 480 736 480z","p-id":"2983"}})])])},staticRenderFns:[]};var m=a("VU/8")({name:"add"},h,!1,function(t){a("Afxe")},"data-v-ff2a4b78",null).exports,g={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("svg",{staticClass:"icon",attrs:{t:"1663664322351",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2111",width:"32",height:"32"}},[e("path",{attrs:{d:"M658.069333 256a64 64 0 0 1 64 64l-0.021333 33.664 49.28-38.4A64 64 0 0 1 874.666667 365.781333v338.368a64 64 0 0 1-103.338667 50.474667l-49.28-38.4v26.496a64 64 0 0 1-64 64H213.333333a64 64 0 0 1-64-64V320a64 64 0 0 1 64-64h444.736z","p-id":"2112"}}),e("path",{attrs:{d:"M376.106667 638.933333l132.224-88.981333a21.333333 21.333333 0 0 0-0.170667-35.498667l-132.202667-87.274666a21.333333 21.333333 0 0 0-33.088 17.792v176.277333a21.333333 21.333333 0 0 0 33.237334 17.706667z","p-id":"2113"}})])])},staticRenderFns:[]};var p=a("VU/8")({name:"send-video"},g,!1,function(t){a("CTDB")},"data-v-6ae3d9c0",null).exports,A=a("mvHQ"),C=a.n(A),w={offerToReceiveVideo:!0,offerToReceiveAudio:!0},k=function(){var t,e,a,n=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{key:"tsm:"+Math.random().toString().substr(2,16),isSelf:!1,handleRemoteStream:null,handleIcecandidate:null},i={RTC:(t=function(t){"function"==typeof n.handleRemoteStream&&n.handleRemoteStream(t,i)},e=function(t){"function"==typeof n.handleIcecandidate&&n.handleIcecandidate(t,i)},a=new RTCPeerConnection({iceServers:[{urls:"stun:stun.l.google.com:19302"},{urls:"stun:stun1.l.google.com:19302"}]}),a.addEventListener("icecandidate",function(t){console.log("icecandidate event:",t),t.candidate?e({type:"candidate",label:t.candidate.sdpMLineIndex,id:t.candidate.sdpMid,candidate:t.candidate.candidate}):console.log("End of candidates.")}),a.addEventListener("track",function(e){"function"==typeof t&&t(e)}),a.setLocalDesc=function(t,e){return a.setLocalDescription(t).then(function(){e(t)}).catch(function(t){return console.log("setLocalDesc error:",t)})},a),key:n.key,isSelf:n.isSelf||!1,CreateOffer:function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:w;console.log("this",i),i.RTC.createOffer(e).then(function(e){i.RTC.setLocalDesc(e,t).then(function(){console.log("createOffer setLocalDesc OK")})}).catch(function(t){return console.log("createOffer error:",t)})},CreateAnswer:function(t){i.RTC.createAnswer().then(function(e){i.RTC.setLocalDesc(e,t).then(function(){console.log("CreateAnswer setLocalDesc OK")})}).catch(function(t){return console.log("createAnswer error:",t)})},HandleOffer:function(t,e){var a=new RTCSessionDescription(t);i.RTC.setRemoteDescription(a).then(function(){i.CreateAnswer(e)}).catch(function(t){return console.log("HandleOffer setRemoteDescription error:",t)})},HandleAnswer:function(t,e){var a=new RTCSessionDescription(t);i.RTC.setRemoteDescription(a).then(function(){e()}).catch(function(t){return console.log("HandleAnswer setRemoteDescription error:",t)})}};return i},b={name:"video-call",props:{roomData:{type:Object,required:!0}},data:function(){return{flag:!1,peerMap:{},memberKeys:[],localStream:null,localVideo:null,remoteVideo:null,localMediaConstraints:{video:!0,audio:!0}}},mounted:function(){console.log("open chat video call"),this.init()},computed:{videoCallRtcCandidate:function(){return this.$store.state.msg.VideoCallRtcCandidate},videoCallMsg:function(){return this.$store.state.msg.VideoCall}},watch:{videoCallRtcCandidate:function(){this.videoCallRtcCandidate&&this.videoCallRtcCandidate.status===o.c.candidate&&this.handleReceivedCandidate()},videoCallMsg:function(){if(this.videoCallMsg)switch(Number(this.videoCallMsg.status)){case o.c.agree:this.$Log("对方同意视频通话",this.videoCallMsg),this.createVideoCall();break;case o.c.createConn:this.$Log("createConn",this.videoCallMsg),this.flag||this.videoCallOfferHandle();break;case o.c.createConnConfirm:this.$Log("收到对方应答",this.videoCallMsg),this.videoCallAnswerHandle()}}},methods:{init:function(){this.localVideo=this.$refs.localVideoRef,this.memberKeys.push(this.roomData.object)},open:function(){var t=this;return!this.videoCallMsg&&(this.getLocalMediaData(function(){t.flag=!0;var e=t.$Chat.videoCall();t.$WebSocket.send(e)}),!0)},close:function(){this.flag=!1},handleRemoteStream:function(t,e){console.log("handleRemoteStream",t,e);var a=this.$refs["remote"+e.key+"VideoRef"][0];this.memberKeys.includes(e.key)&&a&&(a.srcObject=t.streams[0],a.addEventListener("loadedmetadata",function(){a.play()}))},handleReceivedCandidate:function(){if(this.videoCallRtcCandidate&&this.videoCallRtcCandidate.status===o.c.candidate){var t=JSON.parse(this.videoCallRtcCandidate.content),e=new RTCIceCandidate({sdpMLineIndex:t.label,candidate:t.candidate});this.peerMap[this.roomData.object].RTC.addIceCandidate(e).catch(function(t){console.log("addIceCandidate-error",t)})}},handleIcecandidate:function(t,e){var a=this.$Chat.videoCall(o.c.candidate);a.content=C()(t),this.send(a)},createPeerConn:function(t,e){var a=this;if(!this.peerMap[t]){var n=k({key:t,isSelf:e,handleRemoteStream:this.handleRemoteStream,handleIcecandidate:this.handleIcecandidate});e&&(this.localStream.getTracks().map(function(t){n.RTC.addTrack(t,a.localStream)}),n.CreateOffer(function(t){a.$Log("createVideoCall -> createOffer -> description:",t);var e=a.$Chat.videoCall(o.c.createConn);e.content=C()(t),a.send(e)})),this.peerMap[t]=n}return this.peerMap[t]},createVideoCall:function(){var t=this.createPeerConn(this.roomData.object,!0);this.peerMap[this.roomData.object]=t},videoCallAnswerHandle:function(){var t=this.videoCallMsg.sender,e=JSON.parse(this.videoCallMsg.content);this.peerMap[t].HandleAnswer(e,function(){})},getLocalMediaData:function(){var t=this,e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:null;navigator.mediaDevices.getUserMedia(this.localMediaConstraints).then(function(a){t.localStream=a,t.localVideo.srcObject=a,t.localVideo.play(),"function"==typeof e&&e(a)}).catch(function(e){t.$Log("本地视频采集错误「"+e.toString()+"」"),t.$toast.fail("视频通话暂不可用")})},videoCallOfferHandle:function(){var t=this;this.getLocalMediaData(function(e){if(t.$Log("signHandle videoCallMsg:",t.videoCallMsg),t.videoCallMsg&&t.videoCallMsg.status===o.c.createConn){var a=JSON.parse(t.videoCallMsg.content);t.flag=!0;var n=t.createPeerConn(t.roomData.object,!1);t.peerMap[t.roomData.object]=n,t.localStream.getTracks().map(function(e){n.RTC.addTrack(e,t.localStream)}),t.peerMap[t.roomData.object].HandleOffer(a,function(e){t.$Log("videoCallOfferHandle -> createAnswer -> description:",e);var a=t.$Chat.videoCall(o.c.createConnConfirm);a.content=C()(e),t.send(a)})}})},send:function(t){t.parent=this.videoCallMsg.parent||this.videoCallMsg.uuid,t.recipient=this.videoCallMsg.recipient,this.$WebSocket.send(t)}}},S={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{directives:[{name:"show",rawName:"v-show",value:this.flag,expression:"flag"}],staticClass:"video-call-box"},[e("div",{on:{click:this.close}},[this._v("关闭")]),this._v(" "),e("video",{ref:"localVideoRef",staticStyle:{width:"400px"},attrs:{autoplay:"",playsinline:"",controls:"",id:"localVideo"}}),this._v(" "),this._l(this.memberKeys,function(t){return[e("video",{key:t,ref:"remote"+t+"VideoRef",refInFor:!0,staticStyle:{width:"400px"},attrs:{autoplay:"",playsinline:"",controls:"",id:"remoteVideo"}})]})],2)},staticRenderFns:[]};var _={name:"send-message",components:{VideoCall:a("VU/8")(b,S,!1,function(t){a("UaVL")},"data-v-d3c344ac",null).exports,SendVideo:p,Add:m,Emoji:v,Voice:u},props:{roomData:{type:Object,required:!0}},data:function(){return{message:"",extraFlag:!1}},mounted:function(){this.$SetR(this.roomData.room)},methods:{showExtra:function(){this.extraFlag=!0},sendTextMessage:function(){this.message=this.message.trim(),""!==this.message&&(this.$WebSocket.send(this.$Chat.text(this.message)),this.$emit("sendMessage",this.message),this.message="")},sendVideoCall:function(){this.$refs.videoCallRef.open()}}},x={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"send-message"},[a("div",{staticClass:"top"},[a("div",{staticClass:"left"},[a("voice")],1),t._v(" "),a("div",{staticClass:"center"},[a("van-field",{staticStyle:{"border-radius":"8px"},attrs:{rows:"1",autosize:{maxHeight:100},type:"textarea"},on:{blur:t.sendTextMessage},model:{value:t.message,callback:function(e){t.message=e},expression:"message"}})],1),t._v(" "),a("div",{staticClass:"right"},[a("emoji"),t._v(" "),a("span",{on:{click:t.showExtra}},[a("add")],1)],1)]),t._v(" "),t.extraFlag?a("div",{staticClass:"extra"},[a("div",{staticClass:"extra-item",on:{click:t.sendVideoCall}},[a("send-video"),t._v(" "),a("span",{staticClass:"title"},[t._v("视频通话")])],1)]):t._e(),t._v(" "),a("video-call",{ref:"videoCallRef",attrs:{"room-data":t.roomData}})],1)},staticRenderFns:[]};var y={name:"chat-box",components:{SendMessage:a("VU/8")(_,x,!1,function(t){a("AIkw")},"data-v-67cc46db",null).exports,MessageList:l},props:{roomData:{type:Object}},mounted:function(){this.$refs.msgListRef.scrollToBottom()},methods:{sendMsg:function(){this.$refs.msgListRef.scrollToBottom()}}},I={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"chat-box"},[e("message-list",{ref:"msgListRef",attrs:{roomData:this.roomData}}),this._v(" "),e("send-message",{attrs:{roomData:this.roomData},on:{sendMessage:this.sendMsg}})],1)},staticRenderFns:[]};var M=a("VU/8")(y,I,!1,function(t){a("Ktc2")},"data-v-3d0fba8a",null);e.a=M.exports},"8hXn":function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a("mvHQ"),i=a.n(n),s=a("t5DY"),o=a("7paM"),r=a("ouP4"),c={name:"contact-item",props:{contact:{type:Object,default:function(){return{name:"小马",avatar:"https://wwcdn.weixin.qq.com/node/wework/images/kf_head_image_url_4.png",unread_count:0,only_tip:!1,last_msg:"",last_time:""}}}},methods:{}},l={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"contact-option"},[n("div",{staticClass:"left"},[n("div",{class:["avatar",{unread:t.contact.unread_count>0},{"only-tip":t.contact.only_tip}],attrs:{"data-unread":t.contact.unread_count>99?99:t.contact.unread_count}},[t.contact.avatar?n("img",{attrs:{src:t.contact.avatar}}):n("img",{attrs:{src:a("pCOc")}})])]),t._v(" "),n("div",{staticClass:"right"},[n("div",{staticClass:"info"},[n("div",{staticClass:"nickname"},[t._v(t._s(t.contact.name||t.contact.title))]),t._v(" "),n("div",{staticClass:"last-msg"},[t._v(t._s(t.contact.last_msg))])]),t._v(" "),n("div",{staticClass:"extra"},[t._t("extra",[n("div",{staticClass:"last-time"},[t._v(t._s(t.contact.last_time))])])],2)])])},staticRenderFns:[]};var d=a("VU/8")(c,l,!1,function(t){a("2jWG")},"data-v-91512db6",null).exports,u=a("Dd8w"),f=a.n(u),v=a("l+E2"),h={name:"add-friend-modal",props:{userInfo:{type:Object,default:function(){return{}}}},data:function(){return{flag:!1,formState:{type:0,uuid:"",content:"",remark:""},formStateDefault:{}}},mounted:function(){this.formStateDefault=Object(v.a)(this.formState)},methods:{show:function(){this.reset(),this.flag=!0},reset:function(){this.formState=Object(v.a)(this.formStateDefault)},onSubmit:function(){var t=this;if(this.formState.content=this.formState.content.trim(),this.formState.remark=this.formState.remark.trim(),""===this.formState.content)return this.$toast.fail("请输入验证信息");Object(s.k)(f()({},this.formState,{uuid:this.userInfo.uuid})).then(function(e){t.$toast.success(e.msg),t.flag=!1})}}},m={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("van-popup",{staticStyle:{width:"80%",padding:"16px","border-radius":"8px"},model:{value:t.flag,callback:function(e){t.flag=e},expression:"flag"}},[a("h3",{staticStyle:{"margin-top":"0"}},[t._v("添加好友")]),t._v(" "),a("van-form",{on:{submit:t.onSubmit}},[a("van-field",{attrs:{rows:"3",autosize:"",label:"验证信息",type:"textarea",maxlength:"120",placeholder:"请输入验证信息（必填）","show-word-limit":""},model:{value:t.formState.content,callback:function(e){t.$set(t.formState,"content",e)},expression:"formState.content"}}),t._v(" "),a("van-field",{attrs:{label:"备注信息",placeholder:"通过好友申请后自动备注",maxlength:"12","show-word-limit":""},model:{value:t.formState.remark,callback:function(e){t.$set(t.formState,"remark",e)},expression:"formState.remark"}}),t._v(" "),a("div",{staticStyle:{margin:"16px"}},[a("van-button",{attrs:{round:"",block:"",type:"info","native-type":"submit"}},[t._v("提交")])],1)],1)],1)},staticRenderFns:[]};var g={name:"search-box",components:{AddFriendModal:a("VU/8")(h,m,!1,function(t){a("DZls")},"data-v-b8e725fa",null).exports,ContactItem:d},data:function(){return{showSearchFlag:!1,active:0,keyword:"",friends:[],users:[],rooms:[],existSearchMap:{0:"",1:"",2:""},addUserInfo:{}}},watch:{keyword:function(){this.keyword=this.keyword.trim()}},computed:{list:function(){switch(this.active){case 0:return this.friends;case 1:return this.users;case 2:return this.rooms}return[]}},methods:{tabChange:function(){""!==this.keyword&&this.existSearchMap[this.active]!==this.keyword&&this.search()},search:function(){if(""!==this.keyword)switch(this.active){case 0:this.searchFriends();break;case 1:this.searchUser();break;case 2:this.searchRoom()}},searchFriends:function(){var t=this;Object(s.h)({keyword:this.keyword}).then(function(e){t.friends=e.data||[],t.existSearchMap[0]=t.keyword})},searchUser:function(){var t=this;Object(s.j)({keyword:this.keyword}).then(function(e){t.users=e.data||[],t.existSearchMap[1]=t.keyword})},searchRoom:function(){var t=this;Object(s.i)({keyword:this.keyword}).then(function(e){t.rooms=e.data||[],t.existSearchMap[2]=t.keyword})},addFriend:function(t){this.addUserInfo=t,this.$refs.addFriendRef.show()},addGroup:function(){},linkChat:function(){}}},p={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"search-box"},[a("van-search",{attrs:{"show-action":t.showSearchFlag,placeholder:"搜索"},on:{focus:function(e){t.showSearchFlag=!0},cancel:function(e){t.showSearchFlag=!1}}}),t._v(" "),a("van-popup",{staticStyle:{width:"100%",height:"60%"},attrs:{position:"top"},model:{value:t.showSearchFlag,callback:function(e){t.showSearchFlag=e},expression:"showSearchFlag"}},[a("van-search",{attrs:{placeholder:"搜索"},on:{search:t.search},model:{value:t.keyword,callback:function(e){t.keyword=e},expression:"keyword"}}),t._v(" "),a("div",{staticClass:"content"},[a("van-tabbar",{attrs:{fixed:!1},on:{change:t.tabChange},model:{value:t.active,callback:function(e){t.active=e},expression:"active"}},[a("van-tabbar-item",{attrs:{icon:"friends-o"}},[t._v("好友")]),t._v(" "),a("van-tabbar-item",{attrs:{icon:"user-o"}},[t._v("用户")]),t._v(" "),a("van-tabbar-item",{attrs:{icon:"chat-o"}},[t._v("群聊")])],1),t._v(" "),t.list.length>0?a("div",{staticClass:"search-list"},[t._l(t.list,function(e,n){return[a("contact-item",{key:n,attrs:{contact:e},scopedSlots:t._u([{key:"extra",fn:function(){return[1===t.active?a("van-button",{attrs:{type:"info",size:"small"},on:{click:function(a){return t.addFriend(e)}}},[t._v("加好友")]):2===t.active?a("van-button",{attrs:{type:"info",size:"small"},on:{click:function(a){return t.addGroup(e)}}},[t._v("加群")]):a("van-button",{attrs:{type:"info",size:"small"},on:{click:t.linkChat}},[t._v("发消息")])]},proxy:!0}],null,!0)})]})],2):a("van-empty",{attrs:{image:"search",description:"暂无数据"}})],1)],1),t._v(" "),a("add-friend-modal",{ref:"addFriendRef",attrs:{"user-info":t.addUserInfo}})],1)},staticRenderFns:[]};var A=a("VU/8")(g,p,!1,function(t){a("StIe")},"data-v-30d8e6cc",null).exports,C=a("pCOc"),w=a.n(C),k=a("9GPK"),b={name:"index",components:{SearchBox:A,ChatHeader:r.a,ChatBox:o.a},data:function(){return{groupDefaultAvatar:w.a,value:"",loading:!0,finished:!1,data:[],chat:{show:!1,currentRoom:{}}}},mounted:function(){this.init(),this.$WebSocket.Detect()},methods:{init:function(){this.loadData()},loadData:function(){var t=this;s.a.getContacts().then(function(e){t.data=e.data,t.loading=!1,t.finished=!0})},openChat:function(t){console.log("chatRoom",t),Object(k.i)(k.a,i()(t)),this.$router.push({path:"/chat"})},openFriendNoticeBox:function(){this.$router.push({path:"/friend-notice"})}}},S={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"h100"},[n("div",{staticClass:"contact-list"},[n("search-box"),t._v(" "),n("van-list",{attrs:{finished:t.finished},on:{load:t.loadData},model:{value:t.loading,callback:function(e){t.loading=e},expression:"loading"}},[n("div",{staticClass:"contact-option",on:{click:t.openFriendNoticeBox}},[n("div",{staticClass:"left"},[n("div",{staticClass:"avatar"},[n("img",{attrs:{src:a("IO8I")}})])]),t._v(" "),n("div",{staticClass:"right"},[n("div",{staticClass:"info"},[n("div",{staticClass:"nickname"},[t._v("好友通知")]),t._v(" "),n("div",{staticClass:"last-msg"})]),t._v(" "),n("div",{staticClass:"extra"},[n("div",{staticClass:"last-time"})])])]),t._v(" "),t._l(t.data,function(e,a){return n("div",{key:a,staticClass:"contact-option",on:{click:function(a){return t.openChat(e)}}},[n("div",{staticClass:"left"},[n("div",{class:["avatar",{unread:e.unreadCount>0},{"only-tip":e.onlyTip}],attrs:{"data-unread":e.unreadCount>99?99:e.unreadCount}},[n("img",{directives:[{name:"lazy",rawName:"v-lazy",value:e.avatar||t.groupDefaultAvatar,expression:"item.avatar || groupDefaultAvatar"}]})])]),t._v(" "),n("div",{staticClass:"right"},[n("div",{staticClass:"info"},[n("div",{staticClass:"nickname"},[t._v(t._s(e.name))]),t._v(" "),n("div",{staticClass:"last-msg"},[t._v(t._s(e.lastMsg))])]),t._v(" "),n("div",{staticClass:"extra"},[n("div",{staticClass:"last-time"},[t._v(t._s(e.lastTime))])])])])})],2)],1)])},staticRenderFns:[]};var _=a("VU/8")(b,S,!1,function(t){a("fNp4")},"data-v-4f3be4fc",null);e.default=_.exports},AIkw:function(t,e){},Afxe:function(t,e){},CTDB:function(t,e){},CczY:function(t,e){},DZls:function(t,e){},IO8I:function(t,e){t.exports="data:image/png;base64,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"},Ktc2:function(t,e){},StIe:function(t,e){},UaVL:function(t,e){},XPRr:function(t,e){},fNp4:function(t,e){},"l+E2":function(module,__webpack_exports__,__webpack_require__){"use strict";__webpack_require__.d(__webpack_exports__,"b",function(){return isEmail}),__webpack_require__.d(__webpack_exports__,"a",function(){return copyObject});var __WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify__=__webpack_require__("mvHQ"),__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify___default=__webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify__),Tools={},getUrlParam=function(t){var e=window.location.search.substring(1);if(""==e){var a=window.location.href;e=a.substring(a.indexOf("?")+1)}for(var n=e.split("&"),i=0;i<n.length;i++){var s=n[i].split("=");if(s[0]==t)return s[1]}return!1},replaceParamVal=function replaceParamVal(paramName,replaceWith){var oUrl=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";oUrl=oUrl||window.location.href.toString();var re=eval("/("+paramName+"=)([^&]*)/gi");return oUrl.replace(re,paramName+"="+replaceWith)},isWxBrowser=function(){return"micromessenger"==window.navigator.userAgent.toLowerCase().match(/MicroMessenger/i)},getUrlQuery=function(){return window.location.href.substring(window.location.href.indexOf("?"))},queryStringToObject=function(t){t.indexOf("?")>-1&&(t=t.substring(t.indexOf("?")+1));for(var e={},a=t.split("&"),n=0;n<a.length;n++){var i=a[n].split("=");e[i[0]]=i[1]||null}return e},isEmail=function(t){return!!/^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/.test(t)},copyObject=function(t){return JSON.parse(__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify___default()(t))},_unused_webpack_default_export={getUrlParam:getUrlParam,replaceParamVal:replaceParamVal,isWxBrowser:isWxBrowser,getUrlQuery:getUrlQuery,isEmail:isEmail,copyObject:copyObject}},pCOc:function(t,e,a){t.exports=a.p+"static/img/default-group-avatar.ab4db50.jpeg"}});
//# sourceMappingURL=1.0c3956ed8d4694a1fae0.js.map