package services

import (
	"net/http"
	"testing"

	"github.com/Washington-Ksea/golang-microservices/mvc/utils"

	"github.com/Washington-Ksea/golang-microservices/mvc/domain"
	"github.com/stretchr/testify/assert"
)

var (
	userDoaMock usersDaoMock

	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNotError(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{ID: 123}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}
