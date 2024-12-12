package main

import "fmt"

func addNumbers(number1 int, number2 int) int {
	return number1 + number2

}

func proAdder(values ...int) (int, string) {
	result := 1
	for _, val := range values {
		result *= val
	}
	return result, "Hi"

}

func main() {
	fmt.Println("Function in go")
	result := addNumbers(10, 20)
	fmt.Println(result)

	res, msg := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(res, msg)

}
