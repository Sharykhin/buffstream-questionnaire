package service

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/application/model"
	"Sharykhin/buffstream-questionnaire/domains/stream/repository"
	"Sharykhin/buffstream-questionnaire/uuid"
	"context"
	"fmt"
)

type (
	// StreamService describes general api that stream domain provides
	StreamService interface {
		List(ctx context.Context, limit, offset int64) ([]model.Stream, int64, error)
		Create(ctx context.Context, title string) (*model.Stream, error)
	}

	// StreamHandler is a concrete implementation of StreamService interface
	StreamHandler struct {
		uuidGenerator uuid.Generator
		streamRepo    repository.StreamRepository
	}
)

// List returns a limited number of streams along with a total number of them
func (s *StreamHandler) List(ctx context.Context, limit, offset int64) ([]model.Stream, int64, error) {
	var streams []model.Stream
	repoStreams, err := s.streamRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get list of streams from a repository: %v", err)
	}
	total, err := s.streamRepo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total number of streams from a repository: %v", err)
	}
	for _, repoStream := range repoStreams {
		stream := model.NewStreamFromRepository(&repoStream)
		streams = append(streams, *stream)
	}

	return streams, total, nil
}

// Create creates a new stream
func (s *StreamHandler) Create(ctx context.Context, title string) (*model.Stream, error) {
	u := s.uuidGenerator.NewV4()

	repoStream, err := s.streamRepo.Create(ctx, u, title)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new stream: %v", err)
	}

	stream := model.NewStreamFromRepository(repoStream)

	return stream, nil
}

// NewStreamService create a new instance of stream service
func NewStreamService(uuidGenerator uuid.Generator, streamRepo repository.StreamRepository) *StreamHandler {
	srv := StreamHandler{
		uuidGenerator: uuidGenerator,
		streamRepo:    streamRepo,
	}

	return &srv
}
