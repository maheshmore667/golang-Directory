package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Courses struct {
	CourseId   string
	CourseName string
	Price      int
	Author     Author
}

type Author struct {
	AuthorName string
	Website    string
}

// dummy Database
var courses []Courses

// middleware
func isCourseEmpty(c Courses) bool {
	return c.CourseName == "" && c.CourseId == ""
}

func main() {
	fmt.Println("Welcome to build API")
}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Go Server By Mahesh</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
