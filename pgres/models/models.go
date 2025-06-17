package models

type Expense struct {
	Id          int    `json:"id"`
	ExpenseName string `json:"expensename"`
	Amount      int    `json:"amount"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type OperationStatus struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
