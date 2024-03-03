from websockets import connect
import json
import asyncio
import socket

async def main(): 
    url = "ws://localhost:3000/stream/test2"
    ip = socket.gethostbyname(socket.gethostname())
    async with connect(url) as websocket: 
        await websocket.send(json.dumps({ "type": "register:ip", "ip": ip}))
        message = await websocket.recv()
        print(message)
        while True: 
            pass

if __name__ == "__main__":
    asyncio.run(main())
