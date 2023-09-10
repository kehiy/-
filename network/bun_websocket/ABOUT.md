# bun_websocket

An simple chat app (pub/sub feed) with bun js websocket API.

### How to use it?

run the server code:
```
bun run --hot server.ts
```

### Client?

I'm a network and backend programmer, so:

open your browser, open 2 new tabs, inspect and write this codes on both tabs console:

1:
```js
let socket = new WebSocket("ws://localhost:3000/ws");
```

2:
```js
socket.onmessage = function(event) {
  console.log(event.data)
};
```

3:
```js
socket.send("hello")
```


# Result

![Result](https://bafybeifchaykal73ykqcnau5wgyioicona66olg3tinkvrawp3lhdjqvoi.ipfs.w3s.link/Screenshot%20from%202023-09-10%2022-59-01.png)