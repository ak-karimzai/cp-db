package services

import (
	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/repositories"
)

type BillsServices struct {
	billsRepository *repositories.BillsRepository
}

func NewBillsServices(
	billsRepository *repositories.BillsRepository) *BillsServices {
	return &BillsServices{
		billsRepository: billsRepository,
	}
}

func (bs *BillsServices) CreateBill(bill *models.Bill) (
	*models.Bill, *models.ResponseError) {
	if responseErr := validBill(bill); responseErr != nil {
		return nil, responseErr
	}

	response, responseErr := bs.billsRepository.
		CreateBill(bill)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (bs *BillsServices) GetBill(id string) (
	*models.Bill, *models.ResponseError) {
	if responseErr := validBillId(id); responseErr != nil {
		return nil, responseErr
	}

	response, responseErr := bs.billsRepository.
		GetBill(id)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (bs *BillsServices) GetUserBills(userId string) (
	[]*models.Bill, *models.ResponseError) {
	response, responseErr := bs.billsRepository.
		GetUserBills(userId)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (bs *BillsServices) GetAll() (
	[]*models.Bill, *models.ResponseError) {
	response, responseErr := bs.billsRepository.
		GetAll()
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (bs *BillsServices) GetAllPaid() (
	[]*models.Bill, *models.ResponseError) {
	response, responseErr := bs.billsRepository.
		GetAllPaid()
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func (bs *BillsServices) GetAllUnpaid() (
	[]*models.Bill, *models.ResponseError) {
	response, responseErr := bs.billsRepository.
		GetAllUnpaid()
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func validBillId(billId string) *models.ResponseError {
	if billId == "" {
		return &models.ResponseError{
			Message: "Invalid bill ID",
		}
	}
	return nil
}

func validBill(
	bill *models.Bill) *models.ResponseError {
	return nil
}
