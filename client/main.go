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

	var num int32 = 60
	binary.Write(connObj, binary.LittleEndian, num)
	fmt.Println("Sent number:", num)

	var resp int32
	binary.Read(connObj, binary.LittleEndian, &resp)

	fmt.Println("Server replied:", resp)
}
