package main

import "fmt"

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

//middleware
func isCourseEmpty(c Courses) bool {
	return c.CourseName == "" && c.CourseId == ""
}

func main() {
	fmt.Println("Welcome to build API")
}
