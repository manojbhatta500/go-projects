package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	getWorkingDirectory()
	fmt.Println("starting server in port :8000")
	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "downloads/download.html")
	})
	http.HandleFunc("/sendData", downloadContent)
	http.ListenAndServe(":8000", nil)

}
func downloadContent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("this is a post method please post it")
		return
	}
	var inputBody inputJsonBody

	err := json.NewDecoder(r.Body).Decode(&inputBody)

	if err != nil {
		fmt.Println("error while reading body 2", err.Error())
	}
	if inputBody.Content == "" {
		fmt.Println("don't send an empty string")
		return
	}
	fetchFiles(inputBody.Content)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(inputJsonBody{
		Content: "successfully fetched data",
	})

	if err != nil {
		fmt.Println("error while sending data", err.Error())
	}
}

type inputJsonBody struct {
	Content string `json:"content"`
}

func getWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("the err is ", err.Error())
	}
	fmt.Println("the working directory is ", dir)

}

func fetchFiles(endpoint string) {
	fmt.Println("fetch files function called")

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("the error is ", err.Error())
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("the error we are getting  is", err.Error())
	}

	fmt.Println(string(data))

	createFileAndRun(string(data), "download.html")

}

func createFileAndRun(dataBody string, filename string) {
	err := os.MkdirAll("downloads", os.ModePerm)
	if err != nil {
		fmt.Println("error creating folder:", err.Error())
		return
	}
	var file *os.File
	fullFileName := "downloads/" + filename
	if _, err := os.Stat(fullFileName); os.IsNotExist(err) {
		var err error
		file, err = os.Create(fullFileName)
		if err != nil {
			fmt.Println("error while creating file", err.Error())
		}
		output, err := file.WriteString(dataBody)
		if err != nil {
			fmt.Println("error while wirting string in file", err.Error())
		}
		fmt.Println("the output of file is bytes ", output)

	} else {
		fmt.Println("file is already present there")
		file, err = os.OpenFile(fullFileName, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("error while opening existing file", err.Error())
			return
		}
		output, err := file.WriteString(dataBody)
		if err != nil {
			fmt.Println("error while writing string in file", err.Error())
		}
		fmt.Println("the output of file is bytes ", output)

	}

}
