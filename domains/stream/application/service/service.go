package service

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/application/models"
	"Sharykhin/buffstream-questionnaire/domains/stream/repository"
	"context"
)

type (
	// StreamService describes general api that stream domain provides
	StreamService interface {
		List(ctx context.Context, limit, offset int64) ([]models.Stream, int64, error)
	}

	// StreamHandler is a concrete implementation of StreamService interface
	StreamHandler struct {
		streamRepo repository.StreamRepository
	}
)

func (s *StreamHandler) List(ctx context.Context, limit, offset int64) ([]models.Stream, int64, error) {
	return nil, 0, nil
}

// NewStreamService create a new instance of stream service
func NewStreamService(streamRepo repository.StreamRepository) *StreamHandler {
	srv := StreamHandler{
		streamRepo: streamRepo,
	}

	return &srv
}