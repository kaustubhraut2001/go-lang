package main

import "fmt"

type Person struct {
	name  string
	email string
	age   int
}

func (p Person) sayHello() {
	fmt.Println("Hello from", p.name)
}

func main() {
	// P := Person({name : "John Smith " , emaail : "as@gmail.com" , age: 10})

	// var raju Person
	// raju.name ="Raju"

	// fmt.Println("Struvt")
	// sayHello()
	var raju Person
	raju.name = "Raju"
	raju.sayHello()

}
