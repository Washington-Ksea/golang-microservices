package services

import (
	"github.com/Washington-Ksea/golang-microservices/mvc/domain"
	"github.com/Washington-Ksea/golang-microservices/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
