package di

import (
	"Sharykhin/buffstream-questionnaire/database/postgres"
	"Sharykhin/buffstream-questionnaire/domains/stream/application/service"
	"Sharykhin/buffstream-questionnaire/domains/stream/repository/sql"
	"os"
)

var (
	db = postgres.NewConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		)
	streamRepo = sql.NewStreamRepository(db)

	// StreamService is a implementation of stream service that domain provides
	StreamService = service.NewStreamService(streamRepo)
)
