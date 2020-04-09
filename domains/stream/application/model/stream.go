package model

import (
	"time"

	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
)

type (
	// Stream represents steam model on application layer.
	Stream struct {
		UUID      string    `json:"UUID"`
		Title     string    `json:"Title"`
		CreatedAt time.Time `json:"CreatedAt"`
		UpdatedAt time.Time `json:"UpdatedAt"`
	}
)

// NewStreamFromRepository creates a new stream model based on a model that repository returned.
func NewStreamFromRepository(repoModel *models.Stream) *Stream {
	stream := Stream{
		UUID:      repoModel.UUID,
		Title:     repoModel.Title,
		CreatedAt: repoModel.CreatedAt,
		UpdatedAt: repoModel.UpdatedAt,
	}

	return &stream
}
