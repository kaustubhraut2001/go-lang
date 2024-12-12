package main

import "fmt"

func main() {
	fmt.Println("consitinal statemts in go")
	loginCount := 22
	var result string

	if loginCount >= 10 {
		result = "Welcome!"
	} else {
		result = "Access denied!"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if nums := 3; nums < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than or equal to 10")
	}

	// if err != nil {
	// 	fmt.Println("Error occurred:", err)
	// }

}
