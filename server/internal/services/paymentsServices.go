package services

import (
	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/repositories"
)

type PaymentsServices struct {
	paymentsRepository *repositories.PaymentsRepository
}

func NewPaymentsServices(
	paymentsRepository *repositories.PaymentsRepository) *PaymentsServices {
	return &PaymentsServices{
		paymentsRepository: paymentsRepository,
	}
}

func (ps PaymentsServices) CreatePayment(
	payment *models.Payment) (*models.Payment, *models.ResponseError) {
	response, responseErr := ps.paymentsRepository.
		CreatePayment(payment)
	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}

func (ps PaymentsServices) GetAllPayments() (
	[]*models.Payment, *models.ResponseError) {
	response, responseErr := ps.paymentsRepository.
		GetAllPayments()
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (ps PaymentsServices) GetPayment(
	id string) (*models.Payment, *models.ResponseError) {
	response, responseErr := ps.paymentsRepository.
		GetPayment(id)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (ps PaymentsServices) GetUserPayments(
	userId string) ([]*models.Payment, *models.ResponseError) {
	response, responseErr := ps.paymentsRepository.
		GetUserPayments(userId)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}
