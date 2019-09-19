package services

import (
	"strings"

	"github.com/Washington-Ksea/golang-microservices/src/api/config"

	"github.com/Washington-Ksea/golang-microservices/src/api/domain/github"
	"github.com/Washington-Ksea/golang-microservices/src/api/domain/repositories"
	"github.com/Washington-Ksea/golang-microservices/src/api/providers/github_provider"
	"github.com/Washington-Ksea/golang-microservices/src/api/utils/errors"
)

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

type reposService struct{}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
