const socket = Bun.connect({
    hostname: "localhost",
    port: 8080,
  
    socket: {

      data(socket, data) {},

      open(socket) {
        console.log("connected to:", socket.remoteAddress)
        socket.write("Hello")
      },

      close(socket) {
        console.log("disconnected from:", socket.remoteAddress)
      },

      drain(socket) {},
      error(socket, error) {},
  
      // client-specific handlers
      connectError(socket, error) {}, // connection failed
      end(socket) {}, // connection closed by server
      timeout(socket) {}, // connection timed out
    },
});