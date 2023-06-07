package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/models"
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

func (br *BillsRepository) UpdateBill(
	bill *models.Bill) *models.ResponseError {
	query :=
		`
		UPDATE bills
		SET spend_amount = $1
		WHERE id = $2
	`
	_, err := br.dbHandler.Exec(
		query,
		bill.SpendAmount,
		bill.ID,
	)
	if err != nil {
		logger.GetLogger().Error(err)
		return errorJSON(
			err, http.StatusInternalServerError)
	}
	return nil
}

func (br *BillsRepository) GetApartmentBills(
	aprId string) ([]*models.Bill, *models.ResponseError) {
	query := `
		SELECT id, "from", until, spend_amount, payment_amount, paid, apartment_id, service_id, created_at, updated_at
		FROM bills
		WHERE apartment_id = $1
	`
	rows, err := br.dbHandler.Query(query, aprId)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func (br *BillsRepository) CreateBill(
	bill *models.Bill) (*models.Bill, *models.ResponseError) {
	query := `
		INSERT INTO bills("from", until, spend_amount, apartment_id, service_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, payment_amount
	`
	var id string
	var paymentAmount float64
	err := br.dbHandler.QueryRow(query,
		bill.From,
		bill.Until,
		bill.SpendAmount,
		bill.ApartmentId,
		bill.ServiceId).
		Scan(
			&id,
			&paymentAmount,
		)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
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

func (br *BillsRepository) GetBill(
	id string) (*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills
		WHERE id = $1`

	bill := new(models.Bill)
	err := br.dbHandler.
		QueryRow(query, id).
		Scan(
			&bill.ID,
			&bill.From,
			&bill.Until,
			&bill.SpendAmount,
			&bill.PaymentAmount,
			&bill.Paid,
			&bill.ApartmentId,
			&bill.ServiceId,
			&bill.CreatedAt,
			&bill.UpdatedAt,
		)
	if err != nil && err != sql.ErrNoRows {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}

	if bill.ID == "" {
		return nil, &models.ResponseError{
			Message: "not found",
			Status:  http.StatusNotFound,
		}
	}
	return bill, nil
}

func (br *BillsRepository) GetUserBills(userId string) (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT b.* 
		FROM bills b
		JOIN payments p ON b.id = p.bill_id
		WHERE p.user_id = $1
	`
	rows, err := br.dbHandler.Query(query, userId)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func (br *BillsRepository) GetAll() (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills`
	rows, err := br.dbHandler.Query(query)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func (br *BillsRepository) GetAllPaid() (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills
		WHERE paid = true`
	rows, err := br.dbHandler.Query(query)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	return parseBillRows(rows)
}

func (br *BillsRepository) GetAllUnpaid() (
	[]*models.Bill, *models.ResponseError) {
	query := `
		SELECT *
		FROM bills
		WHERE paid = false`
	rows, err := br.dbHandler.Query(query)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
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
			&bill.ServiceId,
			&bill.CreatedAt,
			&bill.UpdatedAt)
		if err != nil {
			logger.GetLogger().Error(err)
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}
		bills = append(bills, bill)
	}

	if rows.Err() != nil {
		logger.GetLogger().Error(rows.Err())
		return nil, errorJSON(
			rows.Err(), http.StatusInternalServerError)
	}

	return bills, nil
}
