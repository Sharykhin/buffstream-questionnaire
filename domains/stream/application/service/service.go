package service

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/application/models"
	"Sharykhin/buffstream-questionnaire/domains/stream/repository"
	"context"
	"fmt"
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

// List returns a limited number of streams along with a total number of them
func (s *StreamHandler) List(ctx context.Context, limit, offset int64) ([]models.Stream, int64, error) {
	var streams []models.Stream
	repoStreams, err := s.streamRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get list of streams from a repository: %v", err)
	}
	total, err := s.streamRepo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total number of streams from a repository: %v", err)
	}
	for _, repoStream := range repoStreams {
		stream := models.NewStreamFromRepository(&repoStream)
		streams = append(streams, *stream)
	}

	return streams, total, nil
}

// NewStreamService create a new instance of stream service
func NewStreamService(streamRepo repository.StreamRepository) *StreamHandler {
	srv := StreamHandler{
		streamRepo: streamRepo,
	}

	return &srv
}
