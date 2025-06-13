package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server 3 handle func finished")

		fmt.Fprintf(w, "hello from server 3 port 8003")
	})

	fmt.Println("server started at port 8003")
	http.ListenAndServe(":8003", nil)
}
