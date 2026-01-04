package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("Error listening: ", err)
	}

	defer listener.Close()

	fmt.Println("Load Balancer running on port :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting: ", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// always close the connection at the end
	defer conn.Close()

	conn1, err := net.Dial("tcp", string(conn.RemoteAddr().String()))

	if err != nil {
		fmt.Printf("Error Connecting: ", err)
	}
	defer conn1.Close()

}
