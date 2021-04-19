package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buffer [128]byte
		len, err := reader.Read(buffer[:]) //Read(slice)
		if err != nil {
			fmt.Println("read failed! Err:", err)
			break
		}
		recv := string(buffer[:len])
		fmt.Println("Received Data From Client:", recv)
		conn.Write([]byte("Sucess!"))
	}
}

// server
func main() {
	// start service
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed! Err:", err)
		return
	}
	for {
		// wait connection
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Connection failed! Err:", err)
			continue
		}
		// create gorountine
		go process(connection)
	}
}
