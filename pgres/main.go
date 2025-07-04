package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/manojbhatta500/pgres/db"
	"github.com/manojbhatta500/pgres/routes"
)

func main() {

	databaseResult := db.ConnectToPostgress()
	if !databaseResult {
		fmt.Println("closing program......... ")
		os.Exit(1)
	}
	http.HandleFunc("/create", routes.CreateExpenses)
	http.HandleFunc("/update", routes.UpdateExpenses)
	http.HandleFunc("/delete", routes.DeleteExpenses)
	http.HandleFunc("/fetch", routes.GetExpenses)

	fmt.Println("server started at port 8000")
	port := ":8000"
	http.ListenAndServe(port, nil)
}

// func insertOneItem() {
// 	_, err := db.Db.Exec(db.Ctx, `
// 	insert into expenses values
// 	($1,$2,$3)
// 	`, 1, "food", 200)
// 	if err != nil {
// 		fmt.Println("cannot perform the database operration", err.Error())
// 	}
// }

// func deleteOne() {
// 	_, err := db.Db.Exec(db.Ctx, `
// 	delete from expenses where id = 1
// 	`)
// 	if err != nil {
// 		fmt.Println("sorry couldn't perfom the operation", err.Error())
// 		return
// 	}
// 	fmt.Println("successfully deleted ")
// }

// func queryOne() {
// 	rows, err := db.Db.Query(db.Ctx, `
// 	select * from expenses
// 	`)
// 	if err != nil {
// 		fmt.Println("sorry couldn't perform the query", err.Error())
// 	}

// 	var id int
// 	var task string
// 	var amount int

// 	for rows.Next() {
// 		err := rows.Scan(&id, &task, &amount)
// 		if err != nil {
// 			fmt.Println("error while scanning id", err.Error())
// 		}
// 		fmt.Printf("%v %v %v ", id, task, amount)
// 	}

// }

// func updateOne(
// 	id int,
// 	exp string,
// 	amt int,
// ) {
// 	_, err := db.Db.Exec(db.Ctx, `
// 	update expenses set id = $1, expensename= $2, amount = $3  where  id = 1
// 	`, id, exp, amt)
// 	if err != nil {
// 		fmt.Println("error updating the db")
// 	}

// 	fmt.Println("finished")
// }
