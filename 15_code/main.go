package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed last")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Defer in go")
	mydefer()

	// defer make use of last in first out fif fashion so the last one will pritn first and then previous things

}

func mydefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // this will print in reverse order as it is defered
	}
}
