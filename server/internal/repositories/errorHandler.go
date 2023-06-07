package repositories

import (
	"net/http"
	"strings"

	"github.com/ak-karimzai/cp-db/internal/models"
)

func errorJSON(
	err error, status ...int) *models.ResponseError {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var finalErr *models.ResponseError

	switch {
	case strings.Contains(err.Error(), "SQLSTATE 23505"):
		finalErr = &models.ResponseError{
			Message: "duplicate value violated unque constrint",
			Status:  http.StatusForbidden,
		}
	case strings.Contains(err.Error(), "SQLSTATE 23403"):
		finalErr = &models.ResponseError{
			Message: "foriegn key violation",
			Status:  http.StatusForbidden,
		}
	case strings.Contains(err.Error(),
		"pq: column reference \"id\" is ambiguous"):
		finalErr = &models.ResponseError{
			Message: "not found",
			Status:  http.StatusNotFound,
		}
	case strings.Contains(err.Error(),
		"sql: no rows in result set"):
		finalErr = &models.ResponseError{
			Message: "not found",
			Status:  http.StatusNotFound,
		}
	case strings.Contains(err.Error(),
		"pq: duplicate key value violates unique constraint"):
		finalErr = &models.ResponseError{
			Message: "Already exist",
			Status:  http.StatusConflict,
		}
	default:
		finalErr = &models.ResponseError{
			Message: err.Error(),
			Status:  statusCode,
		}
	}
	return finalErr
}
