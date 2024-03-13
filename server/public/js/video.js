const buttonUp = document.querySelector("button#up");
const buttonDown = document.querySelector("button#down");
const buttonLeft = document.querySelector("button#left");
const buttonRight = document.querySelector("button#right");

const channel = window.location.href.split("/").pop()

const ws = new WebSocket(`ws://${window.location.host}/stepper/${channel}`)

ws.onopen = () => {
    ws.send(JSON.stringify({"type": "stepper:move", "stepper": "1", "amount": 100}))
    console.log("sent")
}

function sendMove(stepper, direction) {
    ws.send(JSON.stringify({"type": "stepper:move", stepper, "amount": direction ? 100 : -100}))
}
