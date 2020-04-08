package sql

import (
	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
	"context"
	"database/sql"
	"fmt"
)

type (
	StreamRepository struct {
		db *sql.DB

	}
)

// List returns a limited number of rows
func (r *StreamRepository) List(ctx context.Context, limit, offset int64) ([]models.Stream, error) {
	var streams []models.Stream
	query := "SELECT * FROM streams ORDER BY created_at DESC OFFSET $1 LIMIT $2"

	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to execute sql %s query: %v", query, err)
	}
	//TODO: check error
	defer rows.Close()

	for rows.Next() {
		var stream models.Stream
		err := rows.Scan(&stream.ID, &stream.UUID, &stream.Title, &stream.CreatedAt, &stream.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan stream record: %v", err)
		}
		streams = append(streams, stream)
	}

	return streams, rows.Err()
}

// Count return total number of stream records
func (r *StreamRepository) Count(ctx context.Context) (int64, error) {
	var total int64
	query := "SELECT COUNT(id) FROM streams"

	row := r.db.QueryRowContext(ctx, query)
	err := row.Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to execute %s query: %v", query, err)
	}

	return total, nil
}

// NewStreamRepository returns a new instance of sql stream repository
// that should satisfy StreamRepository interface
func NewStreamRepository(db *sql.DB) *StreamRepository {
	repo := StreamRepository{
		db: db,
	}

	return &repo
}