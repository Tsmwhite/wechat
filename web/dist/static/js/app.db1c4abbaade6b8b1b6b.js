webpackJsonp([6],{"++XJ":function(e,t,n){"use strict";n.d(t,"a",function(){return r}),n.d(t,"b",function(){return o}),n.d(t,"c",function(){return i});var r={ping:0,chat:200},o={text:400,image:401,videoCall:406},i={wait:0,cancel:1,agree:2,reject:3,timeouts:4,candidate:5,createConn:10,createConnConfirm:11}},0:function(e,t){},"4B/U":function(e,t,n){"use strict";function r(e){if(null==e)return window;if("[object Window]"!==e.toString()){var t=e.ownerDocument;return t&&t.defaultView||window}return e}function o(e){return e instanceof r(e).Element||e instanceof Element}function i(e){return e instanceof r(e).HTMLElement||e instanceof HTMLElement}function a(e){return"undefined"!=typeof ShadowRoot&&(e instanceof r(e).ShadowRoot||e instanceof ShadowRoot)}n.d(t,"a",function(){return $}),n.d(t,"b",function(){return U});var s=Math.round;function c(e,t){void 0===t&&(t=!1);var n=e.getBoundingClientRect(),r=1,o=1;if(i(e)&&t){var a=e.offsetHeight,c=e.offsetWidth;c>0&&(r=s(n.width)/c||1),a>0&&(o=s(n.height)/a||1)}return{width:n.width/r,height:n.height/o,top:n.top/o,right:n.right/r,bottom:n.bottom/o,left:n.left/r,x:n.left/r,y:n.top/o}}function u(e){var t=r(e);return{scrollLeft:t.pageXOffset,scrollTop:t.pageYOffset}}function f(e){return e?(e.nodeName||"").toLowerCase():null}function l(e){return((o(e)?e.ownerDocument:e.document)||window.document).documentElement}function p(e){return r(e).getComputedStyle(e)}function d(e){var t=p(e),n=t.overflow,r=t.overflowX,o=t.overflowY;return/auto|scroll|overlay|hidden/.test(n+o+r)}function h(e,t,n){void 0===n&&(n=!1);var o,a,p=i(t),h=i(t)&&function(e){var t=e.getBoundingClientRect(),n=s(t.width)/e.offsetWidth||1,r=s(t.height)/e.offsetHeight||1;return 1!==n||1!==r}(t),m=l(t),v=c(e,h),g={scrollLeft:0,scrollTop:0},b={x:0,y:0};return(p||!p&&!n)&&(("body"!==f(t)||d(m))&&(g=(o=t)!==r(o)&&i(o)?{scrollLeft:(a=o).scrollLeft,scrollTop:a.scrollTop}:u(o)),i(t)?((b=c(t,!0)).x+=t.clientLeft,b.y+=t.clientTop):m&&(b.x=function(e){return c(l(e)).left+u(e).scrollLeft}(m))),{x:v.left+g.scrollLeft-b.x,y:v.top+g.scrollTop-b.y,width:v.width,height:v.height}}function m(e){return"html"===f(e)?e:e.assignedSlot||e.parentNode||(a(e)?e.host:null)||l(e)}function v(e,t){var n;void 0===t&&(t=[]);var o=function e(t){return["html","body","#document"].indexOf(f(t))>=0?t.ownerDocument.body:i(t)&&d(t)?t:e(m(t))}(e),a=o===(null==(n=e.ownerDocument)?void 0:n.body),s=r(o),c=a?[s].concat(s.visualViewport||[],d(o)?o:[]):o,u=t.concat(c);return a?u:u.concat(v(m(c)))}function g(e){return["table","td","th"].indexOf(f(e))>=0}function b(e){return i(e)&&"fixed"!==p(e).position?e.offsetParent:null}function y(e){for(var t=r(e),n=b(e);n&&g(n)&&"static"===p(n).position;)n=b(n);return n&&("html"===f(n)||"body"===f(n)&&"static"===p(n).position)?t:n||function(e){var t=-1!==navigator.userAgent.toLowerCase().indexOf("firefox");if(-1!==navigator.userAgent.indexOf("Trident")&&i(e)&&"fixed"===p(e).position)return null;var n=m(e);for(a(n)&&(n=n.host);i(n)&&["html","body"].indexOf(f(n))<0;){var r=p(n);if("none"!==r.transform||"none"!==r.perspective||"paint"===r.contain||-1!==["transform","perspective"].indexOf(r.willChange)||t&&"filter"===r.willChange||t&&r.filter&&"none"!==r.filter)return n;n=n.parentNode}return null}(e)||t}var w="top",C="bottom",O="right",_="left",S="auto",x="start",E="end",k=[].concat([w,C,O,_],[S]).reduce(function(e,t){return e.concat([t,t+"-"+x,t+"-"+E])},[]),M=["beforeRead","read","afterRead","beforeMain","main","afterMain","beforeWrite","write","afterWrite"];function j(e){var t=new Map,n=new Set,r=[];return e.forEach(function(e){t.set(e.name,e)}),e.forEach(function(e){n.has(e.name)||function e(o){n.add(o.name),[].concat(o.requires||[],o.requiresIfExists||[]).forEach(function(r){if(!n.has(r)){var o=t.get(r);o&&e(o)}}),r.push(o)}(e)}),r}function L(e){for(var t=arguments.length,n=new Array(t>1?t-1:0),r=1;r<t;r++)n[r-1]=arguments[r];return[].concat(n).reduce(function(e,t){return e.replace(/%s/,t)},e)}var R='Popper: modifier "%s" provided an invalid %s property, expected %s but got %s',A='Popper: modifier "%s" requires "%s", but "%s" modifier is not available',T=["name","enabled","phase","fn","effect","requires","options"];function I(e){return e.split("-")[0]}function D(e){return e.split("-")[1]}var P="Popper: Invalid reference or popper argument provided. They must be either a DOM element or virtual element.",F="Popper: An infinite loop in the modifiers cycle has been detected! The cycle has been interrupted to prevent a browser crash.",N={placement:"bottom",modifiers:[],strategy:"absolute"};function W(){for(var e=arguments.length,t=new Array(e),n=0;n<e;n++)t[n]=arguments[n];return!t.some(function(e){return!(e&&"function"==typeof e.getBoundingClientRect)})}var V={passive:!0};var H={top:"auto",right:"auto",bottom:"auto",left:"auto"};function q(e){var t,n=e.popper,o=e.popperRect,i=e.placement,a=e.variation,c=e.offsets,u=e.position,f=e.gpuAcceleration,d=e.adaptive,h=e.roundOffsets,m=e.isFixed,v=c.x,g=void 0===v?0:v,b=c.y,S=void 0===b?0:b,x="function"==typeof h?h({x:g,y:S}):{x:g,y:S};g=x.x,S=x.y;var k=c.hasOwnProperty("x"),M=c.hasOwnProperty("y"),j=_,L=w,R=window;if(d){var A=y(n),T="clientHeight",I="clientWidth";if(A===r(n)&&"static"!==p(A=l(n)).position&&"absolute"===u&&(T="scrollHeight",I="scrollWidth"),A=A,i===w||(i===_||i===O)&&a===E)L=C,S-=(m&&A===R&&R.visualViewport?R.visualViewport.height:A[T])-o.height,S*=f?1:-1;if(i===_||(i===w||i===C)&&a===E)j=O,g-=(m&&A===R&&R.visualViewport?R.visualViewport.width:A[I])-o.width,g*=f?1:-1}var D,P=Object.assign({position:u},d&&H),F=!0===h?function(e){var t=e.x,n=e.y,r=window.devicePixelRatio||1;return{x:s(t*r)/r||0,y:s(n*r)/r||0}}({x:g,y:S}):{x:g,y:S};return g=F.x,S=F.y,f?Object.assign({},P,((D={})[L]=M?"0":"",D[j]=k?"0":"",D.transform=(R.devicePixelRatio||1)<=1?"translate("+g+"px, "+S+"px)":"translate3d("+g+"px, "+S+"px, 0)",D)):Object.assign({},P,((t={})[L]=M?S+"px":"",t[j]=k?g+"px":"",t.transform="",t))}var $=function(e){void 0===e&&(e={});var t=e,n=t.defaultModifiers,r=void 0===n?[]:n,i=t.defaultOptions,a=void 0===i?N:i;return function(e,t,n){void 0===n&&(n=a);var i,s,u={placement:"bottom",orderedModifiers:[],options:Object.assign({},N,a),modifiersData:{},elements:{reference:e,popper:t},attributes:{},styles:{}},f=[],l=!1,d={state:u,setOptions:function(n){var i="function"==typeof n?n(u.options):n;m(),u.options=Object.assign({},a,u.options,i),u.scrollParents={reference:o(e)?v(e):e.contextElement?v(e.contextElement):[],popper:v(t)};var s=function(e){var t=j(e);return M.reduce(function(e,n){return e.concat(t.filter(function(e){return e.phase===n}))},[])}(function(e){var t=e.reduce(function(e,t){var n=e[t.name];return e[t.name]=n?Object.assign({},n,t,{options:Object.assign({},n.options,t.options),data:Object.assign({},n.data,t.data)}):t,e},{});return Object.keys(t).map(function(e){return t[e]})}([].concat(r,u.options.modifiers)));u.orderedModifiers=s.filter(function(e){return e.enabled}),function(e){e.forEach(function(t){[].concat(Object.keys(t),T).filter(function(e,t,n){return n.indexOf(e)===t}).forEach(function(n){switch(n){case"name":"string"!=typeof t.name&&console.error(L(R,String(t.name),'"name"','"string"','"'+String(t.name)+'"'));break;case"enabled":"boolean"!=typeof t.enabled&&console.error(L(R,t.name,'"enabled"','"boolean"','"'+String(t.enabled)+'"'));break;case"phase":M.indexOf(t.phase)<0&&console.error(L(R,t.name,'"phase"',"either "+M.join(", "),'"'+String(t.phase)+'"'));break;case"fn":"function"!=typeof t.fn&&console.error(L(R,t.name,'"fn"','"function"','"'+String(t.fn)+'"'));break;case"effect":null!=t.effect&&"function"!=typeof t.effect&&console.error(L(R,t.name,'"effect"','"function"','"'+String(t.fn)+'"'));break;case"requires":null==t.requires||Array.isArray(t.requires)||console.error(L(R,t.name,'"requires"','"array"','"'+String(t.requires)+'"'));break;case"requiresIfExists":Array.isArray(t.requiresIfExists)||console.error(L(R,t.name,'"requiresIfExists"','"array"','"'+String(t.requiresIfExists)+'"'));break;case"options":case"data":break;default:console.error('PopperJS: an invalid property has been provided to the "'+t.name+'" modifier, valid properties are '+T.map(function(e){return'"'+e+'"'}).join(", ")+'; but "'+n+'" was provided.')}t.requires&&t.requires.forEach(function(n){null==e.find(function(e){return e.name===n})&&console.error(L(A,String(t.name),n,n))})})})}((c=[].concat(s,u.options.modifiers),l=function(e){return e.name},h=new Set,c.filter(function(e){var t=l(e);if(!h.has(t))return h.add(t),!0}))),I(u.options.placement)===S&&(u.orderedModifiers.find(function(e){return"flip"===e.name})||console.error(['Popper: "auto" placements require the "flip" modifier be',"present and enabled to work."].join(" ")));var c,l,h,g=p(t);return[g.marginTop,g.marginRight,g.marginBottom,g.marginLeft].some(function(e){return parseFloat(e)})&&console.warn(['Popper: CSS "margin" styles cannot be used to apply padding',"between the popper and its reference element or boundary.","To replicate margin, use the `offset` modifier, as well as","the `padding` option in the `preventOverflow` and `flip`","modifiers."].join(" ")),u.orderedModifiers.forEach(function(e){var t=e.name,n=e.options,r=void 0===n?{}:n,o=e.effect;if("function"==typeof o){var i=o({state:u,name:t,instance:d,options:r});f.push(i||function(){})}}),d.update()},forceUpdate:function(){if(!l){var e=u.elements,t=e.reference,n=e.popper;if(W(t,n)){var r,o,i,a;u.rects={reference:h(t,y(n),"fixed"===u.options.strategy),popper:(r=n,o=c(r),i=r.offsetWidth,a=r.offsetHeight,Math.abs(o.width-i)<=1&&(i=o.width),Math.abs(o.height-a)<=1&&(a=o.height),{x:r.offsetLeft,y:r.offsetTop,width:i,height:a})},u.reset=!1,u.placement=u.options.placement,u.orderedModifiers.forEach(function(e){return u.modifiersData[e.name]=Object.assign({},e.data)});for(var s=0,f=0;f<u.orderedModifiers.length;f++){if((s+=1)>100){console.error(F);break}if(!0!==u.reset){var p=u.orderedModifiers[f],m=p.fn,v=p.options,g=void 0===v?{}:v,b=p.name;"function"==typeof m&&(u=m({state:u,options:g,name:b,instance:d})||u)}else u.reset=!1,f=-1}}else console.error(P)}},update:(i=function(){return new Promise(function(e){d.forceUpdate(),e(u)})},function(){return s||(s=new Promise(function(e){Promise.resolve().then(function(){s=void 0,e(i())})})),s}),destroy:function(){m(),l=!0}};if(!W(e,t))return console.error(P),d;function m(){f.forEach(function(e){return e()}),f=[]}return d.setOptions(n).then(function(e){!l&&n.onFirstUpdate&&n.onFirstUpdate(e)}),d}}({defaultModifiers:[{name:"eventListeners",enabled:!0,phase:"write",fn:function(){},effect:function(e){var t=e.state,n=e.instance,o=e.options,i=o.scroll,a=void 0===i||i,s=o.resize,c=void 0===s||s,u=r(t.elements.popper),f=[].concat(t.scrollParents.reference,t.scrollParents.popper);return a&&f.forEach(function(e){e.addEventListener("scroll",n.update,V)}),c&&u.addEventListener("resize",n.update,V),function(){a&&f.forEach(function(e){e.removeEventListener("scroll",n.update,V)}),c&&u.removeEventListener("resize",n.update,V)}},data:{}},{name:"popperOffsets",enabled:!0,phase:"read",fn:function(e){var t=e.state,n=e.name;t.modifiersData[n]=function(e){var t,n=e.reference,r=e.element,o=e.placement,i=o?I(o):null,a=o?D(o):null,s=n.x+n.width/2-r.width/2,c=n.y+n.height/2-r.height/2;switch(i){case w:t={x:s,y:n.y-r.height};break;case C:t={x:s,y:n.y+n.height};break;case O:t={x:n.x+n.width,y:c};break;case _:t={x:n.x-r.width,y:c};break;default:t={x:n.x,y:n.y}}var u=i?function(e){return["top","bottom"].indexOf(e)>=0?"x":"y"}(i):null;if(null!=u){var f="y"===u?"height":"width";switch(a){case x:t[u]=t[u]-(n[f]/2-r[f]/2);break;case E:t[u]=t[u]+(n[f]/2-r[f]/2)}}return t}({reference:t.rects.reference,element:t.rects.popper,strategy:"absolute",placement:t.placement})},data:{}},{name:"computeStyles",enabled:!0,phase:"beforeWrite",fn:function(e){var t=e.state,n=e.options,r=n.gpuAcceleration,o=void 0===r||r,i=n.adaptive,a=void 0===i||i,s=n.roundOffsets,c=void 0===s||s,u=p(t.elements.popper).transitionProperty||"";a&&["transform","top","right","bottom","left"].some(function(e){return u.indexOf(e)>=0})&&console.warn(["Popper: Detected CSS transitions on at least one of the following",'CSS properties: "transform", "top", "right", "bottom", "left".',"\n\n",'Disable the "computeStyles" modifier\'s `adaptive` option to allow',"for smooth transitions, or remove these properties from the CSS","transition declaration on the popper element if only transitioning","opacity or background-color for example.","\n\n","We recommend using the popper element as a wrapper around an inner","element that can have any CSS property transitioned for animations."].join(" "));var f={placement:I(t.placement),variation:D(t.placement),popper:t.elements.popper,popperRect:t.rects.popper,gpuAcceleration:o,isFixed:"fixed"===t.options.strategy};null!=t.modifiersData.popperOffsets&&(t.styles.popper=Object.assign({},t.styles.popper,q(Object.assign({},f,{offsets:t.modifiersData.popperOffsets,position:t.options.strategy,adaptive:a,roundOffsets:c})))),null!=t.modifiersData.arrow&&(t.styles.arrow=Object.assign({},t.styles.arrow,q(Object.assign({},f,{offsets:t.modifiersData.arrow,position:"absolute",adaptive:!1,roundOffsets:c})))),t.attributes.popper=Object.assign({},t.attributes.popper,{"data-popper-placement":t.placement})},data:{}},{name:"applyStyles",enabled:!0,phase:"write",fn:function(e){var t=e.state;Object.keys(t.elements).forEach(function(e){var n=t.styles[e]||{},r=t.attributes[e]||{},o=t.elements[e];i(o)&&f(o)&&(Object.assign(o.style,n),Object.keys(r).forEach(function(e){var t=r[e];!1===t?o.removeAttribute(e):o.setAttribute(e,!0===t?"":t)}))})},effect:function(e){var t=e.state,n={popper:{position:t.options.strategy,left:"0",top:"0",margin:"0"},arrow:{position:"absolute"},reference:{}};return Object.assign(t.elements.popper.style,n.popper),t.styles=n,t.elements.arrow&&Object.assign(t.elements.arrow.style,n.arrow),function(){Object.keys(t.elements).forEach(function(e){var r=t.elements[e],o=t.attributes[e]||{},a=Object.keys(t.styles.hasOwnProperty(e)?t.styles[e]:n[e]).reduce(function(e,t){return e[t]="",e},{});i(r)&&f(r)&&(Object.assign(r.style,a),Object.keys(o).forEach(function(e){r.removeAttribute(e)}))})}},requires:["computeStyles"]}]});var U={name:"offset",enabled:!0,phase:"main",requires:["popperOffsets"],fn:function(e){var t=e.state,n=e.options,r=e.name,o=n.offset,i=void 0===o?[0,0]:o,a=k.reduce(function(e,n){return e[n]=function(e,t,n){var r=I(e),o=[_,w].indexOf(r)>=0?-1:1,i="function"==typeof n?n(Object.assign({},t,{placement:e})):n,a=i[0],s=i[1];return a=a||0,s=(s||0)*o,[_,O].indexOf(r)>=0?{x:s,y:a}:{x:a,y:s}}(n,t.rects,i),e},{}),s=a[t.placement],c=s.x,u=s.y;null!=t.modifiersData.popperOffsets&&(t.modifiersData.popperOffsets.x+=c,t.modifiersData.popperOffsets.y+=u),t.modifiersData[r]=a}}},"4ml/":function(e,t){},"9GPK":function(e,t,n){"use strict";n.d(t,"i",function(){return c}),n.d(t,"c",function(){return u}),n.d(t,"a",function(){return l}),n.d(t,"d",function(){return h}),n.d(t,"b",function(){return m}),n.d(t,"e",function(){return v}),n.d(t,"h",function(){return g}),n.d(t,"g",function(){return b}),n.d(t,"f",function(){return y});var r=n("mvHQ"),o=n.n(r),i=n("BO1k"),a=n.n(i),s="Tsmwhite:",c=function(e,t){window.localStorage.setItem(s+e,t)},u=function(e){return window.localStorage.getItem(s+e)},f=function(e){return window.localStorage.removeItem(s+e)},l="current-contact",p=["token","user-info",l],d={},h=function(){return d.token||(d.token=u("token")),d.token},m=function(){if(!d.uuid){var e=u("user-info");if(!e)return 0;e=JSON.parse(e),d.uuid=e.uuid}return d.uuid},v=function(){if(!d.userInfo){var e=u("user-info");d.userInfo=JSON.parse(e)}return d.userInfo||{}},g=function(){var e=!0,t=!1,n=void 0;try{for(var r,o=a()(p);!(e=(r=o.next()).done);e=!0){var i=r.value;f(i)}}catch(e){t=!0,n=e}finally{try{!e&&o.return&&o.return()}finally{if(t)throw n}}},b=function(e){c("token",e.token),c("user-info",o()(e))},y=function(){var e=u("token");return null!==e&&""!==e}},NHnr:function(e,t,n){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var r=n("7+uW"),o={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"main-tab"},[t("div",[t("van-icon",{staticClass:"icon",attrs:{name:"chat"}})],1),this._v(" "),t("div",[t("van-icon",{staticClass:"icon",attrs:{name:"friends-o"}})],1),this._v(" "),t("div",[t("van-icon",{staticClass:"icon",attrs:{name:"fire-o"}})],1),this._v(" "),t("div",[t("van-icon",{staticClass:"icon",attrs:{name:"manager-o"}})],1)])},staticRenderFns:[]};var i=n("VU/8")({name:"main-tab"},o,!1,function(e){n("eeW5")},"data-v-0492bc3b",null).exports,a=n("mvHQ"),s=n.n(a),c=n("t5DY"),u=n("++XJ"),f=n("9GPK"),l={name:"video-call-apply",data:function(){return{applyNoticeFlag:!1,friendInfo:{}}},computed:{videoCallMsg:function(){return this.$store.state.msg.VideoCall}},watch:{videoCallMsg:function(){if(this.videoCallMsg)switch(Number(this.videoCallMsg.status)){case u.c.wait:this.show()}}},methods:{show:function(){var e=this;Object(c.c)({friend:this.videoCallMsg.sender}).then(function(t){e.friendInfo=t.data,e.applyNoticeFlag=!0})},close:function(){this.applyNoticeFlag=!1},agree:function(){this.$Log("agree",this.videoCallMsg);var e=this.$Chat.videoCall(u.c.agree);e.parent=this.videoCallMsg.uuid,e.recipient=this.videoCallMsg.recipient,this.$WebSocket.send(e),this.createConnect()},reject:function(){this.$Log("reject",this.videoCallMsg);var e=this.$Chat.videoCall(u.c.reject);e.parent=this.videoCallMsg.uuid,e.recipient=this.videoCallMsg.recipient,this.$WebSocket.send(e),this.close()},createConnect:function(){var e=this;Object(c.b)({room:this.videoCallMsg.recipient}).then(function(t){var n=t.data;e.$Log("chatRoom",n),e.close(),Object(f.i)(f.a,s()(n)),e.$router.replace({path:"/chat"})})}}},p={render:function(){var e=this,t=e.$createElement,n=e._self._c||t;return e.applyNoticeFlag?n("div",{staticClass:"apply-notice-container"},[n("div",{staticClass:"apply-notice-box"},[n("img",{staticClass:"avatar",attrs:{src:e.friendInfo.avatar}}),e._v(" "),n("div",{staticClass:"title"},[n("span",{staticStyle:{color:"#1989fa"}},[e._v(e._s(e.friendInfo.remark||e.friendInfo.name))]),e._v("\n            邀请你视频通话\n        ")]),e._v(" "),n("div",{staticStyle:{"margin-top":"12px"}},[n("van-button",{staticStyle:{width:"120px"},attrs:{type:"warning",size:"small"},on:{click:e.reject}},[e._v("拒绝\n            ")]),e._v(" "),n("van-button",{staticStyle:{width:"120px","margin-left":"24px"},attrs:{type:"info",size:"small"},on:{click:e.agree}},[e._v("同意\n            ")])],1)])]):e._e()},staticRenderFns:[]};var d={name:"global-notice",components:{VideoCallApply:n("VU/8")(l,p,!1,function(e){n("R+zl")},"data-v-fdb412f0",null).exports}},h={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",[t("video-call-apply")],1)},staticRenderFns:[]};var m={name:"App",components:{GlobalNotice:n("VU/8")(d,h,!1,function(e){n("nbfa")},"data-v-7e7f7130",null).exports,MainTab:i},computed:{showMainTab:function(){return this.$route.meta.showMainTab||!1}}},v={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{attrs:{id:"app"}},[t("router-view"),this._v(" "),this.showMainTab?t("main-tab"):this._e(),this._v(" "),t("global-notice")],1)},staticRenderFns:[]};var g=n("VU/8")(m,v,!1,function(e){n("jBzp")},null,null).exports,b=n("/ocq");r.a.use(b.a);var y=new b.a({routes:[{path:"/login",name:"login",component:function(){return n.e(3).then(n.bind(null,"Quw4"))}},{path:"/",name:"home",component:function(){return Promise.all([n.e(0),n.e(1)]).then(n.bind(null,"8hXn"))},meta:{auth:!0,showMainTab:!0}},{path:"/index",name:"index",component:function(){return Promise.all([n.e(0),n.e(1)]).then(n.bind(null,"8hXn"))},meta:{auth:!0,showMainTab:!0}},{path:"/chat",name:"chat",component:function(){return Promise.all([n.e(0),n.e(2)]).then(n.bind(null,"AOzQ"))},meta:{auth:!0}},{path:"/friend-notice",name:"friendNotice",component:function(){return Promise.all([n.e(0),n.e(4)]).then(n.bind(null,"JkKV"))},meta:{auth:!0}}]});y.beforeEach(function(e,t,n){var r=e.meta.auth||!1;console.log("hasAuth",r,Object(f.f)()),!r||Object(f.f)()?n():n("/login")});var w=y,C=n("Fd2+"),O=(n("4ml/"),n("Dd8w")),_=n.n(O),S={Recipient:""},x=function(){return{type:u.a.chat,recipient:S.Recipient,sender:Object(f.b)(),send_time:Math.ceil((new Date).getTime()/1e3)}},E={ping:function(){return{content:"ping",type:u.a.ping}},text:function(e){return _()({content:e,second_type:u.b.text},x())},image:function(e){return _()({content:e,second_type:u.b.image},x())},videoCall:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:u.c.wait,t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[];return _()({content:t.length?t.join(","):Object(f.b)(),second_type:u.b.videoCall,status:Number(e)},x())}},k=n("wy+4"),M=function(e){switch(e.second_type){case u.b.text:case u.b.image:k.c.dispatch("push",e);break;case u.b.videoCall:switch(e.status){case u.c.wait:case u.c.agree:case u.c.createConn:case u.c.createConnConfirm:e.sender!==Object(f.b)()&&k.c.dispatch("videoCallSet",e);break;case u.c.candidate:e.sender!==Object(f.b)()&&k.c.dispatch("videoCallRtcCandidateSet",e);break;case u.c.cancel:case u.c.reject:case u.c.timeouts:}}},j={Error:"",MessageHandleCallback:null,_token:"",_socket:null,_heartTimeout:8e3,_heartTimer:null,_closeTimer:null,_closeFlag:!1,_reConnMax:10,_reConnCounter:0,Socket:function(){return null===this._socket&&this.Init(),this._socket},Init:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:null;if(null===this._socket)return!!this.Func.isSupportWs()&&(this._token=e||Object(f.d)(),this.WsInit(),this)},Detect:function(){this.Socket()},Func:{isSupportWs:function(){return"function"==typeof WebSocket||(this.setError("不支持 websocket 通信协议"),!1)},setError:function(e){this.Error=e},getError:function(){return this.Error}},WsInit:function(){var e=this;this._closeFlag=!1,this._socket=new WebSocket("ws://120.26.1.206:8011/ws",this._token),this._socket.onopen=function(t){e._OPEN(t)},this._socket.onclose=function(t){e._CLOSE(t)},this._socket.onmessage=function(t){e._MESSAGE(t)},this._socket.onerror=function(t){e._ERROR(t)}},WsFuncInit:function(){var e=this;this._OPEN=function(t){console.log("OPEN",t),e._reConnCounter=0,e.heart()},this._CLOSE=function(t){console.log("CLOSE",t),!e._closeFlag&&e.reconnect()},this._MESSAGE=function(t){if("function"==typeof e.MessageHandleCallback)return e.MessageHandleCallback(JSON.parse(t.data))},this._ERROR=function(t){console.log("ERROR",t),e.Func.setError(t)},this.heart=function(){e._heartTimer=setInterval(function(){e.Socket().send(s()(E.ping()))},e._heartTimeout)},this.reconnect=function(){e._reConnCounter++,e._reConnCounter>e._reConnMax||(clearInterval(e._heartTimer),e.WsInit())},this.send=function(t){switch(e.Socket().readyState){case WebSocket.OPEN:e.Socket().send(s()(t)),e.MessageHandleCallback(t);break;case WebSocket.CONNECTING:setTimeout(function(){e.send(t)},3e3);break;case WebSocket.CLOSED:case WebSocket.CLOSING:e.Socket().close()}},this.close=function(){clearInterval(e._heartTimer),e._closeFlag=!0,e._socket.close(),e._socket=null}}};j.MessageHandleCallback=function(e){switch(e.type){case u.a.ping:break;case u.a.chat:M(e)}},j.WsFuncInit();var L=j;r.a.use(C.a),r.a.use(C.c),r.a.prototype.$WebSocket=L,r.a.prototype.$Chat=E,r.a.prototype.$SetR=function(e){S.Recipient=e},r.a.prototype.$CurrentUuid=Object(f.f)()?Object(f.b)():0,r.a.prototype.$Log=function(){for(var e,t=arguments.length,n=Array(t),r=0;r<t;r++)n[r]=arguments[r];(e=console).log.apply(e,["[system]:"].concat(n))},r.a.config.productionTip=!1,new r.a({el:"#app",router:w,store:k.c,components:{App:g},template:"<App/>"})},"R+zl":function(e,t){},eeW5:function(e,t){},jBzp:function(e,t){},"l/JR":function(e,t,n){"use strict";var r=n("//Fk"),o=n.n(r),i=n("mtWM"),a=n.n(i),s=n("mw3O"),c=n.n(s),u=n("Fd2+"),f=n("9GPK"),l=n("wy+4"),p=void 0,d=a.a.create({headers:{"Content-Type":"application/x-www-form-urlencoded","X-Requested-With":"XMLHttpRequest"},timeout:6e5,transformRequest:[function(e){return e=c.a.stringify(e)}]});d.interceptors.request.use(function(e){var t=Object(f.d)();return(t||p)&&(e.headers.authorization=t||p),e},function(e){return o.a.error(e)}),d.interceptors.response.use(function(e){var t=e.status,n=e.data;switch(t){case 200:case 201:switch(n.err_code){case 0:return o.a.resolve(n);case-9:return Object(u.b)("错误"),o.a.reject(n);case 401:Object(u.b)("请重新登录"),Object(f.h)(),Object(l.b)(),setTimeout(function(){},1e3);break;default:return Object(u.b)(n.err_msg),o.a.reject(n)}break;case 500:return Object(u.b)("服务器错误"),o.a.reject(n);default:return Object(u.b)("网络连接出错, error_code: "+t),o.a.reject(n)}},function(e){return Object(u.b)("网络连接出错"),o.a.reject(e)}),t.a=d},nbfa:function(e,t){},t5DY:function(e,t,n){"use strict";n.d(t,"b",function(){return o}),n.d(t,"f",function(){return i}),n.d(t,"e",function(){return a}),n.d(t,"c",function(){return s}),n.d(t,"d",function(){return c}),n.d(t,"h",function(){return u}),n.d(t,"j",function(){return f}),n.d(t,"i",function(){return l}),n.d(t,"k",function(){return p}),n.d(t,"g",function(){return d});var r=n("l/JR"),o=function(e){return r.a.post("/chat/getContactInfo",e)},i=function(e){return r.a.get("/getUserInfo",{params:e})},a=function(e){return r.a.post("/chat/getHistory",e)},s=function(e){return r.a.post("/chat/getFriendInfo",e)},c=function(e){return r.a.post("/chat/getFriendNotice",e)},u=function(e){return r.a.post("/chat/searchFriend",e)},f=function(e){return r.a.post("/chat/searchUser",e)},l=function(e){return r.a.post("/chat/searchRoom",e)},p=function(e){return r.a.post("/chat/addFriends",e)},d=function(e){return r.a.post("/chat/addFriendsHandle",e)};t.a={getContacts:function(){return r.a.post("/chat/contacts")},getUserInfo:i,getHistory:a,getFriendNotice:c,handleAddFriendApply:d}},"wy+4":function(e,t,n){"use strict";var r=n("7+uW"),o=n("NYxO"),i=n("Gu7T"),a=n.n(i),s={state:{VideoCall:null,VideoCallRtcCandidate:null,MessageMapList:{0:[]}},mutations:{PUSH_MSG:function(e,t){if(e.MessageMapList[t.recipient]){var n=[].concat(a()(e.MessageMapList[t.recipient]),[t]);r.a.set(e.MessageMapList,t.recipient,n)}else r.a.set(e.MessageMapList,t.recipient,[t])},READ_MSG:function(e,t){t.recipient,t.messageIds},ASSIGN_HISORY:function(e,t){var n=t.recipient,o=t.messages;if(e.MessageMapList[n]){var i=[].concat(a()(o),a()(e.MessageMapList[n]));r.a.set(e.MessageMapList,n,i)}else r.a.set(e.MessageMapList,n,o)},VIDEO_CALL_SET:function(e,t){e.VideoCall=t},VIDEO_CALL_CLEAR:function(e){e.VideoCall=null},VIDEO_CALL_RTC_CANDIDATE_SET:function(e,t){e.VideoCallRtcCandidate=t},VIDEO_CALL_RTC_CANDIDATE_CLEAR:function(e){e.VideoCallRtcCandidate=null}},actions:{push:function(e,t){(0,e.commit)("PUSH_MSG",t)},history:function(e,t){(0,e.commit)("ASSIGN_HISORY",{recipient:t.recipient,messages:t.messages})},videoCallSet:function(e,t){(0,e.commit)("VIDEO_CALL_SET",t)},videoCallClear:function(e){(0,e.commit)("VIDEO_CALL_CLEAR")},videoCallRtcCandidateSet:function(e,t){(0,e.commit)("VIDEO_CALL_RTC_CANDIDATE_SET",t)},videoCallRtcCandidateClear:function(e){(0,e.commit)("VIDEO_CALL_RTC_CANDIDATE_CLEAR")}},getters:{}},c=n("BO1k"),u=n.n(c),f=["token","uuid","name","mobile","mail","avatar","role_id"],l={state:{token:"",uuid:"",name:"",mobile:"",mail:"",avatar:"",role_id:""},mutations:{SET:function(e,t){var n=!0,r=!1,o=void 0;try{for(var i,a=u()(f);!(n=(i=a.next()).done);n=!0){var s=i.value;e[s]=t[s]}}catch(e){r=!0,o=e}finally{try{!n&&a.return&&a.return()}finally{if(r)throw o}}},CLEAR:function(e){var t=!0,n=!1,r=void 0;try{for(var o,i=u()(f);!(t=(o=i.next()).done);t=!0){e[o.value]=""}}catch(e){n=!0,r=e}finally{try{!t&&i.return&&i.return()}finally{if(n)throw r}}}},actions:{},getter:{}},p={state:{FriendMap:{0:{}}},mutations:{UPDATE:function(e,t){r.a.set(e.FriendMap,t.uuid,t)}},actions:{update:function(e,t){(0,e.commit)("UPDATE",t)}},getters:{}};n.d(t,"a",function(){return h}),n.d(t,"b",function(){return m}),r.a.use(o.a);var d=new o.a.Store({modules:{msg:s,user:l,friend:p}}),h=(t.c=d,function(e){d.commit("SET",e)}),m=function(){d.commit("CLEAR")}}},["NHnr"]);
//# sourceMappingURL=app.db1c4abbaade6b8b1b6b.js.map