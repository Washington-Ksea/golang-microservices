package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Washington-Ksea/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Fade", LastName: "Leon", Email: "myemail@gmail.com"},
	}

	UserDao usersDaoInterface
)

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
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
