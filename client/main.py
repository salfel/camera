#! /usr/bin/python

from websockets import connect
import json
import asyncio
import requests


async def getIp():
    response = await requests.get("https://api.ipify.org")
    return response.text


async def main():
    url = "ws://192.168.253.132:3000/stream/test2"
    ip = await getIp()
    async with connect(url) as websocket:
        await websocket.send(json.dumps({"type": "register:ip", "ip": ip}))
        message = await websocket.recv()
        print(message)
        while True:
            pass

if __name__ == "__main__":
    asyncio.run(main())
