package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// function to hash password

func ConvertToHash(input string) (string, error) {
	output, err := bcrypt.GenerateFromPassword([]byte(input), 12)
	if err != nil {
		fmt.Println("the error is ", err.Error())
		return "", err
	}
	return string(output), nil
}

func VerifyPassword(input string, password string) error {
	result := bcrypt.CompareHashAndPassword([]byte(input), []byte(password))
	if result != nil {
		return result
	} else {
		return nil
	}
}
