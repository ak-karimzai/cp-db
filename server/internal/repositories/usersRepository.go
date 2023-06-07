package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/models"
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

func (ur *UsersRepository) GetUser(userId string) (
	*models.User, *models.ResponseError) {
	query := `
		SELECT id, first_name, last_name, username, user_role
		FROM users
		WHERE id = $1`

	var user models.User
	err := ur.dbHandler.QueryRow(query, userId).
		Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Role,
		)
	if err != nil {
		return nil, errorJSON(err)
	}
	return &user, nil
}

func (ur *UsersRepository) GetAll() (
	[]*models.User, *models.ResponseError) {
	query := `
		SELECT id, first_name, last_name, username, user_role
		FROM users
		ORDER BY user_role`

	rows, err := ur.dbHandler.Query(query)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user = new(models.User)
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Role,
		)
		if err != nil {
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur UsersRepository) Login(
	username, password string) (*models.User, *models.ResponseError) {
	query := `
		SELECT id, first_name, last_name, username, user_role
		FROM users
		WHERE username = $1
		AND user_password = $2`

	var user models.User
	err := ur.dbHandler.QueryRow(
		query, username, password).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Role,
	)
	if err != nil {
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}

	return &user, nil
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
		logger.GetLogger().Error(err)
		return nil, errorJSON(
			err, http.StatusInternalServerError)
	}
	defer rows.Close()

	var Id string
	for rows.Next() {
		err := rows.Scan(&Id)
		if err != nil {
			logger.GetLogger().Error(err)
			return nil, errorJSON(
				err, http.StatusInternalServerError)
		}
	}

	if rows.Err() != nil {
		logger.GetLogger().Error(rows.Err())
		return nil, errorJSON(
			rows.Err(), http.StatusInternalServerError)
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
