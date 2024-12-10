package main

import "fmt"

func main() {
	// we make use o new or make

	// garbage colection happesn automatically

	// the differnace betwen the new and make is thatconst
	// new allocates memory but no init  zeroed storage
	// make allocate memory but init non-zero storage

	fmt.Println("Go Memeory management")

	// pointers in go ang
	fmt.Println("pointers in go")
	var ptr *int
	fmt.Println("VAule of pointers", ptr) // nil is default value

	number := 23
	var pt = &number
	fmt.Println(*pt, pt) // pt will retunr address and *pt will return actual value

	*pt = *pt + 2
	fmt.Println(number)

}
