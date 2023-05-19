package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/models"
)

type ApartmentsRepository struct {
	dbHandler *sql.DB
}

func NewApartmentsRepository(
	dbHandler *sql.DB) *ApartmentsRepository {
	return &ApartmentsRepository{
		dbHandler: dbHandler,
	}
}

func (ar ApartmentsRepository) CreateApartment(
	apartment *models.Apartment) (*models.Apartment, *models.ResponseError) {
	query := `
		INSERT INTO apartments(size, room_numbers, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	err := ar.dbHandler.QueryRow(query,
		apartment.Size,
		apartment.RoomNumbers,
		apartment.UserId).
		Scan(
			&apartment.ID,
		)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}

	return apartment, nil
}

func (ar ApartmentsRepository) UpdateApartment(
	apartment *models.Apartment) *models.ResponseError {
	query := `
		UPDATE apartments
		SET
			size = $1,
			room_numbers = $2,
			user_id = $3
		WHERE id = $4`
	result, err := ar.dbHandler.Exec(query,
		apartment.Size,
		apartment.RoomNumbers,
		apartment.UserId,
		apartment.ID)
	if err != nil {
		logger.GetLogger().Error(err)
		return errorJSON(
			err, http.StatusInternalServerError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.GetLogger().Error(err)
		return errorJSON(
			err, http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (ar ApartmentsRepository) GetAll() (
	[]*models.Apartment, *models.ResponseError) {
	query := `
		SELECT *
		FROM apartments`
	rows, err := ar.dbHandler.Query(query)
	if err != nil {
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	apartments := make([]*models.Apartment, 0)
	for rows.Next() {
		apartment := new(models.Apartment)
		err := rows.Scan(&apartment.ID,
			&apartment.Size,
			&apartment.RoomNumbers,
			&apartment.UserId)
		if err != nil {
			logger.GetLogger().Error(err)
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}

		apartments = append(apartments, apartment)
	}

	if rows.Err() != nil {
		logger.GetLogger().Error(rows.Err())
		return nil, errorJSON(
			rows.Err(), http.StatusInternalServerError)
	}

	return apartments, nil
}

func (ar ApartmentsRepository) GetApartment(
	id string) (*models.Apartment, *models.ResponseError) {
	query := `
		SELECT *
		FROM apartments
		WHERE id = $1
	`
	apartment := new(models.Apartment)
	err := ar.dbHandler.
		QueryRow(query, id).
		Scan(
			&apartment.ID,
			&apartment.Size,
			&apartment.RoomNumbers,
			&apartment.UserId,
		)
	if err != nil && err != sql.ErrNoRows {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}

	if apartment.ID == "" {
		return nil, &models.ResponseError{
			Message: "not found",
			Status:  http.StatusNotFound,
		}
	}

	return apartment, nil
}

func (ar ApartmentsRepository) GetUserAparments(
	userId string) ([]*models.Apartment, *models.ResponseError) {
	query := `
		SELECT *
		FROM apartments
		WHERE user_id = $1
	`
	rows, err := ar.dbHandler.Query(query, userId)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	apartments := make([]*models.Apartment, 0)
	for rows.Next() {
		apartment := new(models.Apartment)
		err := rows.Scan(&apartment.ID,
			&apartment.Size,
			&apartment.RoomNumbers,
			&apartment.UserId)
		if err != nil {
			logger.GetLogger().Error(err)
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}
		apartments = append(apartments, apartment)
	}

	if rows.Err() != nil {
		logger.GetLogger().Error(rows.Err())
		return nil, errorJSON(
			rows.Err(), http.StatusInternalServerError)
	}

	return apartments, nil
}
