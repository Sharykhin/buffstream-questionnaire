package models

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
	"time"
)

type (
	Stream struct {
		UUID      string    `json:"UUID"`
		Title     string    `json:"Title"`
		CreatedAt time.Time `json:"CreatedAt"`
		UpdatedAt time.Time `json:"UpdatedAt"`
	}
)

func NewStreamFromRepository(repoModel *models.Stream) *Stream {
	stream := Stream{
		UUID:      repoModel.UUID,
		Title:     repoModel.Title,
		CreatedAt: repoModel.CreatedAt,
		UpdatedAt: repoModel.UpdatedAt,
	}

	return &stream
}
