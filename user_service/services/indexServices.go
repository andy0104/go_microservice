package services

import (
	"go_microservice/user_service/dto"
	"go_microservice/user_service/repository"

	"go.uber.org/zap"
)

type IndexServices struct {
	UserService interface {
		UserSignup(dto.UserSignupRequest) (string, error)
	}
}

func NewServices(repo repository.Repository, logger *zap.SugaredLogger) IndexServices {
	return IndexServices{
		UserService: &UserService{repo: repo, logger: logger},
	}
}
