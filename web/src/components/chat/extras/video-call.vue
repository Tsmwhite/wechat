<template>
    <div v-show="flag"
         class="video-call-box">
        <div @click="close">关闭</div>
        <video autoplay playsinline controls id="localVideo" ref="localVideoRef" style="width: 400px"></video>
        <template v-for="key in memberKeys">
            <video autoplay
                   playsinline
                   controls
                   id="remoteVideo"
                   :key="key"
                   :ref="'remote'+key+'VideoRef'"
                   style="width: 400px"></video>
        </template>
    </div>
</template>

<script>
import {VideoCallStatus} from "../../../library/message/const";
import {NewWebRTC} from "../../../library/webrtc/main";

export default {
    name: "video-call",
    props: {
        roomData: {
            type: Object,
            required: true,
        }
    },
    data() {
        return {
            flag: false,
            peerMap: {},
            memberKeys: [],
            localStream: null,
            localVideo: null,
            remoteVideo: null,
            localMediaConstraints: {
                video: true,
                audio: true,
            }
        }
    },
    mounted() {
        console.log("open chat video call")
        this.init()
    },
    computed: {
        videoCallRtcCandidate() {
            let storeData = this.$store.state
            return storeData.msg.VideoCallRtcCandidate
        },
        videoCallMsg() {
            let storeData = this.$store.state
            return storeData.msg.VideoCall
        },
    },
    watch: {
        videoCallRtcCandidate() {
            if (this.videoCallRtcCandidate && this.videoCallRtcCandidate.status === VideoCallStatus.candidate) {
                this.handleSendCandidate()
            }
        },
        videoCallMsg() {
            if (this.videoCallMsg) {
                switch (Number(this.videoCallMsg.status)) {
                    case VideoCallStatus.agree:
                        // 对方同意视频通话
                        // 创建连接
                        this.$Log("对方同意视频通话", this.videoCallMsg)
                        this.createVideoCall()
                        break
                    case VideoCallStatus.createConn:
                        // 收到对方信令(offer)
                        // 创建连接
                        this.$Log("createConn", this.videoCallMsg)
                        if (!this.flag) {
                            this.videoCallOfferHandle()
                        }
                        break
                    case VideoCallStatus.createConnConfirm:
                        // 收到对方应答
                        this.$Log("收到对方应答", this.videoCallMsg)
                        this.videoCallAnswerHandle()
                        break
                }
            }
        },
    },
    methods: {
        init() {
            this.localVideo = this.$refs.localVideoRef
            this.memberKeys.push(this.roomData['object'])
        },
        open() {
            if (this.videoCallMsg) {
                return false
            }
            this.getLocalMediaData(() => {
                this.flag = true
                let message = this.$Chat.videoCall()
                this.$WebSocket.send(message)
            })
            return true
        },
        close() {
            this.flag = false
        },
        handleRemoteStream(event, pc) {
            console.log("handleRemoteStream", event, pc)
            let remoteVideo = this.$refs['remote' + pc.key + 'VideoRef'][0]
            if (this.memberKeys.includes(pc.key) && remoteVideo) {
                remoteVideo.srcObject = event.streams[0];
                remoteVideo.addEventListener('loadedmetadata', () => {
                    remoteVideo.play()
                })
            }
        },
        handleSendCandidate() {
            if (this.videoCallRtcCandidate && this.videoCallRtcCandidate.status === VideoCallStatus.candidate) {
                let message = JSON.parse(this.videoCallRtcCandidate.content)
                let candidate = new RTCIceCandidate({
                    sdpMLineIndex: message.label,
                    candidate: message.candidate
                })
                this.peerMap[this.roomData['object']].RTC.addIceCandidate(candidate).catch(err => {
                    console.log('addIceCandidate-error', err)
                })
            }
        },
        handleIcecandidate(data, pc) {
            let message = this.$Chat.videoCall(VideoCallStatus.candidate)
            message.parent = this.videoCallMsg.parent || this.videoCallMsg.uuid
            message.recipient = this.videoCallMsg.recipient
            message.content = JSON.stringify(data)
            this.$WebSocket.send(message)
        },
        createPeerConn(key, isSelf) {
            if (!this.peerMap[key]) {
                let pc = NewWebRTC({
                    key: key,
                    isSelf: isSelf,
                    handleRemoteStream: this.handleRemoteStream,
                    handleIcecandidate: this.handleIcecandidate
                })
                if (isSelf) {
                    this.localStream.getTracks().map(track => {
                        pc.RTC.addTrack(track, this.localStream)
                    })
                    pc.CreateOffer((description) => {
                        this.$Log("createVideoCall -> createOffer -> description:", description)
                        // 传输 description
                        let message = this.$Chat.videoCall(VideoCallStatus.createConn)
                        message.parent = this.videoCallMsg.parent || this.videoCallMsg.uuid
                        message.recipient = this.videoCallMsg.recipient
                        message.content = JSON.stringify(description)
                        this.$WebSocket.send(message)
                    })
                }
                this.peerMap[key] = pc
            }
            return this.peerMap[key]
        },
        createVideoCall() {
            // 作为发送方
            // @1、创建连接
            // @2、生成offer
            // @3、设置本地会话描述
            // @4、向对方发送信令（offer）
            const PC = this.createPeerConn(this.roomData['object'], true)
            this.peerMap[this.roomData['object']] = PC
        },
        videoCallAnswerHandle() {
            // 处理对方应答
            const otherUuid = this.videoCallMsg.sender
            const answerDescription = JSON.parse(this.videoCallMsg.content)
            const PC = this.peerMap[otherUuid]
            PC.HandleAnswer(answerDescription, () => {

            })
        },
        getLocalMediaData(callback = null) {
            navigator.mediaDevices.getUserMedia(this.localMediaConstraints).then(stream => {
                this.localStream = stream
                this.localVideo.srcObject = stream
                this.localVideo.play()
                typeof callback === "function" && callback(stream)
            }).catch(err => {
                this.$Log("本地视频采集错误「" + err.toString() + "」")
                this.$toast.fail("视频通话暂不可用")
            })
        },
        videoCallOfferHandle() {
            this.getLocalMediaData((stream) => {
                // 作为接收方，收到信令(offer)进行应答
                // @1、创建连接
                // @2、生成answer
                // @3、设置本地会话描述
                // @4、向对方发送信令（answer）
                // @5、创建远端会话连接
                // @6、处理对方信令（offer）
                // @7、设置远端会话描述（offer）
                this.$Log("signHandle videoCallMsg:", this.videoCallMsg)
                if (this.videoCallMsg && this.videoCallMsg.status === VideoCallStatus.createConn) {
                    const offerDescription = JSON.parse(this.videoCallMsg.content)
                    this.flag = true
                    const PC = this.createPeerConn(this.roomData['object'], false)
                    this.peerMap[this.roomData['object']] = PC
                    this.localStream.getTracks().map(track => {
                        PC.RTC.addTrack(track, this.localStream)
                    })
                    this.peerMap[this.roomData['object']].HandleOffer(offerDescription, (answerDescription) => {
                        this.$Log("videoCallOfferHandle -> createAnswer -> description:", answerDescription)
                        let message = this.$Chat.videoCall(VideoCallStatus.createConnConfirm)
                        message.parent = this.videoCallMsg.parent || this.videoCallMsg.uuid
                        message.recipient = this.videoCallMsg.recipient
                        message.content = JSON.stringify(answerDescription)
                        this.$WebSocket.send(message)
                    })
                }
            })
        },
    }
}
</script>

<style scoped lang="less">
.video-call-box {
    position: absolute;
    z-index: 999;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background: rgba(0, 0, 0, 0.9);

    video {
        max-width: 88%;
    }
}
</style>
