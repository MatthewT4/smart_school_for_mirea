package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"smart_school_for_mirea/internal/model"
)

func (p *PgStorage) CreateUser(ctx context.Context, user model.User) (result model.User, err error) {
	const op = "storage.CreateUser"

	err = p.connections.QueryRow(
		ctx,
		queryCreateUser,
		user.UUID,
		user.Email,
		user.HashPassword,
	).Scan(&result.UUID, &result.Email, &result.HashPassword)
	if err != nil {
		return result, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}

func (p *PgStorage) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	const op = "storage.GetUserByUsername"

	rows, err := p.connections.Query(
		ctx,
		queryGetUserByEmail,
		username,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (p *PgStorage) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	const op = "storage.GetUserByUsername"

	rows, err := p.connections.Query(
		ctx,
		queryGetUserByID,
		id,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
