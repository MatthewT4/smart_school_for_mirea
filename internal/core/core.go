package core

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"smart_school_for_mirea/internal/model"
)

type Storage interface {
	CreateUser(ctx context.Context, user model.User) (result model.User, err error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error)

	FindCourses(ctx context.Context, invitedUserID *uuid.UUID, nameLike *string) ([]model.Course, error)
	FindUserCourseIDs(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
	GetCourse(ctx context.Context, courseID uuid.UUID) (model.Course, error)
	AddUserInCourse(ctx context.Context, userID uuid.UUID, courseID uuid.UUID) error

	GetTopic(ctx context.Context, topicID uuid.UUID) (model.Topic, error)
	AddViewedTopicMark(ctx context.Context, topicID uuid.UUID, userID uuid.UUID) error

	GetTestWithResult(ctx context.Context, testID uuid.UUID, userID uuid.UUID) (model.TestEntity, error)
	ApplyTestResult(ctx context.Context, userID uuid.UUID, test model.TestEntity) error
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
