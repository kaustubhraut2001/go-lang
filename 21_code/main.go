package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

}

func perfromfromreq() {
	const myurl = "http://localhost:4000/postformat"
	data := url.Values()
	data.Add("firstname", "Kaustubh")

	res, errr := http.PostForm(myurl, data)

	if errr != nil {
		panic(errr)
	}

	defer res.Body.Close()

	cont, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cont)

}
