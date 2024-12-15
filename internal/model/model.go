package model

import "github.com/google/uuid"

type ElementType int

const (
	TopicElementType ElementType = iota
	TestElementType
)

type Course struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`

	UserInvitedInCourse bool `json:"user_invited_in_course"`

	Topics []Topic `json:"topics"`

	Elements []CourseElement `json:"elements"`
}

type Topic struct {
	ID       uuid.UUID `json:"id"`
	CourseID uuid.UUID `json:"course_id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
}

type CourseElement struct {
	ElementID   uuid.UUID `json:"element_id"`
	ElementType string    `json:"element_type"`
	Index       int64     `json:"index"`
}
