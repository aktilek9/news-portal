package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

type AppError struct {
	Err     error  `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{Err: err, Code: code, Message: message}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func MapPostgresError(err error) *AppError {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return NewAppError(http.StatusConflict, "Resource already exists", err)
		case "23503":
			return NewAppError(http.StatusNotFound, "Related resource not found", err)
		case "23502":
			return NewAppError(http.StatusBadRequest, "Required field is missing", err)
		case "22001":
			return NewAppError(http.StatusBadRequest, "Field too long", err)
		default:
			return NewAppError(http.StatusInternalServerError, "Database error", err)
		}
	}

	return NewAppError(http.StatusInternalServerError, "Unknown error", err)
}

