package main

import (
	"fmt"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Hello, Fir!")

	// User input
	var input int
	fmt.Scanf("%d", &input)
	fmt.Printf("You entered: %d\n", input)

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	// if err != nil {
	// 	log.Fatalf("Failed to make HTTP request: %v", err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Failed to read response body: %v", err)
	// }

	// var todos []Todo
	// err = json.Unmarshal(body, &todos)
	// if err != nil {
	// 	log.Fatalf("Failed to unmarshal JSON: %v", err)
	// }

	// for _, todo := range todos {
	// 	fmt.Printf("%+v\n", todo)
	// }

	mp := map[string]int{
		"completed":     8,
		"not completed": 0,
	}
	fetchmapvalues := mp["completed"]
	fmt.Println(fetchmapvalues)

}
