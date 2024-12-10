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
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)
}
