package repository

import (
	"context"

	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
)

type (
	// CompanyRepository describes storage interface
	StreamRepository interface {
		List(ctx context.Context, limit, offset int64) ([]models.Stream, error)
		Count(cxt context.Context) (int64, error)
	}

	// CreateStream represents income request to create a new stream in repository
	CreateStream struct {
		UUID  string
		Title string
	}
)
