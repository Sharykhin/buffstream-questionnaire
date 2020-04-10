package repository

//go:generate mockery -name StreamRepository

import (
	"context"

	"Sharykhin/buffstream-questionnaire/domains/stream/repository/model"
)

type (
	// CompanyRepository describes storage interface
	StreamRepository interface {
		List(ctx context.Context, limit, offset int64) ([]model.Stream, error)
		Count(cxt context.Context) (int64, error)
		Create(ctx context.Context, UUID, title string) (*model.Stream, error)
	}
)
