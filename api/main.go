package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var fakedb []Course

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	http.HandleFunc("/", serveCourse)

	fmt.Println("server started in port 8000")
	http.ListenAndServe(":8000", nil)

}

func serveCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fakedb)
}

// controllers
