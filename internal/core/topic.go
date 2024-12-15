package core

import (
	"context"
	"github.com/google/uuid"
	"smart_school_for_mirea/internal/model"
)

func (c *Core) GetTopic(ctx context.Context, topicID uuid.UUID) (model.Topic, error) {
	topic, err := c.storage.GetTopic(ctx, topicID)
	if err != nil {
		if _, ok := err.(*model.ErrNotFound); ok {
			return model.Topic{}, &model.ErrNotFound{BaseError: model.BaseError{Message: "topic not found"}}
		}
		return model.Topic{}, &model.ErrInternal{}
	}
	return topic, nil
}

func (c *Core) MarkTopicAsViewed(ctx context.Context, topicID uuid.UUID, userID uuid.UUID) error {
	_, err := c.storage.GetTopic(ctx, topicID)
	if err != nil {
		if _, ok := err.(*model.ErrNotFound); ok {
			return &model.ErrBadRequest{BaseError: model.BaseError{Message: "topic not found"}}
		}
		return &model.ErrInternal{}
	}

	err = c.storage.AddViewedTopicMark(ctx, topicID, userID)
	if err != nil {
		return &model.ErrInternal{}
	}
	return nil
}
