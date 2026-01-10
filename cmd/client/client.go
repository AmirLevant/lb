package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// the LB port to connect to
	argPortClient := os.Args[1]

	// write the message we want to deliver
	message := "hello my name is Amir"

	// connect with the LB
	conn, err := net.Dial("tcp", ":"+argPortClient)

	// check if the connection is correct
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()
	remoteAdd := conn.RemoteAddr()
	localAdd := conn.LocalAddr()

	fmt.Printf("The remote address in the client is : %s\n", remoteAdd)
	fmt.Printf("The local address in the client is: %s\n", localAdd)

	// write into the connection
	conn.Write([]byte(message))

}
