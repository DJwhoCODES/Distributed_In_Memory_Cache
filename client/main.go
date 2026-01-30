package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	connObj, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		panic(err)
	}

	defer connObj.Close()

	msg := "hello world"
	data := []byte(msg)

	binary.Write(connObj, binary.LittleEndian, int32(len(data)))

	binary.Write(connObj, binary.LittleEndian, data)

	fmt.Println("Sent:", msg)

	var respLen int32

	binary.Read(connObj, binary.LittleEndian, &respLen)

	respBuf := make([]byte, respLen)
	binary.Read(connObj, binary.LittleEndian, &respBuf)

	fmt.Println("Received:", string(respBuf[:respLen]))
}
