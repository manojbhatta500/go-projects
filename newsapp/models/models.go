package models

type NewsArticleModel struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category string `json:"category"`
	AuthorId int    `json:"authorid"`
}

type User struct {
	Id           int    `json:"id"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	IsAdmin      bool   `json:"isadmin"`
}

type LoginInputModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInputModel struct {
	UserName     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type LoginSuccessModel struct {
	Status bool   `json:"status"`
	Token  string `json:"Token"`
}

type GetAllNewsModel struct {
	News []NewsArticleModel `json:"news"`
}
