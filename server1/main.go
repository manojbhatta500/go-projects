package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server 1 handle func finished")
		fmt.Fprintf(w, "hello from server 1 port 8001")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		m := Message{
			Msg: "hello form server1",
		}
		json.NewEncoder(w).Encode(m)
	})

	fmt.Println("server started at port 8001")
	http.ListenAndServe(":8001", nil)
}

type Message struct {
	Msg string `json:"msg"`
}
