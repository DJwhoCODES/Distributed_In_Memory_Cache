package main

import (
	"encoding/binary"
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

	var num int32

	err := binary.Read(conn, binary.LittleEndian, &num)

	if err != nil {
		fmt.Println("binary.Read error:", err)
		return
	}

	fmt.Println("Received number:", num)

	binary.Write(conn, binary.LittleEndian, int32(num*2))
}
