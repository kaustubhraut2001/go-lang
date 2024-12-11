package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slics in go")
	var fruitList = []string{}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	// we can add many values no needs to woryy about size
	fruitList = append(fruitList, "Mongo", "apple")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 80
	highScores[2] = 90
	highScores[3] = 70
	// highScores[4] = 00

	highScores = append(highScores, 88, 89)
	fmt.Println(highScores)

	sort.Ints(highScores)
	// fmt.Println(sort.IsSorted(highScores))
	// removing vlaue from slice

	var coursec = []string{"react", "js", "python"}
	fmt.Println(coursec)

	var index int = 2
	coursec = append(coursec[:index])
	coursec = append(coursec, coursec[index:]...)
	fmt.Println(coursec)
}
