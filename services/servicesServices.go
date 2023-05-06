package services

import (
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
	"github.com/ak-karimzai/cp-db/repositories"
)

type ServicesServices struct {
	serviceRepository *repositories.ServicesRepository
}

func NewServicesServices(
	serviceRepository *repositories.ServicesRepository) *ServicesServices {
	return &ServicesServices{
		serviceRepository: serviceRepository,
	}
}

func (ss ServicesServices) CreateService(
	service *models.Service) (*models.Service, *models.ResponseError) {
	if responseErr := validService(service); responseErr != nil {
		return nil, responseErr
	}

	response, responseErr := ss.serviceRepository.
		CreateService(service)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (ss ServicesServices) UpdateService(
	service *models.Service) *models.ResponseError {
	if service.ID == "" {
		return &models.ResponseError{
			Message: "Incorect service id",
			Status:  http.StatusBadRequest,
		}
	}
	if responseErr := validService(service); responseErr != nil {
		return responseErr
	}

	responseErr := ss.serviceRepository.
		UpdateService(service)
	if responseErr != nil {
		return responseErr
	}
	return nil
}

func (ss ServicesServices) GetService(id string) (
	*models.Service, *models.ResponseError) {
	response, responseErr := ss.serviceRepository.
		GetService(id)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (ss ServicesServices) GetAllServices() (
	[]*models.Service, *models.ResponseError) {
	response, responseErr := ss.serviceRepository.
		GetAll()
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func validService(service *models.Service) *models.ResponseError {
	return nil
}
