package model

import "time"

type (
	// Question represents a model on application level that can be converted into view model
	// by using json marshaling
	Question struct {
		UUID      string    `json:"UUID"`
		Text      string    `json:"Text"`
		CreatedAt time.Time `json:"CreatedAt"`
		UpdatedAt time.Time `json:"UpdatedAt"`
		Answers   []Answer  `json:"Answers"`
	}
)
