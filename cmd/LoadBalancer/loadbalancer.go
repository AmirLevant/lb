package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("Error listening: %s", err)
	}

	defer listener.Close()

	fmt.Println("Load Balancer running on port :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting: %s", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// always close the connection at the end
	defer conn.Close()

	conn1, err := net.Dial("tcp", ":9090")

	if err != nil {
		fmt.Printf("Error Connecting: %s ", err)
	}

	defer conn1.Close()

}
