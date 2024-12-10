package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting")
	reader := bufio.NewReader(os.Stdin)
	ratings, _ := reader.ReadString('\n')

	fmt.Println(ratings)

	numstring, err := strconv.ParseFloat(strings.TrimSpace(ratings), 64)
	if err != nil {
		fmt.Println("Error parsing", err)
	} else {
		fmt.Println("Number: ", numstring+1)
	}

}
