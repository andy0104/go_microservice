package services

import (
	"go_microservice/user_service/dto"
	"go_microservice/user_service/repository"
	"go_microservice/user_service/utility/hasher"
	"go_microservice/user_service/utility/responses"

	"go.uber.org/zap"
)

type UserService struct {
	repo   repository.Repository
	logger *zap.SugaredLogger
}

func (us *UserService) UserSignup(payload dto.UserSignupRequest) (string, error) {
	// check if email is already used
	user, err := us.repo.User.GetUserByEmailID(payload.Email)
	if err != nil {
		return "", err
	}

	us.logger.Infow("repo.User.GetUserByEmailID", "user", user)

	if user != nil {
		return "", responses.ErrEmailInUseServer
	}

	// hash the password
	passHash, err := hasher.HashText(payload.Password)
	if err != nil {
		return "", err
	}
	payload.Password = string(passHash)

	// create a new user
	rowsAffected, err := us.repo.User.Create(payload)
	if err != nil && rowsAffected == 0 {
		return "", err
	}

	us.logger.Infow("repo.User.Create", "rowsAffected", rowsAffected)

	return "user signup", nil
}
