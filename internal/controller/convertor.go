package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smart_school_for_mirea/internal/model"
)

func convertErrorToResponse(err error) error {
	code := http.StatusInternalServerError
	message := err.Error()

	switch err.(type) {
	case *model.ErrNotFound:
		code = http.StatusNotFound
	case *model.ErrBadRequest:
		code = http.StatusBadRequest
	}

	return echo.NewHTTPError(code, message)
}
