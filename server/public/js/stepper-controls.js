const defaultValue = 5
const controls = {
    up: {
        axis: "y",
        value: defaultValue,
    }, 
    down: {
        axis: "y",
        value: -defaultValue,
    },
    left: {
        axis: "x",
        value: -defaultValue,
    },
    right: {
        axis: "x",
        value: defaultValue,
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
            }, 100)
        })

        function removeInterval() {
            clearInterval(interval)
        }

        button.addEventListener("mouseup", removeInterval)
        button.addEventListener("mouseleave", removeInterval)
    }
}
