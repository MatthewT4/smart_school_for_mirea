package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"log/slog"
	"net/http"
	"smart_school_for_mirea/api"
	"smart_school_for_mirea/internal/core"
)

func (h *handlers) FindCourses(ctx echo.Context, params api.FindCoursesParams) error {
	userID := ctx.Get(userIDCtx).(uuid.UUID)
	courses, err := h.core.FindCourses(ctx.Request().Context(), params.Like, params.My, userID)
	if err != nil {
		return convertErrorToResponse(err)
	}
	return ctx.JSON(http.StatusOK, courses)
}

func (h *handlers) GetCourse(ctx echo.Context, courseId openapi_types.UUID) error {
	course, err := h.core.GetCourse(ctx.Request().Context(), courseId)
	if err != nil {
		return convertErrorToResponse(err)
	}
	return ctx.JSON(http.StatusOK, course)
}

func (h *handlers) InviteInCourse(ctx echo.Context, courseId openapi_types.UUID) error {
	//TODO implement me
	panic("implement me")
}

func newHandlers(core *core.Core, authSecretKey string, logger *slog.Logger) *handlers {
	return &handlers{
		core: core,

		authSecretKey: authSecretKey,

		logger: logger,
	}
}
