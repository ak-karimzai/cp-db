package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
)

type UsersRepository struct {
	dbHandler *sql.DB
}

func NewUsersRepository(
	dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}

func (ur UsersRepository) Login(
	username, password string) (*models.User, *models.ResponseError) {
	query := `
		SELECT *
		FROM users
		WHERE username = $1
		AND user_password = crypt($2, user_password)`
	rows, err := ur.dbHandler.Query(
		query, username, password)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	user := new(models.User)
	for rows.Next() {
		err := rows.Scan(&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Password,
			&user.Role)
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
	return user, nil
}

func (ur UsersRepository) SignUp(user *models.User) (*models.User, *models.ResponseError) {
	query := `
		INSERT INTO users(first_name, last_name, username, user_password, user_role)
		VALUES ($1, $2, $3, crypt($4, gen_salt('bf')), $5) RETURNING id`
	rows, err := ur.dbHandler.Query(query,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Password,
		user.Role)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	var Id string
	for rows.Next() {
		err := rows.Scan(&Id)
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

	return &models.User{
		ID:        Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Password:  user.Password,
		Role:      user.Role,
	}, nil
}
