const buttonUp = document.querySelector("button#up");
const buttonDown = document.querySelector("button#down");
const buttonLeft = document.querySelector("button#left");
const buttonRight = document.querySelector("button#right");

const channel = window.location.href.split("/").pop()

const ws = new WebSocket(`ws://${window.location.host}/stepper/${channel}`)

ws.onopen = () => {
    const speed = 3

    let interval = null
    buttonUp.addEventListener("mousedown", () => interval = setInterval(() => sendMove("y", speed), 50))
    buttonUp.addEventListener("mouseup", () => clearInterval(interval))
    buttonDown.addEventListener("mousedown", () => interval = setInterval(() => sendMove("y", -speed), 50))
    buttonDown.addEventListener("mouseup", () => clearInterval(interval))
    buttonLeft.addEventListener("mousedown", () => interval = setInterval(() => sendMove("x", -speed), 50))
    buttonLeft.addEventListener("mouseup", () => clearInterval(interval))
    buttonRight.addEventListener("mousedown", () => interval = setInterval(() => sendMove("x", speed), 50))
    buttonRight.addEventListener("mouseup", () => clearInterval(interval))
}

function sendMove(stepper, amount) {
    ws.send(JSON.stringify({"type": "stepper:move", stepper, "amount": amount}))
}
