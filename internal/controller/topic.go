package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"net/http"
)

// (GET /topics/{topic_id})
func (h *handlers) GetTopic(ctx echo.Context, topicId openapi_types.UUID) error {
	topic, err := h.core.GetTopic(ctx.Request().Context(), topicId)
	if err != nil {
		return convertErrorToResponse(err)
	}
	return ctx.JSON(http.StatusOK, topic)
}

// (POST /topics/{topic_id})
func (h *handlers) ViewedTopic(ctx echo.Context, topicId openapi_types.UUID) error {
	userID := ctx.Get(userIDCtx).(uuid.UUID)

	err := h.core.MarkTopicAsViewed(ctx.Request().Context(), topicId, userID)
	if err != nil {
		return convertErrorToResponse(err)
	}
	ctx.Response().Status = http.StatusNoContent
	return nil
}
