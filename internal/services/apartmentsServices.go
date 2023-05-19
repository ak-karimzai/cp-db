package services

import (
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/repositories"
)

type ApartmentsServices struct {
	apartmentsRepository *repositories.ApartmentsRepository
}

func NewApartmentsService(
	apartmentsRepository *repositories.ApartmentsRepository) *ApartmentsServices {
	return &ApartmentsServices{
		apartmentsRepository: apartmentsRepository,
	}
}

func (as ApartmentsServices) CreateApartment(
	apartment *models.Apartment) (*models.Apartment,
	*models.ResponseError) {
	if responseErr := validApartment(
		apartment); responseErr != nil {
		return nil, responseErr
	}

	response, responseErr := as.
		apartmentsRepository.CreateApartment(apartment)
	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}

func (as ApartmentsServices) UpdateApartment(
	apartment *models.Apartment) *models.ResponseError {
	if responseErr := validId(
		apartment.ID); responseErr != nil {
		return responseErr
	}

	if responseErr := validApartment(
		apartment); responseErr != nil {
		return responseErr
	}

	responseErr := as.
		apartmentsRepository.UpdateApartment(apartment)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

func (as ApartmentsServices) GetApartment(
	id string) (*models.Apartment, *models.ResponseError) {
	if responseErr := validId(
		id); responseErr != nil {
		return nil, responseErr
	}

	response, responseErr := as.
		apartmentsRepository.GetApartment(id)
	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}

func (as ApartmentsServices) GetUserAparments(
	userId string) ([]*models.Apartment, *models.ResponseError) {
	if userId == "" {
		return nil, &models.ResponseError{
			Message: "Invalid user id",
			Status:  http.StatusBadRequest,
		}
	}
	response, responseErr := as.
		apartmentsRepository.GetUserAparments(userId)
	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}

func (as ApartmentsServices) GetAll() (
	[]*models.Apartment, *models.ResponseError) {
	response, responseErr := as.
		apartmentsRepository.GetAll()
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func validApartment(
	apartment *models.Apartment) *models.ResponseError {
	return nil
}

func validId(id string) *models.ResponseError {
	if id == "" {
		return &models.ResponseError{
			Message: "Invalid apartment id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
