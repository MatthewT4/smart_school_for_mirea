package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"smart_school_for_mirea/internal/model"
)

type Storage interface {
	CreateUser(ctx context.Context, user model.User) (result model.User, err error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error)

	GetProduct(ctx context.Context, productID uuid.UUID) (model.Product, error)
}

type Core struct {
	storage Storage

	authSecretKey string
	authTTL       int64

	logger *slog.Logger
}

func NewCore(storage Storage, authSecretKey string, authTTL int64, logger *slog.Logger) *Core {
	return &Core{
		storage: storage,

		authSecretKey: authSecretKey,
		authTTL:       authTTL,

		logger: logger,
	}
}

func (c *Core) GetProduct(ctx context.Context, productID uuid.UUID) (model.Product, error) {
	product, err := c.storage.GetProduct(ctx, productID)
	if err != nil {
		if errors.Is(err, &model.ErrNotFound{}) {
			return model.Product{}, err
		}
		return model.Product{}, fmt.Errorf("get product: %w", err)
	}
	return product, err
}
