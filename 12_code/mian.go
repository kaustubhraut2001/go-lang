package main

import "fmt"

func main() {
	fmt.Println("Loops in go")
	dyas := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for index, day := range dyas {
		fmt.Printf("%d. %s\n", index+1, day)
	}

	for d := 0; d < len(dyas); d++ {
		fmt.Printf("%d. %s\n", d+1, dyas[d])
	}

	for i := range dyas {
		fmt.Println(dyas[i])
	}

	for _, day := range dyas {
		fmt.Println(day)
	}

	rougValue := 1
	for rougValue < 10 { // kind of simliar to while
		if rougValue == 5 {
			rougValue++
			continue
		}
		fmt.Println(rougValue)
		rougValue++
	}

}
