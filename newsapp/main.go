package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/manojbhatta500/newsapp/database"
	"github.com/manojbhatta500/newsapp/middleware"
	"github.com/manojbhatta500/newsapp/routers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("something went wrong while loading env file")
	}

	res := database.ConnectToPostgress()
	if !res {
		fmt.Println("can't connect to databse")
		os.Exit(1)
	}

	// handlers

	http.HandleFunc("/login", routers.Login)
	http.HandleFunc("/signup", routers.Signup)

	http.HandleFunc("/save-news", routers.SaveNews)

	http.HandleFunc("/all-news", middleware.Logger(middleware.CheckToken(middleware.CheckPostOnlyMethod(routers.GetAllNews))))

	fmt.Println("starting server at port ", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)

}
