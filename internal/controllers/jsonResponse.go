package controllers

import (
	"github.com/ak-karimzai/cp-db/internal/models"
)

type envelope map[string]interface{}

func jsonResponse(
	responseErr *models.ResponseError, data interface{}) *models.JsonResponse {
	return &models.JsonResponse{
		Error: responseErr,
		Data:  envelope{"data": data},
	}
}
