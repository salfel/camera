from RpiMotorLib import RpiMotorLib

motors = {
    "x": {
        "motor": RpiMotorLib.BYJMotor("motorX"),
        "pins": [18, 23, 24, 25]
    },
    "y": {
        "motor": RpiMotorLib.BYJMotor("motorY"),
        "pins": [17, 22, 10, 9]
    }
}

def run_motor(axis: str, amount: int):
    motor = motors[axis]["motor"]
    motor.motor_run(motors[axis]["pins"], steps=abs(amount), ccwise=amount < 0)
