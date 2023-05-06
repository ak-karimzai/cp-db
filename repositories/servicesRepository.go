package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
)

type ServicesRepository struct {
	dbHandler *sql.DB
}

func NewServicesRepository(
	dbHandler *sql.DB) *ServicesRepository {
	return &ServicesRepository{
		dbHandler: dbHandler,
	}
}

func (sr ServicesRepository) CreateService(
	service *models.Service) (*models.Service, *models.ResponseError) {
	query := `
		INSERT INTO services(name, description, m_amount, cost) 
		VALUES ($1, $2, $3, $4)`
	rows, err := sr.dbHandler.Query(query,
		service.Name,
		service.Description,
		service.MAmount,
		service.Cost)
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

	return &models.Service{
		ID:          id,
		Name:        service.Name,
		Description: service.Description,
		MAmount:     service.MAmount,
		Cost:        service.Cost,
	}, nil
}

func (sr ServicesRepository) UpdateService(
	service *models.Service) *models.ResponseError {
	query := `
		UPDATE services
		SET 
			name = $1,
			description = $2,
			m_amount = $3,
			cost = $4,
		WHERE id = $5`
	result, err := sr.dbHandler.Exec(query,
		service.Name,
		service.Description,
		service.MAmount,
		service.Cost)
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
			Message: "Service not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (sr ServicesRepository) GetService(
	id string) (*models.Service, *models.ResponseError) {
	query := `
		SELECT *
		FROM services
		WHERE id = $1
		`
	rows, err := sr.dbHandler.Query(query,
		id)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	service := new(models.Service)
	for rows.Next() {
		err := rows.Scan(&service.ID,
			&service.Name,
			&service.Description,
			&service.MAmount,
			&service.Cost)
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

	if service.ID == "" {
		return nil, &models.ResponseError{
			Message: "Service not found",
			Status:  http.StatusNotFound,
		}
	}
	return service, nil
}

func (sr ServicesRepository) GetAll() (
	[]*models.Service, *models.ResponseError) {
	query := `
		SELECT *
		FROM services
		`
	rows, err := sr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	services := make([]*models.Service, 0)
	for rows.Next() {
		service := new(models.Service)
		err := rows.Scan(&service.ID,
			&service.Name,
			&service.Description,
			&service.MAmount,
			&service.Cost)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		services = append(services, service)
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return services, nil
}
