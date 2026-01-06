package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// the LB port to connect to
	argPortClient := os.Args[1]

	// write the message we want to deliver
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Send a message:")
	message, _ := reader.ReadString('\n')

	// connect with the LB
	conn, err := net.Dial("tcp", ":"+argPortClient)

	// check if the connection is correct
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()

	// write into the connection
	conn.Write([]byte(message))

}
