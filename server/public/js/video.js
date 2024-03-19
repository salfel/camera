const buttonUp = document.querySelector("button#up");
const buttonDown = document.querySelector("button#down");
const buttonLeft = document.querySelector("button#left");
const buttonRight = document.querySelector("button#right");

const channel = window.location.href.split("/").pop()

const ws = new WebSocket(`ws://${window.location.host}/stepper/${channel}`)

ws.onopen = () => {
    const speed = 15
    const interval = 100

    let intervalId = null
    buttonUp.addEventListener("mousedown", () => intervalId = setInterval(() => sendMove("y", speed), interval))
    buttonUp.addEventListener("mouseup", () => clearInterval(intervalId))
    buttonDown.addEventListener("mousedown", () => intervalId = setInterval(() => sendMove("y", -speed), interval))
    buttonDown.addEventListener("mouseup", () => clearInterval(intervalId))
    buttonLeft.addEventListener("mousedown", () => intervalId = setInterval(() => sendMove("x", -speed), interval))
    buttonLeft.addEventListener("mouseup", () => clearInterval(intervalId))
    buttonRight.addEventListener("mousedown", () => intervalId = setInterval(() => sendMove("x", speed), interval))
    buttonRight.addEventListener("mouseup", () => clearInterval(intervalId))
}

function sendMove(stepper, amount) {
    ws.send(JSON.stringify({"type": "stepper:move", stepper, "amount": amount}))
}
