#! usr/bin/python

from websockets import connect
import json
import asyncio
import requests
import json
import threading
from stepper import run_motor

thread = None
queue = []

def getIp():
    response = requests.get("https://api.ipify.org")
    return response.text

async def main():
    url = "ws://192.168.70.132:3000/stream/test2"
    ip = getIp()
    async with connect(url) as websocket:
        await websocket.send(json.dumps({"type": "register:ip", "ip": ip}))
        message = await websocket.recv()
        print(message)
        while True:
            message = await websocket.recv()
            message = json.loads(message)

            if message["type"] == "stepper:move": 
                global thread
                if thread != None and thread.is_alive(): 
                    queue.append(message["amount"])
                else:
                    thread = threading.Thread(target=motor_thread, args=(message["amount"],))
                    thread.daemon = True
                    thread.start()

def motor_thread(amount): 
    run_motor(amount)

    if len(queue): 
        amount = queue.pop(0)
        thread = threading.Thread(target=motor_thread, args=(amount,))
        thread.daemon = True
        thread.start()
    
if __name__ == "__main__":
    asyncio.run(main())
