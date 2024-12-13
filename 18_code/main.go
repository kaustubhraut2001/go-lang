package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.com"

func main() {
	fmt.Println("Url handing in go")
	fmt.Println(myurl)

	res, _ := url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qParam := res.Query()

	fmt.Printf("Tyope %T\n", qParam) // urlValues

	for _, val := range qParam {
		fmt.Printf("Key : %s , Value : %s\n", val[0], val[1])
	}

	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		Path:     "/",
		RawQuery: "q=golang",
	}
	fmt.Println(partsofUrl.String())

}
