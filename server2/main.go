package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server 2 handle func finished")

		fmt.Fprintf(w, "hello from server 2 port 8002")
	})

	fmt.Println("server started at port 8002")
	http.ListenAndServe(":8002", nil)
}
