package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	perfromGetdata()                     // calling the function
	fmt.Println("Finished getting data") // print after getting data

}
func perfromGetdata() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url) //

	if err != nil {
		panic(err)
	}

	cont, _ := ioutil.ReadAll(res.Body)

	var ress strings.Builder

	fmt.Println(string(cont))
	byteCount, _ := ress.Write(cont)
	fmt.Println(byteCount)

	fmt.Println(ress.String())
	fmt.Println(ress.Len())

	defer res.Body.Close() // closing the response body
}
