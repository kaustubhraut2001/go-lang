package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("go cod")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	// comma ok syntax
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)
	fmt.Printf("Type of this is %T \n", name)
}
