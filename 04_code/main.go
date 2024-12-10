package main

import (
	"fmt"
	"time"
)

func main() {
	// time
	// fmt.Println("Time: ", time.Now)
	parseTime := time.Now()
	fmt.Println(parseTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2022, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)

}
