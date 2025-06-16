package main

import (
	"fmt"

	"github.com/manojbhatta500/pgres/db"
)

func main() {
	fmt.Println("calling db go file ")
	insertOneItem()
	// deleteOne()
	queryOne()

}

func insertOneItem() {
	_, err := db.Db.Exec(db.Ctx, `
	insert into expenses values
	($1,$2,$3)
	`, 1, "food", 200)
	if err != nil {
		fmt.Println("cannot perform the database operration", err.Error())
	}
}

func deleteOne() {
	_, err := db.Db.Exec(db.Ctx, `
	delete from expenses where id = 1
	`)
	if err != nil {
		fmt.Println("sorry couldn't perfom the operation", err.Error())
		return
	}
	fmt.Println("successfully deleted ")
}

func queryOne() {
	rows, err := db.Db.Query(db.Ctx, `
	select * from expenses
	`)
	if err != nil {
		fmt.Println("sorry couldn't perform the query", err.Error())
	}

	var id int
	var task string
	var amount int

	for rows.Next() {
		err := rows.Scan(&id, &task, &amount)
		if err != nil {
			fmt.Println("error while scanning id", err.Error())
		}
		fmt.Printf("%v %v %v ", id, task, amount)
	}

}
