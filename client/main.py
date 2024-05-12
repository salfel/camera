from websockets import connect, exceptions
from dotenv import load_dotenv
import json
import asyncio
import json
import os
import socket
from stepper import run_motor

load_dotenv()

thread = None
queue = []

def getIp():
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect(("8.8.8.8", 80))
    ip = s.getsockname()[0]
    s.close()
    return ip

async def main():
    server_ip = os.getenv("SERVER_IP")
    url = "ws://" + str(server_ip) + ":3000/stream/" + (os.getenv("CHANNEL") or "")
    ip = getIp()
    async for websocket in connect(url):
        try: 
            await websocket.send(json.dumps({"type": "register:ip", "ip": ip, "authToken": os.getenv("AUTH_TOKEN")}))
            message = await websocket.recv()
            print(message)
            while True:
                message = await websocket.recv()
                message = json.loads(message)

                if message["type"] == "stepper:move": 
                    run_motor(message["axis"], message["amount"])
        except exceptions.ConnectionClosed as e:
            print(e)

if __name__ == "__main__":
    asyncio.run(main())
