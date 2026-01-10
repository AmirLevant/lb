package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// load balancer runs on 8080
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

	remoteAdd := conn.RemoteAddr()
	localAdd := conn.LocalAddr()

	fmt.Printf("The remote address in the client is : %s\n", remoteAdd)
	fmt.Printf("The local address in the client is: %s\n", localAdd)

	RxBuffer := make([]byte, 1024)

	_, err := conn.Read(RxBuffer)

	if err != nil {
		log.Printf("Connection error: %v", err)
		return
	}

	fmt.Printf("I got from the client: %s", RxBuffer)

	conn1, err := net.Dial("tcp", ":9090")

	if err != nil {
		fmt.Printf("Error Connecting: %s ", err)
	}

	defer conn1.Close()

	conn1.Write(RxBuffer)

}
