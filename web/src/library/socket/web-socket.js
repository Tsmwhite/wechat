import {Chat} from "../message/chat"
import mesHandleFunc from "../message/handle"
import {GetToken} from "../../utis/cache";

const WEB_SOCKET = {
    Error: "",
    MessageHandleCallback: null,
    _token: "",
    _socket: null,
    _heartTimeout: 8000, // 心跳间隔
    _heartTimer: null, // 保持心跳
    _closeTimer: null, // 延时关闭
    _closeFlag: false, // 主动关闭
    _reConnMax: 60,    // 重连限制
    _reConnCounter: 0, // 重连计数
    _waitSendMessages: [],
    Socket() {
        if (this._socket === null) {
            this.Init()
        }
        return this._socket
    },
    Init(token = null) {
        if (this._socket !== null) {
            return
        }
        if (!this.Func.isSupportWs()) {
            return false
        }
        this._token = token || GetToken()
        this.WsInit()
        return this
    },
    Detect() {
        this.Socket()
    },
    Func: {
        isSupportWs() {
            if (typeof (WebSocket) != "function") {
                this.setError("不支持 websocket 通信协议")
                return false
            }
            return true
        },
        setError(err) {
            this.Error = err
        },
        getError() {
            return this.Error
        },
    },
    WsInit() {
        let _this = this
        this._closeFlag = false
        this._socket = new WebSocket(process.env.WEB_SOCKET_SERVER, this._token)
        this._socket.onopen = function (e) {
            _this._OPEN(e)
        }
        this._socket.onclose = function (e) {
            _this._CLOSE(e)
        }
        this._socket.onmessage = function (e) {
            _this._MESSAGE(e)
        }
        this._socket.onerror = function (e) {
            _this._ERROR(e)
        }
    },
    WsFuncInit() {
        this._OPEN = (e) => {
            console.log('OPEN', e)
            this._reConnCounter = 0
            this.heart()
            if (this._waitSendMessages.length > 0) {
                for (let msg of this._waitSendMessages) {
                    this.send(msg)
                }
            }
        }
        this._CLOSE = (e) => {
            console.log('CLOSE', e)
            !this._closeFlag && this.reconnect()
        }
        this._MESSAGE = (e) => {
            if (typeof this.MessageHandleCallback === "function") {
                return this.MessageHandleCallback(JSON.parse(e.data))
            }
            //clearTimeout(this._closeTimer)
        }

        this._ERROR = (e) => {
            console.log('ERROR', e)
            this.Func.setError(e)
        }
        this.heart = () => {
            this._heartTimer = setInterval(() => {
                this.Socket().send(JSON.stringify(Chat.ping()))
                // this._closeTimer = setTimeout(() => {
                //     this.close()
                // }, this._heartTimeout)
            }, this._heartTimeout)
        }
        this.reconnect = () => {
            this._reConnCounter++
            if (this._reConnCounter > this._reConnMax) {
                return
            }
            setTimeout(() => {
                clearInterval(this._heartTimer)
                this.WsInit()
            }, 1000)
        }
        this.send = (message) => {
            switch (this.Socket().readyState) {
                case WebSocket.OPEN:
                    this.Socket().send(JSON.stringify(message))
                    this.MessageHandleCallback(message)
                    break
                case WebSocket.CONNECTING:
                    setTimeout(() => {
                        this.send(message)
                    }, 3000)
                    break
                case WebSocket.CLOSED:
                case WebSocket.CLOSING:
                    this._waitSendMessages.push(message)
                    this.Socket().close()
                    break
                default:
            }
        }
        this.close = () => {
            //clearTimeout(this._closeTimer)
            clearInterval(this._heartTimer)
            this._closeFlag = true
            this._socket.close()
            this._socket = null
        }
    }
}

WEB_SOCKET.MessageHandleCallback = mesHandleFunc
WEB_SOCKET.WsFuncInit()

export default WEB_SOCKET
