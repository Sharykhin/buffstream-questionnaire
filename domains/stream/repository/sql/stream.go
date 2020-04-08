package sql

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
	"context"
	"database/sql"
)

type (
	StreamRepository struct {
		db *sql.DB

	}
)

func (r *StreamRepository)  List(ctx context.Context, limit, offset int64) ([]models.Stream, error) {
	return nil, nil
}

func (r *StreamRepository) Count(cxt context.Context) (int64, error) {
	return 0, nil
}

// NewStreamRepository returns a new instance of sql stream repository
// that should satisfy StreamRepository interface
func NewStreamRepository(db *sql.DB) *StreamRepository {
	repo := StreamRepository{
		db: db,
	}

	return &repo
}