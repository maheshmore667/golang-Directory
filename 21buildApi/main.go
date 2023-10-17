package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to build API")

	courses = append(courses, Courses{CourseId: "2", CourseName: "Angular", Price: 800, Author: Author{AuthorName: "Nahesh More", Website: "go.dev"}})
	courses = append(courses, Courses{CourseId: "3", CourseName: "React", Price: 900, Author: Author{AuthorName: "Nahesh More", Website: "go.dev"}})

	//handlers
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/getCourse/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/createCourse", createOneCourse).Methods("POST")
	r.HandleFunc("/updateCourse", updateCourse).Methods("PUT")
	r.HandleFunc("/deleteCourse", deleteCourse).Methods("DELETE")

	//serve
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Go Server By Mahesh</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course found with the given Id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var course Courses
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty payload")
	}
	json.NewDecoder(r.Body).Decode(&course)
	if isCourseEmpty(course) {
		json.NewEncoder(w).Encode("invalid payload")
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var course Courses
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Payload")
	}
	json.NewDecoder(r.Body).Decode(&course)
	isCoursePresent := false
	for index, courseInfo := range courses {
		if courseInfo.CourseId == course.CourseId {
			isCoursePresent = true
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Courses = course
			newCourse.CourseId = course.CourseId
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}
	if !isCoursePresent {
		json.NewEncoder(w).Encode("No Course is present with give courseId")
		return
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete the course")
	w.Header().Add("Content-Type", "application/json")
	var course Courses
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Payload")
	}
	json.NewDecoder(r.Body).Decode(&course)
	isCoursePresent := false
	for index, courseInfo := range courses {
		if courseInfo.CourseId == course.CourseId {
			isCoursePresent = true
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully deleted the course")
			break
		}
	}
	if !isCoursePresent {
		json.NewEncoder(w).Encode("No Course is present with give courseId")
		return
	}

}
