package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	// we listen on port 8080
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server running on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// always close the connection at the end
	defer conn.Close()

	// prevent zombie connections
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error while Reading:", err)
			return
		}
		fmt.Printf("Recieved: %s", buffer[:n])
		conn.Write([]byte("Message received \n"))
	}

}
