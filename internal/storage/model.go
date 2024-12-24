package storage

import "github.com/google/uuid"

type TestResult struct {
	ID     uuid.UUID `db:"id"`
	TestID uuid.UUID `db:"test_id"`
	UserID uuid.UUID `db:"user_id"`

	CountCorrectAnswers int64 `db:"count_correct_answers"`
	CountAnswers        int64 `db:"count_answers"`
}

type TestElementResult struct {
	ID           uuid.UUID `db:"id"`
	TestResultID uuid.UUID `db:"test_result_id"`
	ElementID    uuid.UUID `db:"element_id"`

	UserAnswer string `db:"user_answer"`
	Score      int64  `db:"score"`
}
