package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/manojbhatta500/newsapp/models"
	"github.com/manojbhatta500/newsapp/utils"
)

func CheckPostOnlyMethod(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  false,
				Message: "post method only allowed",
			})
			return
		}
		next(w, r)
		// ctx := context.WithValue(r.Context(), "userID", userID)
		// ctx = context.WithValue(ctx, "email", email)
		// next(w, r.WithContext(ctx))
		// here we can aceess it
		// userID, ok := r.Context().Value("userID").(string)
		// if !ok {
		//     // Handle case where userID is not set or not a string
		// }
		// email, ok := r.Context().Value("email").(string)
		// if !ok {
		//     // Handle case where email is not set or not a string
		// }
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logFile := "log.txt"
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening log file:", err)
		} else {
			logEntry := fmt.Sprintf("Method: %s URL: %s time: %s\n", r.Method, r.URL.String(), time.Now().Format(time.RFC3339))
			file.WriteString(logEntry)
			file.Close()
		}
		next(w, r)
	}
}

type contextKey string

const (
	UserIDKey contextKey = "userid"
	EmailKey  contextKey = "email"
)

func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")

		if authHeader == " " {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  false,
				Message: "please provide authorization header",
			})
			return
		}

		partsOfAuthHeader := strings.Split(authHeader, " ")

		if len(partsOfAuthHeader) != 2 || partsOfAuthHeader[0] != "Bearer" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  false,
				Message: "Invalid authorization format. Use Bearer <token>",
			})
			return
		}

		token := partsOfAuthHeader[1]

		data, err := utils.ValidateToken(token)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  false,
				Message: "invalid token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, data.UserID)
		ctx = context.WithValue(ctx, EmailKey, data.Email)

		next(w, r.WithContext(ctx))
	}
}
