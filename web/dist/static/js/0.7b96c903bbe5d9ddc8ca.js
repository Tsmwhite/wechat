webpackJsonp([0],{"T+Xt":function(e,r){},"l+E2":function(module,__webpack_exports__,__webpack_require__){"use strict";__webpack_require__.d(__webpack_exports__,"b",function(){return isEmail}),__webpack_require__.d(__webpack_exports__,"a",function(){return copyObject});var __WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify__=__webpack_require__("mvHQ"),__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify___default=__webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify__),Tools={},getUrlParam=function(e){var r=window.location.search.substring(1);if(""==r){var t=window.location.href;r=t.substring(t.indexOf("?")+1)}for(var a=r.split("&"),n=0;n<a.length;n++){var i=a[n].split("=");if(i[0]==e)return i[1]}return!1},replaceParamVal=function replaceParamVal(paramName,replaceWith){var oUrl=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";oUrl=oUrl||window.location.href.toString();var re=eval("/("+paramName+"=)([^&]*)/gi");return oUrl.replace(re,paramName+"="+replaceWith)},isWxBrowser=function(){return"micromessenger"==window.navigator.userAgent.toLowerCase().match(/MicroMessenger/i)},getUrlQuery=function(){return window.location.href.substring(window.location.href.indexOf("?"))},queryStringToObject=function(e){e.indexOf("?")>-1&&(e=e.substring(e.indexOf("?")+1));for(var r={},t=e.split("&"),a=0;a<t.length;a++){var n=t[a].split("=");r[n[0]]=n[1]||null}return r},isEmail=function(e){return!!/^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/.test(e)},copyObject=function(e){return JSON.parse(__WEBPACK_IMPORTED_MODULE_0_babel_runtime_core_js_json_stringify___default()(e))},_unused_webpack_default_export={getUrlParam:getUrlParam,replaceParamVal:replaceParamVal,isWxBrowser:isWxBrowser,getUrlQuery:getUrlQuery,isEmail:isEmail,copyObject:copyObject}},ouP4:function(e,r,t){"use strict";var a={name:"chat-header",props:{backFunc:{type:Function},roomData:{type:Object},showMoreBtn:{type:Boolean,default:!0}},data:function(){return{}},methods:{back:function(){"function"==typeof this.backFunc?this.backFunc():this.$router.go(-1)}}},n={render:function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("div",{staticClass:"chat-header"},[t("div",{staticClass:"left"},[t("van-icon",{staticClass:"icon",attrs:{name:"arrow-left"},on:{click:e.back}})],1),e._v(" "),t("div",{staticClass:"center"},[1===e.roomData.type?t("div",{staticClass:"title"},[e._v(e._s(e.roomData.name)+"（"+e._s(e.roomData.member_num)+"）")]):t("div",{staticClass:"title"},[e._v(e._s(e.roomData.name))])]),e._v(" "),t("div",{staticClass:"right"},[e.showMoreBtn?t("van-icon",{staticClass:"icon",attrs:{name:"ellipsis"}}):e._e()],1)])},staticRenderFns:[]};var i=t("VU/8")(a,n,!1,function(e){t("T+Xt")},"data-v-648546a1",null);r.a=i.exports}});
//# sourceMappingURL=0.7b96c903bbe5d9ddc8ca.js.map