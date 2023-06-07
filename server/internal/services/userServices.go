package services

import (
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/jwt"
	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/repositories"
)

type UsersServices struct {
	userRepository *repositories.UsersRepository
}

func NewUsersServices(userRepository *repositories.UsersRepository) *UsersServices {
	return &UsersServices{
		userRepository: userRepository,
	}
}

func (uc *UsersServices) GetUser(userId string) (
	*models.User, *models.ResponseError) {
	user, err := uc.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UsersServices) GetAll() (
	[]*models.User, *models.ResponseError) {
	users, err := us.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us UsersServices) Login(username,
	password string) (*models.User, string, *models.ResponseError) {
	if username == "" || password == "" {
		return nil, "", &models.ResponseError{
			Message: "Incorect username or password",
			Status:  http.StatusBadRequest,
		}
	}
	user, responseErr := us.userRepository.
		Login(username, password)
	if responseErr != nil {
		return nil, "", responseErr
	}

	token, responseErr := jwt.GenerateToken(user.ID,
		user.Role)
	if responseErr != nil {
		return nil, "", responseErr
	}
	return user, token, nil
}

func (us UsersServices) SignUp(
	user *models.User) (*models.User, *models.ResponseError) {
	response, responseErr := us.userRepository.SignUp(user)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}
