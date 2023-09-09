package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("SET amir Kay 100"))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
