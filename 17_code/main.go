package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web requests in go")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	defer res.Body.Close()
	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
