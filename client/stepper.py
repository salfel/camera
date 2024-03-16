from RpiMotorLib import RpiMotorLib

motor1 = RpiMotorLib.BYJMotor("motor1")

Motor1Pins = [18, 23, 24, 25]

def run_motor(value):
    motor1.motor_run(Motor1Pins, steps=abs(value), ccwise=value < 0)
