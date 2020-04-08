package repository

import "context"

type (
	// CompanyRepository describes storage interface
	StreamRepository interface {
		Create(ctx context.Context, req CreateStream) error
	}

	// CreateStream represents income request to create a new stream in repository
	CreateStream struct {
		UUID    string
		Title    string
	}
)
