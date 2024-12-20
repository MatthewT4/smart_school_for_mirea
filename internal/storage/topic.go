package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"smart_school_for_mirea/internal/model"
)

func (p *PgStorage) GetTopic(ctx context.Context, topicID uuid.UUID) (model.Topic, error) {
	rows, err := p.connections.Query(ctx, queryGetTopic, topicID)
	if err != nil {
		return model.Topic{}, fmt.Errorf("queryex: %w", err)
	}
	topic, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Topic])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Topic{}, &model.ErrNotFound{}
		}

		return model.Topic{}, err
	}

	return topic, nil
}

func (p *PgStorage) AddViewedTopicMark(ctx context.Context, topicID uuid.UUID, userID uuid.UUID) error {
	_, err := p.connections.Exec(ctx, execAddTopicViewedRow, userID, topicID)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
