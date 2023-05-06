package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
)

type BillsRepository struct {
	dbHandler *sql.DB
}

func NewBillsRepository(
	dbHandler *sql.DB) *BillsRepository {
	return &BillsRepository{
		dbHandler: dbHandler,
	}
}

func (br BillsRepository) CreateBill(
	bill *models.Bill) (*models.Bill, *models.ResponseError) {
	query := `
		INSERT INTO bills("from", until, spend_amount, apartment_id, service_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, payment_amount`
	rows, err := br.dbHandler.Query(query,
		bill.From,
		bill.Until,
		bill.SpendAmount,
		bill.ApartmentId,
		bill.ServiceId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	var id string
	var paymentAmount float64
	for rows.Next() {
		err := rows.Scan(&id, &paymentAmount)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.Bill{
		ID:            id,
		From:          bill.From,
		Until:         bill.Until,
		SpendAmount:   bill.SpendAmount,
		PaymentAmount: paymentAmount,
		ApartmentId:   bill.ApartmentId,
		ServiceId:     bill.ServiceId,
	}, nil
}

func (br BillsRepository) GetBill(
	id string) (*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills
		WHERE id = $1`
	rows, err := br.dbHandler.Query(query, id)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	bill := new(models.Bill)
	for rows.Next() {
		err := rows.Scan(&bill.ID,
			&bill.From,
			&bill.Until,
			&bill.SpendAmount,
			&bill.PaymentAmount,
			&bill.Paid,
			&bill.ApartmentId,
			&bill.ServiceId)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return bill, nil
}

func (br BillsRepository) GetAll() (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills`
	rows, err := br.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func (br BillsRepository) GetAllPaid() (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills
		WHERE paid = true`
	rows, err := br.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func (br BillsRepository) GetAllUnpaid() (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills
		WHERE paid = false`
	rows, err := br.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func parseBillRows(rows *sql.Rows) (
	[]*models.Bill, *models.ResponseError) {
	bills := make([]*models.Bill, 0)
	for rows.Next() {
		bill := new(models.Bill)
		err := rows.Scan(&bill.ID,
			&bill.From,
			&bill.Until,
			&bill.SpendAmount,
			&bill.PaymentAmount,
			&bill.Paid,
			&bill.ApartmentId,
			&bill.ServiceId)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		bills = append(bills, bill)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return bills, nil
}
