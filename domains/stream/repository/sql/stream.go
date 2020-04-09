package sql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"Sharykhin/buffstream-questionnaire/domains/stream/repository/models"
)

type (
	// StreamRepository is a sql implementation of stream repository interface
	StreamRepository struct {
		db *sql.DB
	}
)

// List returns a limited number of streams records. If there are no records empty slice will be returned, not nil.
func (r *StreamRepository) List(ctx context.Context, limit, offset int64) ([]models.Stream, error) {
	streams := make([]models.Stream, 0)
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

// Create inserts a new stream records into the database
func (r *StreamRepository) Create(ctx context.Context, UUID, title string) (*models.Stream, error) {
	stream := models.Stream{
		UUID:      UUID,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "INSERT INTO streams(uuid, title, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id"

	var id int64
	err := r.db.QueryRowContext(ctx, query, UUID, title, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert a new stream record: %v", err)
	}

	stream.ID = id

	return &stream, nil
}

// NewStreamRepository returns a new instance of sql stream repository
// that should satisfy StreamRepository interface
func NewStreamRepository(db *sql.DB) *StreamRepository {
	repo := StreamRepository{
		db: db,
	}

	return &repo
}
