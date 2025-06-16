package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var jsonData []Task

func init() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error while reading data", err.Error())
		return
	}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("error while reading json file ", err.Error())
	}
}

func main() {

	displayTaskJson()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome to manager app")
	})

	http.HandleFunc("/task", addTaskToJson)
	http.HandleFunc("/gettask", fetchTaskAndServe)

	http.ListenAndServe(":8000", nil)

}

func fetchTaskAndServe(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonData)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorMessage{
			Msg: "can't send the json pelase try agin",
		})
	}

}

func addTaskToJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(ErrorMessage{
			Msg: "please use appropiate method",
		})
		return
	}

	var taskFromUser Task

	byteCount, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while reading body of add task to json ", err.Error())
	}
	err = json.Unmarshal(byteCount, &taskFromUser)
	if err != nil {
		fmt.Println("error while saving the request body to task form user")
	}

	result := checkId(taskFromUser)

	if result {
		json.NewEncoder(w).Encode(ErrorMessage{
			Msg: "please give another id",
		})
		return
	}

	jsonData = append(jsonData, taskFromUser)

	byteData, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("sorry we couldn't marshal the json data", err.Error())
	}

	err = os.WriteFile("data.json", byteData, 0644)
	if err != nil {
		fmt.Println("sorry we couldn't marshal the json data", err.Error())
	}

	fmt.Println("data successfully included on the data.json")

	json.NewEncoder(w).Encode(ErrorMessage{
		Msg: "successfully done ",
	})

}

func checkId(t Task) bool {
	id := t.Id
	if len(jsonData) == 0 {
		return false
	}
	found := false
	for i, v := range jsonData {
		fmt.Printf("the i  : %v  and value is %v \n", i, v.Id)
		if v.Id == id {
			found = true
		}
	}
	return found

}

type Task struct {
	Id    int    `json:"Id"`
	Title string `json:"Title"`
	Done  bool   `json:"Done"`
}

func displayTaskJson() {
	if len(jsonData) == 0 {
		fmt.Print("the task are empty please add the tasks")
		return
	}

	for i := 0; i < len(jsonData); i++ {
		fmt.Printf("%v .  %s  ->   %v \n", jsonData[i].Id, jsonData[i].Title, jsonData[i].Done)

	}
}

type ErrorMessage struct {
	Msg string `json:"msg"`
}

func writeToFile(taskData Task) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error while reading data", err.Error())
		return
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("error while reading json file ", err.Error())
	}

}
