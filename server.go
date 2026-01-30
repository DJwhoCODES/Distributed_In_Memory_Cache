package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5050")

	if err != nil {
		panic(err)
	}

	fmt.Println("TCP Server Is Listening On Port 5050")

	for {
		connObj, err := listener.Accept()

		if err != nil {
			fmt.Println("Error Accepting Request:", err)
			continue
		}

		fmt.Println("Client Connected!")

		go handleConn(connObj)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error Reading:", err)
		return
	}

	fmt.Println("Received:", string(buf[:n]))

	conn.Write([]byte("OK"))
}
