package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in go ")

	content := "File goin this line"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(fileName string) {
	cont, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("file conetent", string(cont)) // if we do not add string then it shows liek this [70 105 108 101 32 103 111 105 110 32 116 104 105 115 32 108 105 110 101]
}
