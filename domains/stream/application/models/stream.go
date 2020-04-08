package models

import "time"

type (
	Stream struct {
		UUID string
		Title string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)