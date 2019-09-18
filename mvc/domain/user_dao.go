package domain

import (
	"fmt"
	"net/http"

	"github.com/Washington-Ksea/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Fade", LastName: "Leon", Email: "myemail@gmail.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		fmt.Printf("Success get user: %v", user)
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
