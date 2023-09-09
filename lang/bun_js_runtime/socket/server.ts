const socket = Bun.listen({
    hostname: "localhost",
    port: 8080,
    socket: {

      data(socket, data) {
        console.log("data:", String(data), "from:", socket.remoteAddress)
      }, // message received from client

      open(socket) {
        console.log("new incoming connection:", socket.remoteAddress)
        socket.write("hello, you are connected")
      }, // socket opened

      close(socket) {
        console.log("connection closed:", socket.remoteAddress)
      }, // socket closed

      drain(socket) {}, // socket ready for more data
      error(socket, error) {}, // error handler
   },
});

console.log("listening on port 8080....")
