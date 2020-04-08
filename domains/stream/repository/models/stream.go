package models

import "time"

type (
	// Stream is a model on a repository level that maps data from storage to a struct
	Stream struct {
		ID int64
		UUID string
		Title string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)