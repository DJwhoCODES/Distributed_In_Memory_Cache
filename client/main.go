package main

import (
	"fmt"
	"net"
)

func main() {
	connObj, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		panic(err)
	}

	defer connObj.Close()

	msg := "Hello from client!"
	connObj.Write([]byte(msg))

	buf := make([]byte, 1024)
	n, err := connObj.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server responded:", string(buf[:n]))
}
