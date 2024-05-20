from piservo import Servo

motors = {
    "x": {
        "motor": Servo(18),
        "pin": 18,
        "value": 0,
        "running": False
    },
    "y": {
        "motor": Servo(23),
        "pin": 23,
        "value": 0,
        "running": False
    }
}

def run_motor(axis: str, amount: int):
    motor = motors[axis]
    motor["value"] = amount
    motor["motor"].set_angle(amount)
