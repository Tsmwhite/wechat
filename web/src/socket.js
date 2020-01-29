const SocketConfig =  {
    Host:"127.0.0.1",
    Port:"8088",
}

var WsUrl = "ws://"+SocketConfig.Host+":"+SocketConfig.Port+"/ws"
window.console.log(WsUrl)
var WsSocket;
var HeartTime = 6000;
var WsHeartTime;
var WsCloseTime;
var WsToken;
var CloseStatus = false;    //是否主动关闭
var ReConnCounter = 0;      //重连计数器

function WsInit(token){
   WsToken = token
   try{
       if( typeof(WebSocket) != "function" ) {
           throw "您的浏览器不支持 websocket 通信协议，请使用 chrome 或 firefox"
       }

       if (token != "") {
           WsSocket = new WebSocket(WsUrl,token)
       }else{
           WsSocket = new WebSocket(WsUrl)
       }
       WsSocket.onopen    = WsOpen
       WsSocket.onerror   = WsError
       WsSocket.onclose   = WsClose
       WsSocket.onmessage = WsMessage
   }catch (err) {
       return err
   }
}

//重连
function reconnect() {
    ReConnCounter++
    if (ReConnCounter > 100) {
        return
    }
    clearInterval(WsHeartTime)
    WsInit(WsToken)
}

function WsOpen() {
    //开启心跳
    Heart()
    //重连计数器置0
    ReConnCounter = 0
}

function WsError(err) {
    window.console.log("WsError:"+err)
    reconnect()
}

function WsClose() {
    !CloseStatus && reconnect()
}

function WsMessage(e) {
    clearTimeout(WsCloseTime)
    window.console.log(e)
    var data = e.data
    if (data != "pong") {
        data = JSON.parse(data)
        window.console.log(data)
    }
}


function SendMsg(data) {
    if (WsSocket !== null && WsSocket.readyState === 3) {
        WsSocket.close()
        //WsInit(WsToken) //重连
    } else if (WsSocket.readyState === 1) {
        WsSocket.send(JSON.stringify(data))
    } else if (WsSocket.readyState === 0) {
        setTimeout(() => {
            WsSocket.send(JSON.stringify(data))
        }, 3000)
    }
}

function Heart() {
    WsHeartTime = setInterval(() => {
        WsSocket.send('ping')
        WsCloseTime = setTimeout(()=> {
            WsSocket.close()
        },HeartTime)
    }, HeartTime)
}


var Ws = {
    Init:WsInit,
    Send:SendMsg
}

export default Ws