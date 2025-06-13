package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	// performGetReq()
	performPostReq()

}

func performGetReq() {

	endpoint := "http://localhost:8001/hello"

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("the err is ", err.Error())
	}
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	fmt.Println("the response form the server is")
	// fmt.Println(string(response))
	var responseString strings.Builder
	byteCount, err := responseString.Write(response)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	fmt.Println("the byte count is ", byteCount)
	fmt.Println("the actual response is ", responseString.String())

}

func performPostReq() {
	reqbody := strings.NewReader(`
	{
		"msg" : "hello we are learning golang"
	}
	`)
	endpoint := "http://localhost:8001/hello"

	res, err := http.Post(endpoint, "application/json", reqbody)

	if err != nil {
		fmt.Println("the error is ", err.Error())
	}
	low, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("the error is ", err.Error())
	}
	fmt.Println("count of byte is ", low)
	var responseString strings.Builder

	count, err := responseString.Write(low)
	if err != nil {
		fmt.Println("the actual error is", err.Error())
	}
	fmt.Println("the count is ", count)
	fmt.Println("the actual response is ", responseString.String())

}
