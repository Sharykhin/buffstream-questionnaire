package model

import "time"

type (
	// Answer model represents answer on an application level
	Answer struct {
		ID        int64     `json:"ID"`
		Text      string    `json:"Text"`
		IsCorrect bool      `json:"IsCorrect"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
	}
)
