package main

import "fmt"

func main() {
	fmt.Println("Maps in go")

	fruitMap := make(map[string]string)
	fruitMap["Apple"] = "red"
	fruitMap["Banana"] = "yellow"
	fmt.Println(fruitMap["Apple"], fruitMap)

	delete(fruitMap, "Banana") // to delete some things from map
	fmt.Println(fruitMap)

	// loops on hashmap

	for key, value := range fruitMap {
		fmt.Printf("Key : %s , Value : %s \n", key, value)
	}
	for _, value := range fruitMap {
		fmt.Printf("Key : %s , Value : %s \n", value)
	}
}
