package models

type KmcRegisterPostApiBody struct {
	Usertype string `json:"usertype"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type Target struct {
	T int `json:"t"`
}

type LoginRksPostBody struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
