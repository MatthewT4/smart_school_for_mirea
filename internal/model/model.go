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

	Topics []Topic      `json:"topics"`
	Tests  []TestEntity `json:"tests"`

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

type TestEntity struct {
	ID       uuid.UUID `json:"id"`
	CourseID uuid.UUID `json:"course_id"`
	Title    string    `json:"title"`

	Elements []TestElement `json:"elements"`

	MaxScore    *int64 `json:"max_score"`
	ResultScore *int64 `json:"result_score"`
}

type TestElement struct {
	ID            uuid.UUID `json:"id"`
	TestID        uuid.UUID `json:"test_id"`
	Title         string    `json:"title"`
	CorrectAnswer string    `json:"correct_answer"`

	Index int64 `json:"index"`

	UserAnswer *string `json:"user_answer"`
	Score      *int64  `json:"score"`
}

type TestElementAnswer struct {
	ElementID uuid.UUID `json:"element_id"`
	Answer    string    `json:"answer"`
}
