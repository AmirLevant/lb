package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println("position 0 is :  ", os.Args[0])
	fmt.Println("position 1 is :  ", os.Args[1])
	fmt.Println("position 2 is :  ", os.Args[2])
	fmt.Println("position 3 is :  ", os.Args[3])

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
}
