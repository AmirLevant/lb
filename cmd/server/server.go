package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	argPort := os.Args[1]

	// we listen on port 8080
	listener, err := net.Listen("tcp", ":"+argPort)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server running on port :" + argPort)

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

	remoteAdd := conn.RemoteAddr()
	localAdd := conn.LocalAddr()
	fmt.Printf("The remote address in the server is : %s\n", remoteAdd)
	fmt.Printf("The local address in the server is: %s\n", localAdd)

	message := make([]byte, 1024)

	_, err := conn.Read(message)

	if err != nil && err != io.EOF {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Message Recieved: %s", message)

}
