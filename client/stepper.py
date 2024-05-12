from RpiMotorLib import RpiMotorLib

motors = {
    "x": {
        "motor": RpiMotorLib.BYJMotor("motorX"),
        "pins": [18, 23, 24, 25],
        "value": 0,
        "running": False
    },
    "y": {
        "motor": RpiMotorLib.BYJMotor("motorY"),
        "pins": [17, 22, 10, 9],
        "value": 0,
        "running": False
    }
}

def run_motor(axis: str, amount: int):
    motors[axis]["value"] += amount
    if motors[axis]["running"]:
        return

    value = motors[axis]["value"]

    motors[axis]["running"] = True
    motor = motors[axis]["motor"]
    motor.motor_run(motors[axis]["pins"], steps=abs(value), ccwise=value < 0)

    motors[axis]["value"] -= value
    motors[axis]["running"] = False

    if motors[axis]["value"] != 0:
        run_motor(axis, motors[axis]["value"])
