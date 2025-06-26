package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manojbhatta500/newsapp/database"
	"github.com/manojbhatta500/newsapp/middleware"
	"github.com/manojbhatta500/newsapp/models"
	"github.com/manojbhatta500/newsapp/utils"
)

func GetAllNews(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "Could not retrieve user ID",
		})
		return
	}

	fmt.Println("getAll news function running  userid : ", userID)

	var newsArticles []models.NewsArticleModel

	data, err := database.Db.Query(database.Ctx, "select * from newsarticle ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "sorry no users exists",
		})
		return
	}

	for data.Next() {
		var article models.NewsArticleModel
		err := data.Scan(
			&article.Id,
			&article.Title,
			&article.Body,
			&article.Category,
			&article.AuthorId,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  false,
				Message: "scanning issue for article",
			})
			return
		}
		newsArticles = append(newsArticles, article)

	}

	fmt.Println(newsArticles)

	w.WriteHeader(200)

	json.NewEncoder(w).Encode(&models.GetAllNewsModel{
		News: newsArticles,
	})

}

func SaveNews(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "login is post method",
		})
		return
	}
	fmt.Println("1")
	var loginBody models.LoginInputModel

	err := json.NewDecoder(r.Body).Decode(&loginBody)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "login and password is required",
		})
		return
	}
	fmt.Println("2")

	databaseresponse, err := database.Db.Query(database.Ctx, `select * from users where email=$1 `, loginBody.Email)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "user does not exits",
		})
		return
	}

	fmt.Println("3")

	var userData models.User
	for databaseresponse.Next() {
		err := databaseresponse.Scan(&userData.Id, &userData.UserName, &userData.Email, &userData.PasswordHash, &userData.IsAdmin)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  false,
				Message: "error while scanning the data from database response",
			})
			return
		}
	}
	fmt.Println("4")

	err = utils.VerifyPassword(userData.PasswordHash, loginBody.Password)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "please provide the correct password",
		})
		return
	}
	fmt.Println("6")

	// generate token and send it

	token, err := utils.GenerateToken(userData.Id, userData.Email)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "error while generating token",
		})
		return
	}

	var sendingModel models.LoginSuccessModel
	sendingModel.Status = true
	sendingModel.Token = token

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(sendingModel)
	fmt.Println("successfully logged in")

}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "signup is post method",
		})
		return
	}

	fmt.Println("1")
	var signupBody models.SignUpInputModel
	err := json.NewDecoder(r.Body).Decode(&signupBody)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "username,email,password are  required",
		})
		return
	}

	fmt.Println(" step 2")

	if len(signupBody.PasswordHash) < 8 || len(signupBody.Email) < 10 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "password must be of 8 character and email must be of 10",
		})
		return
	}

	fmt.Println(" step 3")

	rows, err := database.Db.Query(
		database.Ctx,
		"SELECT username FROM users WHERE username=$1",
		signupBody.UserName,
	)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "error while executing database command",
		})
		return
	}
	defer rows.Close()

	if rows.Next() {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "username already exists.please select another username",
		})
		return
	}

	hashPassword, err := utils.ConvertToHash(signupBody.PasswordHash)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "error while converting password to hash",
		})
		return
	}

	result, err := database.Db.Exec(
		database.Ctx,
		`INSERT INTO users (username, email, password,isadmin) VALUES ($1, $2, $3,$4)`,
		signupBody.UserName, signupBody.Email, hashPassword, false,
	)

	if err != nil {
		fmt.Println("error is ", err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  false,
			Message: "error while storing signup data to database",
		})
		return
	}

	fmt.Println(result.Insert(), "result insert printed")

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  true,
		Message: "successfully registered user.",
	})
}

//  learn about middleware so i can use are they valid
