package main

import (
	"fmt"

	util "github.com/manojbhatta500/auth/internal/utils"
)

func main() {
	// err := config.Load()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// config.ConfigInstance.ConfigPrinter()
	// database.ConnectToPostgres()
	// fmt.Println("program finished")

	output, err := util.ConvertToHash("manoj")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(output)

	matchOutput := util.VerifyPassword(output, "manoj1")
	if matchOutput != nil {
		fmt.Println("sorry password did not matched", matchOutput.Error())
	} else {
		fmt.Println("password matched")
	}

}
