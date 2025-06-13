package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rental.com/routes" // Remove underscore
)

func main() {
	handleRoutes()
	startServer()
}

func handleRoutes() {
	http.HandleFunc("/", routes.HandleRoot)
	http.HandleFunc("/kmc", routes.HandleKmc)
	http.HandleFunc("/leapfrog", routes.BruteForceLeapFrog)
}

func startServer() {
	fmt.Println("server starting at 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
