package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	performPostJson()
}

func performPostJson() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	reqBody, err := string.NewReader(`{
	    "name": "John Doe",
        "age": 30
		}`)
	if err != nil {
		panic(err)
	}
	res, err := http.Post(reqBody, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res)
	cont, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(cont))
}
