webpackJsonp([3],{QPNa:function(e,t,i){e.exports=i.p+"static/img/login-bg.d5d498a.png"},Quw4:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var o=i("QPNa"),a=i.n(o),n=i("l+E2"),s=i("l/JR"),r=i("9GPK"),l=i("wy+4"),c={name:"login",data:function(){return{Bg:a.a,checked:!0,form:{phone:"",email:"",verifyCode:""},sendCodeWait:0}},watch:{"form.phone":function(){this.form.phone=this.form.phone.trim()},"form.email":function(){this.form.email=this.form.email.trim()},"form.verifyCode":function(){this.form.verifyCode=this.form.verifyCode.trim()}},methods:{sendCode:function(){var e,t=this;return""===this.form.email?this.$toast.fail("请输入邮箱"):Object(n.b)(this.form.email)?(this.sendCodeWait=60,void(e={account:this.form.email,type:1},s.a.post("/sendCode",e)).then(function(){var e=setInterval(function(){t.sendCodeWait-=1,t.sendCodeWait<1&&clearInterval(e)},1e3);return t.$toast.success("发送完成")}).catch(function(){t.sendCodeWait=0})):this.$toast.fail("邮箱格式不正确")},onSubmit:function(e){var t,i=this;return""===this.form.email?this.$toast.fail("请输入邮箱"):Object(n.b)(this.form.email)?""===this.form.verifyCode?this.$toast.fail("请输入验证码"):void(t={account:this.form.email,verify_code:this.form.verifyCode},s.a.post("/loginReg",t)).then(function(e){Object(r.g)(e.data),Object(l.a)(e.data),i.$WebSocket.Init(),i.$router.push({path:"/index"})}):this.$toast.fail("邮箱格式不正确")}}},f={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"loginApp"},[i("img",{staticClass:"bgImage",attrs:{src:e.Bg}}),e._v(" "),i("div",{staticClass:"form-box"},[i("h3",[e._v("尬聊啊！兄嘚！")]),e._v(" "),i("van-form",[i("van-field",{attrs:{label:"邮箱",clearable:!0,clickable:!0,placeholder:"请输入邮箱",maxlength:50},model:{value:e.form.email,callback:function(t){e.$set(e.form,"email",t)},expression:"form.email"}}),e._v(" "),i("van-field",{attrs:{type:"number",clearable:!0,clickable:!0,maxlength:6,center:"",clearable:"",label:"验证码",placeholder:"请输入验证码"},scopedSlots:e._u([{key:"button",fn:function(){return[e.sendCodeWait>0?i("van-button",{attrs:{size:"small",type:"info",disabled:""}},[e._v(e._s(e.sendCodeWait)+" 秒后重新发送")]):i("van-button",{attrs:{size:"small",type:"info"},on:{click:e.sendCode}},[e._v("发送验证码")])]},proxy:!0}]),model:{value:e.form.verifyCode,callback:function(t){e.$set(e.form,"verifyCode",t)},expression:"form.verifyCode"}}),e._v(" "),i("van-button",{staticStyle:{width:"80%","margin-top":"20px"},attrs:{round:"",type:"info"},on:{click:e.onSubmit}},[e._v("登录\n            ")]),e._v(" "),i("van-checkbox",{staticStyle:{"margin-top":"20px","font-size":"12px",display:"flex","justify-content":"center"},attrs:{"icon-size":"14px",shape:"square"},model:{value:e.checked,callback:function(t){e.checked=t},expression:"checked"}},[e._v("已阅读并同意使用协议\n            ")])],1)],1)])},staticRenderFns:[]};var d=i("VU/8")(c,f,!1,function(e){i("ujDh")},"data-v-1888fb9b",null);t.default=d.exports},ujDh:function(e,t){}});
//# sourceMappingURL=3.0d8297661d618e116761.js.map