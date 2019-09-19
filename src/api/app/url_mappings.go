package app

import (
	"github.com/Washington-Ksea/golang-microservices/src/api/controllers/polo"
	"github.com/Washington-Ksea/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
