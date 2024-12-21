package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

// fake db
var course []Course

// middleware or helper -file

func IsEmpty(c *Course) bool {
	return c.CourseID == "" && c.CourseName == ""

}

func main() {
	fmt.Println("Creating APIS in go")

}

func serName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating APIS in go")
	w.Write([]byte("<h1>APIS</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	// id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	fmt.Println(params)

	//loop through course and find the course with the id
	for _, item := range course {
		if item.CourseID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found")

}
