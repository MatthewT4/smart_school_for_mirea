package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"io"
	"net/http"
	"smart_school_for_mirea/internal/model"
)

func (h *handlers) GetTest(ctx echo.Context, testId openapi_types.UUID) error {
	userID := ctx.Get(userIDCtx).(uuid.UUID)

	test, err := h.core.GetTest(ctx.Request().Context(), testId, userID)
	if err != nil {
		return convertErrorToResponse(err)
	}
	return ctx.JSON(http.StatusOK, test)
}

func (h *handlers) ApplyTest(ctx echo.Context, testId openapi_types.UUID) error {
	userID := ctx.Get(userIDCtx).(uuid.UUID)

	answers := make([]model.TestElementAnswer, 0)
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return convertErrorToResponse(&model.ErrInternal{})
	}
	err = json.Unmarshal(body, &answers)
	if err != nil {
		return convertErrorToResponse(&model.ErrBadRequest{BaseError: model.BaseError{Message: "Unmarshal error"}})
	}

	err = h.core.ApplyTestResult(ctx.Request().Context(), userID, testId, answers)
	if err != nil {
		return convertErrorToResponse(err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
