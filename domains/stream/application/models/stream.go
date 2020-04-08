package models

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
	"time"
)

type (
	Stream struct {
		UUID string	`json:"UUID"`
		Title string `json:"Title"`
		CreatedAt time.Time `json:"CreatedAt"`
		UpdatedAt time.Time `json:"UpdatedAt"`
	}
)

func NewStreamFromRepository(repoStream models.Stream) *Stream {
	stream := Stream{
		UUID: repoStream.UUID,
		Title: repoStream.Title,
		CreatedAt: repoStream.CreatedAt,
		UpdatedAt: repoStream.UpdatedAt,
	}

	return &stream
}