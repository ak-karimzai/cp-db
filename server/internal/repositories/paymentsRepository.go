package repositories

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/models"
)

type PaymentsRepository struct {
	dbHandler *sql.DB
}

func NewPaymentsRepository(dbHandler *sql.DB) *PaymentsRepository {
	return &PaymentsRepository{
		dbHandler: dbHandler,
	}
}

func (pr PaymentsRepository) CreatePayment(
	payment *models.Payment) (*models.Payment, *models.ResponseError) {
	query := `
		INSERT INTO payments(bill_id, user_id, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := pr.dbHandler.
		QueryRow(
			query,
			payment.BillId,
			payment.UserId,
			time.Now(),
		).
		Scan(
			&payment.ID,
		)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	return payment, nil
}

func (pr PaymentsRepository) GetAllPayments() (
	[]*models.Payment, *models.ResponseError) {
	query := `
			SELECT p.id, b.payment_amount, p.bill_id, p.user_id, p.created_at
			FROM payments p
			JOIN bills b
			ON p.bill_id = b.id
	`

	rows, err := pr.dbHandler.Query(query)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		payment := new(models.Payment)
		err := rows.Scan(
			&payment.ID,
			&payment.Amount,
			&payment.BillId,
			&payment.UserId,
			&payment.CreatedAt,
		)
		if err != nil {
			logger.GetLogger().Error(err)
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (pr PaymentsRepository) GetUserPayments(userId string) (
	[]*models.Payment, *models.ResponseError) {
	query := `
		SELECT p.id, b.payment_amount, p.bill_id, p.user_id, p.created_at
		FROM payments p
		JOIN bills b
		ON p.bill_id = b.id
		WHERE user_id = $1
	`
	rows, err := pr.dbHandler.Query(query, userId)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		payment := new(models.Payment)
		err := rows.Scan(
			&payment.ID,
			&payment.Amount,
			&payment.BillId,
			&payment.UserId,
			&payment.CreatedAt,
		)
		if err != nil {
			logger.GetLogger().Error(err)
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (pr PaymentsRepository) GetPayment(
	id string) (*models.Payment, *models.ResponseError) {
	query := `
		SELECT id, b.amount, bill_id, user_id, created_at
		FROM payments p
		JOIN bills b
		ON p.bill_id = b.id
		WHERE p.id = $1
	`
	payment := new(models.Payment)
	err := pr.dbHandler.
		QueryRow(query, id).
		Scan(
			&payment.ID,
			&payment.Amount,
			&payment.BillId,
			&payment.UserId,
			&payment.CreatedAt,
		)
	if err != nil && err != sql.ErrNoRows {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}

	if payment.ID == "" {
		return nil, &models.ResponseError{
			Message: "not found",
			Status:  http.StatusNotFound,
		}
	}
	return payment, nil
}
