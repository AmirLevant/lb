package main

import (
	"bufio"
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

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err)
			}
			return
		}

		fmt.Printf("Message Recieved: %s", message)

	}

}
