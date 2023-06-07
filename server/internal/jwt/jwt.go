package jwt

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var jwtSecretKey []byte

const (
	USER_ID    = "ID"
	USER_ROLE  = "ROLE"
	CREATED_AT = "CREATED_AT"
	EXPIRED_AT = "EXPIRED_AT"
)

func newUserInfo(userId, userRole string) *models.User {
	return &models.User{
		ID:   userId,
		Role: userRole,
	}
}

func InitJwtSerectKey(config *viper.Viper) {
	key := config.GetString("jwt.secret_key")
	if key == "" {
		log.Fatalf("Empty string cannot be passed as secret key")
	}
	jwtSecretKey = []byte(key)
}

func GenerateToken(
	id, role string) (string, *models.ResponseError) {
	claims := jwt.MapClaims{
		USER_ID:    id,
		USER_ROLE:  role,
		CREATED_AT: time.Now().Unix(),
		EXPIRED_AT: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.User,
	*models.ResponseError) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (
			interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return jwtSecretKey, nil
		})
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims[EXPIRED_AT].(float64))
		if exp < time.Now().Unix() {
			return nil, &models.ResponseError{
				Message: "Login credintials are expired",
				Status:  http.StatusUnauthorized,
			}
		}
		return newUserInfo(claims[USER_ID].(string),
			claims[USER_ROLE].(string)), nil
	}
	return nil, &models.ResponseError{
		Message: "Invalid token",
		Status:  http.StatusUnauthorized,
	}
}
