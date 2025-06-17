package controller

import (
	"fmt"

	"github.com/manojbhatta500/pgres/db"
	"github.com/manojbhatta500/pgres/models"
)

func SaveExpenseToPostGres(exp models.Expense) (bool, error) {
	_, err := db.Db.Exec(db.Ctx, `
	insert into expenses values
	($1,$2,$3)
	`, exp.Id, exp.ExpenseName, exp.Amount)
	if err != nil {
		fmt.Println("error while excuting postgres command", err.Error())
		return false, err
	}
	fmt.Println("data successfully inserted")
	return true, nil
}

func GetExpensesFromPostgres() (data []models.Expense, err error) {
	var output []models.Expense
	rows, err := db.Db.Query(db.Ctx, `
	select * from expenses;
	`)
	if err != nil {
		fmt.Println("error", err.Error())
		return nil, err
	}
	var id int
	var task string
	var amount int
	for rows.Next() {
		err := rows.Scan(&id, &task, &amount)
		if err != nil {
			fmt.Println("error while scanning id", err.Error())
			return nil, err
		}
		output = append(output, models.Expense{
			Id:          id,
			ExpenseName: task,
			Amount:      amount,
		})
		fmt.Println("added 1 item to output")
	}
	fmt.Println("the length of output is ", len(output))
	return output, nil

}

func UpdateExpanseFromPostgres(id int, name string, amt int) bool {

	output, err := db.Db.Exec(db.Ctx, `
	update expenses set expensename = $1 , amount = $2 where id = $3
	`, name, amt, id)

	if err != nil {
		fmt.Println("the  error ", err.Error())
		return false
	}

	fmt.Println("update task completed", output.Update())
	return true

}

func DeleteExpense(id int) bool {
	res, err := db.Db.Exec(db.Ctx, `delete from expenses  where id = $1`, id)
	if err != nil {
		fmt.Println("error while deleting from the table", err.Error())
		return false
	}
	fmt.Println("successfully delete from database ", res.Delete())
	return true
}
