package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status string
}

func main() {
	fmt.Println("Structs in go lang")

	// no inheritance in go lang

	k := User{"Kaustubh", 20, "kaus@gmail.com", "completed"}

	fmt.Println(k)

	fmt.Printf("struct deatils are %+v\n", k)

}
 