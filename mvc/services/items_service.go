package services

import (
	"net/http"

	"github.com/Washington-Ksea/golang-microservices/mvc/domain"
	"github.com/Washington-Ksea/golang-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
