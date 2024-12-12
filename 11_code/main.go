package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case")

	rand.Seed(time.Now().UnixNano())

	dics := rand.Intn(6) + 1

	switch dics {
	case 1:
		fmt.Println("dice value is 1")

	case 2:
		fmt.Println("dice valeus is 2")

	case 3:
		fmt.Println("dice values is 3")
	case 4:
		fmt.Println("dice values is 4")
	case 5:
		fmt.Println("dice values is 5")
	case 6:
		fmt.Println("dice values is 6")
	}
}
