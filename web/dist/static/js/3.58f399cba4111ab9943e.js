webpackJsonp([3],{QPNa:function(e,t,r){e.exports=r.p+"static/img/login-bg.d5d498a.png"},Quw4:function(e,t,r){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var a=r("QPNa"),i=r.n(a),n=r("l+E2"),o=r("l/JR"),s=r("9GPK"),l=r("wy+4"),c={name:"login",data:function(){return{Bg:i.a,checked:!0,form:{phone:"",email:"",verifyCode:""},sendCodeWait:0}},watch:{"form.phone":function(){this.form.phone=this.form.phone.trim()},"form.email":function(){this.form.email=this.form.email.trim()},"form.verifyCode":function(){this.form.verifyCode=this.form.verifyCode.trim()}},methods:{sendCode:function(){var e,t=this;return""===this.form.email?this.$toast.fail("请输入邮箱"):Object(n.b)(this.form.email)?(this.sendCodeWait=60,void(e={account:this.form.email,type:1},o.a.post("/sendCode",e)).then(function(){var e=setInterval(function(){t.sendCodeWait-=1,t.sendCodeWait<1&&clearInterval(e)},1e3);return t.$toast.success("发送完成")}).catch(function(){t.sendCodeWait=0})):this.$toast.fail("邮箱格式不正确")},onSubmit:function(e){var t,r=this;return""===this.form.email?this.$toast.fail("请输入邮箱"):Object(n.b)(this.form.email)?""===this.form.verifyCode?this.$toast.fail("请输入验证码"):void(t={account:this.form.email,verify_code:this.form.verifyCode},o.a.post("/loginReg",t)).then(function(e){Object(s.g)(e.data),Object(l.a)(e.data),r.$WebSocket.Init(),r.$router.push({path:"/index"})}):this.$toast.fail("邮箱格式不正确")}}},_={render:function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"loginApp"},[r("img",{staticClass:"bgImage",attrs:{src:e.Bg}}),e._v(" "),r("div",{staticClass:"form-box"},[r("h3",[e._v("尬聊啊！兄嘚！")]),e._v(" "),r("van-form",[r("van-field",{attrs:{label:"邮箱",clearable:!0,clickable:!0,placeholder:"请输入邮箱",maxlength:50},model:{value:e.form.email,callback:function(t){e.$set(e.form,"email",t)},expression:"form.email"}}),e._v(" "),r("van-field",{attrs:{type:"number",clearable:!0,clickable:!0,maxlength:6,center:"",clearable:"",label:"验证码",placeholder:"请输入验证码"},scopedSlots:e._u([{key:"button",fn:function(){return[e.sendCodeWait>0?r("van-button",{attrs:{size:"small",type:"info",disabled:""}},[e._v(e._s(e.sendCodeWait)+" 秒后重新发送")]):r("van-button",{attrs:{size:"small",type:"info"},on:{click:e.sendCode}},[e._v("发送验证码")])]},proxy:!0}]),model:{value:e.form.verifyCode,callback:function(t){e.$set(e.form,"verifyCode",t)},expression:"form.verifyCode"}}),e._v(" "),r("van-button",{staticStyle:{width:"80%","margin-top":"20px"},attrs:{round:"",type:"info"},on:{click:e.onSubmit}},[e._v("登录\n            ")]),e._v(" "),r("van-checkbox",{staticStyle:{"margin-top":"20px","font-size":"12px",display:"flex","justify-content":"center"},attrs:{"icon-size":"14px",shape:"square"},model:{value:e.checked,callback:function(t){e.checked=t},expression:"checked"}},[e._v("已阅读并同意使用协议\n            ")])],1)],1)])},staticRenderFns:[]};var f=r("VU/8")(c,_,!1,function(e){r("ujDh")},"data-v-1888fb9b",null);t.default=f.exports},"l+E2":function(module,__webpack_exports__,__webpack_require__){"use strict";__webpack_require__.d(__webpack_exports__,"b",function(){return isEmail}),__webpack_require__.d(__webpack_exports__,"a",function(){return copyObject});var __WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify__=__webpack_require__("mvHQ"),__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify___default=__webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify__),Tools={},getUrlParam=function(e){var t=window.location.search.substring(1);if(""==t){var r=window.location.href;t=r.substring(r.indexOf("?")+1)}for(var a=t.split("&"),i=0;i<a.length;i++){var n=a[i].split("=");if(n[0]==e)return n[1]}return!1},replaceParamVal=function replaceParamVal(paramName,replaceWith){var oUrl=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";oUrl=oUrl||window.location.href.toString();var re=eval("/("+paramName+"=)([^&]*)/gi");return oUrl.replace(re,paramName+"="+replaceWith)},isWxBrowser=function(){return"micromessenger"==window.navigator.userAgent.toLowerCase().match(/MicroMessenger/i)},getUrlQuery=function(){return window.location.href.substring(window.location.href.indexOf("?"))},queryStringToObject=function(e){e.indexOf("?")>-1&&(e=e.substring(e.indexOf("?")+1));for(var t={},r=e.split("&"),a=0;a<r.length;a++){var i=r[a].split("=");t[i[0]]=i[1]||null}return t},isEmail=function(e){return!!/^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/.test(e)},copyObject=function(e){return JSON.parse(__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify___default()(e))},_unused_webpack_default_export={getUrlParam:getUrlParam,replaceParamVal:replaceParamVal,isWxBrowser:isWxBrowser,getUrlQuery:getUrlQuery,isEmail:isEmail,copyObject:copyObject}},ujDh:function(e,t){}});
//# sourceMappingURL=3.58f399cba4111ab9943e.js.map