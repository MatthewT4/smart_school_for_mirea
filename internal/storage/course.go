package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"smart_school_for_mirea/internal/model"
)

func (p *PgStorage) GetCourse(ctx context.Context, courseID uuid.UUID) (model.Course, error) {
	rows, err := p.connections.Query(ctx, queryGetCourse, courseID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Course{}, &model.ErrNotFound{}
		}
		return model.Course{}, fmt.Errorf("query course: %w", err)
	}
	course, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Course])
	if err != nil {
		return model.Course{}, fmt.Errorf("collect row from table 'course': %w", err)
	}

	topicRows, err := p.connections.Query(ctx, queryGetCourseTopics, courseID)
	if err != nil {
		return model.Course{}, fmt.Errorf("query topics: %w", err)
	}
	course.Topics, err = pgx.CollectRows(topicRows, pgx.RowToStructByNameLax[model.Topic])
	if err != nil {
		return model.Course{}, fmt.Errorf("collect topics: %w", err)
	}

	elementRows, err := p.connections.Query(ctx, queryGetCourseElements, courseID)
	if err != nil {
		return model.Course{}, fmt.Errorf("query elements: %w", err)
	}
	course.Elements, err = pgx.CollectRows(elementRows, pgx.RowToStructByNameLax[model.CourseElement])
	if err != nil {
		return model.Course{}, fmt.Errorf("collect elements: %w", err)
	}

	return course, nil
}

func (p *PgStorage) FindCourses(ctx context.Context, invitedUserID *uuid.UUID, nameLike *string) ([]model.Course, error) {
	var courseIDs *[]uuid.UUID
	if invitedUserID != nil {
		userCourseIDs, err := p.FindUserCourseIDs(ctx, *invitedUserID)
		if err != nil {
			return nil, fmt.Errorf("find user courses: %w", err)
		}
		courseIDs = &userCourseIDs
	}

	courseRows, err := p.connections.Query(ctx, queryFindCourses, courseIDs, nameLike)
	if err != nil {
		return nil, fmt.Errorf("query courses: %w", err)
	}
	courses, err := pgx.CollectRows(courseRows, pgx.RowToStructByNameLax[model.Course])
	if err != nil {
		return nil, fmt.Errorf("collect courses: %w", err)
	}
	return courses, nil
}

func (p *PgStorage) FindUserCourseIDs(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	userCourseRows, err := p.connections.Query(ctx, queryUserCourses, userID)
	if err != nil {
		return nil, fmt.Errorf("query user courses: %w", err)
	}
	userCourseIDs, err := pgx.CollectRows(userCourseRows, pgx.RowTo[uuid.UUID])
	if err != nil {
		return nil, fmt.Errorf("collect user courses: %w", err)
	}
	return userCourseIDs, nil
}
