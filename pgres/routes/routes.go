package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/manojbhatta500/pgres/controller"
	"github.com/manojbhatta500/pgres/models"
)

func CreateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		fmt.Println("sorry GetExpense is  a post method")
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "this is a post request",
		})
		return
	}
	var inputModel models.Expense
	err := json.NewDecoder(r.Body).Decode(&inputModel)
	if err != nil {
		fmt.Println("error :", err.Error())
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: err.Error(),
		})
	}
	output, err := controller.SaveExpenseToPostGres(inputModel)
	if err != nil {
		fmt.Println("error :", err.Error())
	}
	json.NewEncoder(w).Encode(models.OperationStatus{
		Status:  output,
		Message: "successfully inserted data",
	})

}

func UpdateExpenses(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "this is put method",
		})
		return
	}
	w.Header().Add("Content-Type", "application/json")

	param := r.URL.Query()
	strid := param.Get("id")
	id, err := strconv.Atoi(strid)
	if err != nil || id < 1 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "please provide a correct id",
		})
		return
	}
	task := param.Get("task")
	stramt := param.Get("amount")
	amt, err := strconv.Atoi(stramt)
	if err != nil || amt < 1 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "please provide a correct amount",
		})
		return
	}
	res := controller.UpdateExpanseFromPostgres(id, task, amt)
	if res {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(models.OperationStatus{
			Status:  res,
			Message: "successfully updated.",
		})
		return
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(models.OperationStatus{
			Status:  res,
			Message: "successfully updated.",
		})
		return
	}

}

func DeleteExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		json.NewEncoder(w).Encode(
			models.ErrorMessage{
				Message: "you can only send delete request here",
			},
		)
		return
	}

	query := r.URL.Query()
	strId := query.Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0 {
		json.NewEncoder(w).Encode(
			models.ErrorMessage{
				Message: "please give appropriate id",
			},
		)
		return
	}

	result := controller.DeleteExpense(id)

	if result {
		json.NewEncoder(w).Encode(
			models.OperationStatus{
				Status:  result,
				Message: "successfully deleted the item",
			},
		)
		return
	} else {
		json.NewEncoder(w).Encode(
			models.OperationStatus{
				Status:  result,
				Message: "error deleting the item",
			},
		)
		return

	}

}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(
			models.ErrorMessage{
				Message: "you can only send get request here",
			},
		)
	}

	output, err := controller.GetExpensesFromPostgres()
	if err != nil {
		fmt.Println("the error is : ", err.Error())
		json.NewEncoder(w).Encode(
			models.ErrorMessage{
				Message: err.Error(),
			},
		)
	}

	json.NewEncoder(w).Encode(output)

}
