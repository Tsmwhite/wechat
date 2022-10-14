const defaultOfferOptions = {
    offerToReceiveVideo: true,
    offerToReceiveAudio: true,
}

const randStr = () => {
    return "tsm:" + Math.random().toString().substr(2, 16)
}

const NewPeerConn = (handleRemoteStream, handleIcecandidate) => {
    const PC = new RTCPeerConnection({
        iceServers: [
            {urls: 'stun:stun.l.google.com:19302'},
            {urls: 'stun:stun1.l.google.com:19302'}
        ]
    })
    PC.addEventListener("icecandidate", event => {
        console.log('icecandidate event:', event)
        if (event.candidate) {
            handleIcecandidate({
                type: 'candidate',
                label: event.candidate.sdpMLineIndex,
                id: event.candidate.sdpMid,
                candidate: event.candidate.candidate
            })
        } else {
            console.log('End of candidates.')
        }
    })
    PC.addEventListener("track", event => {
        if (typeof handleRemoteStream === "function") {
            handleRemoteStream(event)
        }
    })
    PC.setLocalDesc = (description, success) => {
        return PC.setLocalDescription(description).then(() => {
            success(description)
        }).catch((err) => console.log("setLocalDesc error:", err))
    }
    return PC
}

export const NewWebRTC = (option = {
    key: randStr(),
    isSelf: false,
    handleRemoteStream: null,
    handleIcecandidate: null,
}) => {
    const RTCPeer = {
        RTC: NewPeerConn((event) => {
            typeof option.handleRemoteStream === "function" && option.handleRemoteStream(event, RTCPeer)
        }, (data) => {
            typeof option.handleIcecandidate === "function" && option.handleIcecandidate(data, RTCPeer)
        }),
        key: option.key,
        isSelf: option.isSelf || false,
        CreateOffer: (
            success,
            options = defaultOfferOptions,
        ) => {
            console.log("this", RTCPeer)
            RTCPeer.RTC.createOffer(options).then(description => {
                RTCPeer.RTC.setLocalDesc(description, success).then(() => {
                    console.log("createOffer setLocalDesc OK")
                })
            }).catch((err) => console.log("createOffer error:", err))
        },
        CreateAnswer: (success) => {
            RTCPeer.RTC.createAnswer().then((description) => {
                RTCPeer.RTC.setLocalDesc(description, success).then(() => {
                    console.log("CreateAnswer setLocalDesc OK")
                })
            }).catch((err) => console.log("createAnswer error:", err))
        },
        HandleOffer: (
            offerDescription,
            success
        ) => {
            const RTCDesc = new RTCSessionDescription(offerDescription)
            RTCPeer.RTC.setRemoteDescription(RTCDesc).then(() => {
                RTCPeer.CreateAnswer(success)
            }).catch((err) => console.log("HandleOffer setRemoteDescription error:", err))
        },
        HandleAnswer: (answerDescription, success) => {
            const RTCDesc = new RTCSessionDescription(answerDescription)
            RTCPeer.RTC.setRemoteDescription(RTCDesc).then(() => {
                success()
            }).catch((err) => console.log("HandleAnswer setRemoteDescription error:", err))
        },
    }
    return RTCPeer
}
