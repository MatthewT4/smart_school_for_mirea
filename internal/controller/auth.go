package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"smart_school_for_mirea/api"

	"smart_school_for_mirea/internal/model"
)

func (h *handlers) SignUp(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	request := model.SignUpRequest{}
	if err = json.Unmarshal(body, &request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.core.SignUp(ctx.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := api.JWTToken{Token: &token}
	return ctx.JSON(http.StatusOK, resp)
}

func (h *handlers) SignIn(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	request := model.SignInRequest{}

	if err = json.Unmarshal(body, &request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.core.SignIn(ctx.Request().Context(), request)
	if err != nil {
		return convertErrorToResponse(err)
	}

	resp := api.JWTToken{Token: &token}
	return ctx.JSON(http.StatusOK, resp)
}
