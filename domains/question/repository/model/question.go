package model

import "time"

type (
	// Question is a model on a repository level that would maps data to storage
	Question struct {
		ID        int64
		UUID      string
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Answers   []Answer
	}
)
