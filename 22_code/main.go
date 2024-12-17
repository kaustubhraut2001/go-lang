package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	name     string
	price    int
	platform string
	password string
	tags     []string
}

func main() {
	fmt.Println("handling jso data")
	encodejson()

}

func encodejson() {
	courssee := []course{
		{"Reactjs", 233, "Learcodeonline", "abc123", []string{"web-dev", "dsa"}},
		{"Python", 233, "Learcodeonline", "abc123", []string{"web-dev", "dsa"}},
		{"C++", 233, "Learcodeonline", "abc123", nil},
	}

	// packaging data as an json data
	finaljson, _ := json.MarshalIndent(courssee, "", "\t")
	fmt.Printf("%s\n", finaljson)
}
