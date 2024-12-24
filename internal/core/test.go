package core

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"smart_school_for_mirea/internal/model"
)

func (c *Core) GetTest(ctx context.Context, testID uuid.UUID, userID uuid.UUID) (model.TestEntity, error) {
	test, err := c.storage.GetTestWithResult(ctx, testID, userID)
	if err != nil {
		if _, ok := err.(*model.ErrNotFound); ok {
			return model.TestEntity{}, &model.ErrNotFound{BaseError: model.BaseError{Message: "test not found"}}
		}
		return model.TestEntity{}, &model.ErrInternal{}
	}
	return test, nil
}

func (c *Core) ApplyTestResult(
	ctx context.Context,
	userID uuid.UUID,
	testID uuid.UUID,
	answers []model.TestElementAnswer,
) error {
	test, err := c.storage.GetTestWithResult(ctx, testID, userID)
	if err != nil {
		if _, ok := err.(*model.ErrNotFound); ok {
			return &model.ErrNotFound{BaseError: model.BaseError{Message: "test not found"}}
		}
		return &model.ErrInternal{}
	}

	if test.ResultScore != nil {
		return &model.ErrBadRequest{BaseError: model.BaseError{Message: "test already applied"}}
	}

	answersInMap := make(map[uuid.UUID]model.TestElementAnswer, len(answers))
	for _, a := range answers {
		answersInMap[a.ElementID] = a
	}

	var resultScore int64
	for idx, element := range test.Elements {
		var score int64
		var userAnswer string
		if uAnsr, ok := answersInMap[element.ID]; ok {
			if uAnsr.Answer == element.CorrectAnswer {
				score = 1
			}
			userAnswer = uAnsr.Answer
		}

		resultScore += score
		element.Score = &score
		element.UserAnswer = &userAnswer

		test.Elements[idx] = element
	}

	test.ResultScore = &resultScore
	maxScore := int64(len(test.Elements))
	test.MaxScore = &maxScore

	err = c.storage.ApplyTestResult(ctx, userID, test)
	if err != nil {
		c.logger.Error("Error apply test result", slog.Any("error", err))
		return &model.ErrInternal{}
	}
	return nil
}
