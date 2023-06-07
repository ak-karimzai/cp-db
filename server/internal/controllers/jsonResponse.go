package controllers

import (
	"github.com/ak-karimzai/cp-db/internal/models"
)

func jsonResponse(
	responseErr *models.ResponseError, data interface{}) *models.JsonResponse {
	return &models.JsonResponse{
		Error: responseErr,
		Data:  data,
	}
}
