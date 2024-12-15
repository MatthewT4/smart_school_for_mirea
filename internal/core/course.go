package core

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"smart_school_for_mirea/internal/model"
)

func (c *Core) GetCourse(ctx context.Context, courseID uuid.UUID) (model.Course, error) {
	course, err := c.storage.GetCourse(ctx, courseID)
	if err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func (c *Core) FindCourses(ctx context.Context, nameLike *string, isMyCourses *bool, userID uuid.UUID) ([]model.Course, error) {
	var invitedUserID *uuid.UUID
	if isMyCourses != nil && *isMyCourses == true {
		invitedUserID = &userID
	}
	courses, err := c.storage.FindCourses(ctx, invitedUserID, nameLike)
	if err != nil {
		c.logger.Error("fail find courses", slog.Any("error", err))
		return nil, &model.ErrInternal{}
	}

	userInviteToCourseIDs, err := c.storage.FindUserCourseIDs(ctx, userID)
	if err != nil {
		c.logger.Error("fail find user invite to courses", slog.Any("error", err))
		return nil, &model.ErrInternal{}
	}
	userInvitedTo := make(map[uuid.UUID]struct{})
	for _, courseID := range userInviteToCourseIDs {
		userInvitedTo[courseID] = struct{}{}
	}

	for key, course := range courses {
		if _, find := userInvitedTo[course.ID]; find {
			course.UserInvitedInCourse = true
			courses[key] = course
		}
	}
	return courses, nil
}
