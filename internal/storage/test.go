package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"smart_school_for_mirea/internal/model"
)

func (p *PgStorage) GetTestInfo(ctx context.Context, testID uuid.UUID) (model.TestEntity, error) {
	testRows, err := p.connections.Query(ctx, queryGetTestInfo, testID)
	if err != nil {
		return model.TestEntity{}, fmt.Errorf("query test info: %w", err)
	}
	test, err := pgx.CollectOneRow(testRows, pgx.RowToStructByNameLax[model.TestEntity])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.TestEntity{}, &model.ErrNotFound{}
		}
		return model.TestEntity{}, fmt.Errorf("collect test info: %w", err)
	}

	return test, nil
}

func (p *PgStorage) GetTestElementsInfo(ctx context.Context, testID uuid.UUID) ([]model.TestElement, error) {
	testElementRows, err := p.connections.Query(ctx, queryGetTestElementsInfo, testID)
	if err != nil {
		return nil, fmt.Errorf("query test info: %w", err)
	}
	elements, err := pgx.CollectRows(testElementRows, pgx.RowToStructByNameLax[model.TestElement])
	if err != nil {
		return nil, fmt.Errorf("collect test info: %w", err)
	}

	return elements, nil
}

func (p *PgStorage) GetTestWithResult(ctx context.Context, testID uuid.UUID, userID uuid.UUID) (model.TestEntity, error) {
	test, err := p.GetTestInfo(ctx, testID)
	if err != nil {
		return model.TestEntity{}, err
	}

	test.Elements, err = p.GetTestElementsInfo(ctx, testID)
	if err != nil {
		return model.TestEntity{}, err
	}

	testResultRows, err := p.connections.Query(ctx, queryGetTestResult, testID, userID)
	if err != nil {
		return model.TestEntity{}, fmt.Errorf("query test info: %w", err)
	}
	testResult, err := pgx.CollectOneRow(testResultRows, pgx.RowToStructByNameLax[TestResult])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Its ok, test may haven't result for user, if user not finished test.
			return test, nil
		}
		return model.TestEntity{}, fmt.Errorf("collect test result info: %w", err)
	}

	test.ResultScore = &testResult.CountCorrectAnswers
	test.MaxScore = &testResult.CountAnswers

	testElementResultRows, err := p.connections.Query(ctx, queryGetTestElementResults, testResult.ID)

	elementResults, err := pgx.CollectRows(testElementResultRows, pgx.RowToStructByNameLax[TestElementResult])
	if err != nil {
		return model.TestEntity{}, fmt.Errorf("collect test elements result info: %w", err)
	}

	mapElementResults := make(map[uuid.UUID]TestElementResult, len(elementResults))
	for _, elementResult := range elementResults {
		mapElementResults[elementResult.ElementID] = elementResult
	}

	for idx, element := range test.Elements {
		if result, ok := mapElementResults[element.ID]; ok {
			test.Elements[idx].UserAnswer = &result.UserAnswer
			test.Elements[idx].Score = &result.Score
		}
	}
	return test, nil
}

func (p *PgStorage) ApplyTestResult(ctx context.Context, userID uuid.UUID, test model.TestEntity) error {
	testResultID := uuid.New()
	_, err := p.connections.Exec(ctx, execAddTestResult, testResultID, test.ID, userID, test.ResultScore, test.MaxScore)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	for _, el := range test.Elements {
		_, err := p.connections.Exec(ctx, execAddTestElementResult, uuid.New(), testResultID, el.ID, *el.UserAnswer, *el.Score)
		if err != nil {
			return fmt.Errorf("exec: %w", err)
		}
	}
	return nil
}
