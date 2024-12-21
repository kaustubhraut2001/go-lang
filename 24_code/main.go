package main

import "fmt"

// Model for course -file

type Course struct {
	CourseID   string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	Price      float64 `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

//fake db
var course []Course

// middleware or helper -file

func IsEmpty(c *Course) bool {
	return c.CourseID == "" && c.CourseName == ""

}

func main() {
	fmt.Println("Creating APIS in go")

}
