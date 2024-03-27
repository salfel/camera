let stream = new MediaStream();
let suuid = window.location.pathname.split("/").pop();

let config = {
  iceServers: [{
    urls: ["stun:stun.l.google.com:19302"]
  }]
};

const pc = new RTCPeerConnection(config);
pc.onnegotiationneeded = handleNegotiationNeededEvent;

pc.ontrack = function(event) {
    stream.addTrack(event.track)
    const videoElement = document.querySelector("video")
    videoElement.srcObject = stream
}

async function handleNegotiationNeededEvent() {
    let offer = await pc.createOffer()
    await pc.setLocalDescription(offer)
    getRemoteSdp();
}

async function getCodecInfo() {
    const response = await fetch(`/stream/codec/${suuid}`)
    const data = await response.json()
    for (const value of data) {
        pc.addTransceiver(value.Type, {direction: "sendrecv"})
    }
}

async function getRemoteSdp() {
    const formaData = new FormData()
    formaData.append("suuid", suuid)
    formaData.append("data", btoa(pc.localDescription.sdp))

    const response = await fetch("/stream/receiver/" + suuid, {
        method: "POST",
        body: formaData
    });

    pc.setRemoteDescription(new RTCSessionDescription({
      type: 'answer',
      sdp: atob(await response.text())
    }))
}

getCodecInfo()
