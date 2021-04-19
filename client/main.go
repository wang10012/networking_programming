package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// creat connection(with server)
	connection, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial failed! Err:", err)
		return
	}
	defer connection.Close()
	// Send and Receive data
	input := bufio.NewReader(os.Stdin)
	for {
		// data processing
		str, _ := input.ReadString('\n')
		str = strings.TrimSpace(str)
		if strings.ToUpper(str) == "QUIT" {
			return
		}
		// Send
		_, err := connection.Write([]byte(str))
		if err != nil {
			fmt.Println("Send failed! Err:", err)
			return
		}
		// Receive
		var buffer [1024]byte
		len, err := connection.Read(buffer[:])
		if err != nil {
			fmt.Println("Recieve failed! Err:", err)
			return
		}
		fmt.Println("Received Data From Server:",string(buffer[:len]))
	}
}
