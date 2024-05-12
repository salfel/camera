const defaultValue = 3
const controls = {
    up: {
        axis: "y",
        amount: defaultValue,
    }, 
    down: {
        axis: "y",
        amount: -defaultValue,
    },
    left: {
        axis: "x",
        amount: -defaultValue,
    },
    right: {
        axis: "x",
        amount: defaultValue,
    },
}

const socket = new WebSocket(`ws://${window.location.host}/stepper/${suuid}`)

socket.onopen = () => {
    for (const [key, control] of Object.entries(controls)) {
        const button = document.getElementById(key)

        let interval = null

        button.addEventListener("mousedown", () => {
            interval = setInterval(() => {
                socket.send(JSON.stringify({
                    type: "stepper:move",
                    ...control,
                }))
                console.log("sending " + key)
            }, 35)
        })

        function removeInterval() {
            clearInterval(interval)
        }

        button.addEventListener("mouseup", removeInterval)
        button.addEventListener("mouseleave", removeInterval)
    }
}

socket.onclose = () => {
    console.log("")
}
