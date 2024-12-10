package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Cherry"
	// fruitList[3] = "peaches"

	fmt.Println(len(fruitList))

}
