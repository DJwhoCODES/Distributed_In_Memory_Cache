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

	var strLen int32

	if err := binary.Read(conn, binary.LittleEndian, &strLen); err != nil {
		fmt.Println("Read Length Error:", err)
		return
	}

	buf := make([]byte, strLen)
	if err := binary.Read(conn, binary.LittleEndian, &buf); err != nil {
		fmt.Println("Read String Error:", err)
		return
	}

	str := string(buf[:strLen])
	fmt.Println("Received string:", str)

	binary.Write(conn, binary.LittleEndian, int32(len(buf)))

	binary.Write(conn, binary.LittleEndian, buf)
}
