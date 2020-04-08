package di

import (
	"Sharykhin/buffstream-questionnaire/database/postgres"
	questionSrv "Sharykhin/buffstream-questionnaire/domains/question/application/service"
	questionSQL "Sharykhin/buffstream-questionnaire/domains/question/repository/sql"
	streamSrv "Sharykhin/buffstream-questionnaire/domains/stream/application/service"
	streamSQL "Sharykhin/buffstream-questionnaire/domains/stream/repository/sql"
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
	streamRepo   = streamSQL.NewStreamRepository(db)
	questionRepo = questionSQL.NewQuestionRepository(db)
	// StreamService is a implementation of stream service that domain provides
	StreamService   = streamSrv.NewStreamService(streamRepo)
	QuestionService = questionSrv.NewQuestionService(questionRepo)
)
