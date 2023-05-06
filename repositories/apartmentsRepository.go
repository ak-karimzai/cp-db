package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
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
	rows, err := ar.dbHandler.Query(query,
		apartment.Size,
		apartment.RoomNumbers,
		apartment.UserId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	var id string
	for rows.Next() {
		err := rows.Scan(&id)
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
	return &models.Apartment{
		ID:          id,
		Size:        apartment.Size,
		RoomNumbers: apartment.RoomNumbers,
		UserId:      apartment.UserId,
	}, nil
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
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "Apartment not found",
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
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
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
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		apartments = append(apartments, apartment)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
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
	rows, err := ar.dbHandler.Query(query, id)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	apartment := new(models.Apartment)
	for rows.Next() {
		err := rows.Scan(&apartment.ID,
			&apartment.Size,
			&apartment.RoomNumbers,
			&apartment.UserId)
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
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
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
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		apartments = append(apartments, apartment)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return apartments, nil
}
