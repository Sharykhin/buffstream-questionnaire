package model

import "time"

type (
	// Answer is a model on repository level that would maps data to a storage
	Answer struct {
		ID         int64
		QuestionID int64
		Text       string
		IsCorrect  bool
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
