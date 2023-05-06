package services

import (
	"net/http"

	"github.com/ak-karimzai/cp-db/jwt"
	"github.com/ak-karimzai/cp-db/models"
	"github.com/ak-karimzai/cp-db/repositories"
)

type UsersServices struct {
	userRepository *repositories.UsersRepository
}

func NewUsersServices(userRepository *repositories.UsersRepository) *UsersServices {
	return &UsersServices{
		userRepository: userRepository,
	}
}

func (us UsersServices) Login(username,
	password string) (string, *models.ResponseError) {
	if username == "" || password == "" {
		return "", &models.ResponseError{
			Message: "Incorect username or password",
			Status:  http.StatusBadRequest,
		}
	}
	response, responseErr := us.userRepository.
		Login(username, password)
	if responseErr != nil {
		return "", responseErr
	}

	token, responseErr := jwt.GenerateToken(response.ID,
		response.Role)
	if responseErr != nil {
		return "", responseErr
	}
	return token, nil
}

func (us UsersServices) SignUp(
	user *models.User) (*models.User, *models.ResponseError) {
	response, responseErr := us.userRepository.SignUp(user)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}
