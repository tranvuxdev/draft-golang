package apperror

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadReuqest = errors.New("bad request")
	ErrInternal   = errors.New("internal server error")
	ErrDuplicate  = errors.New("duplicate")
)

func StatusCode(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound // 404
	case errors.Is(err, ErrBadReuqest):
		return http.StatusBadRequest // 400
	case errors.Is(err, ErrDuplicate):
		return http.StatusConflict // 409
	default:
		return http.StatusInternalServerError // 500
	}
}
