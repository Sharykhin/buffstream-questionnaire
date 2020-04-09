package repository

import (
	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	"context"
)

type (
	// QuestionRepository describes API on the repository layer
	QuestionRepository interface {
		FindOneByIDWithAnswers(ctx context.Context, UUID string) (*model.Question, error)
		FindListByStreamID(ctx context.Context, UUID string) ([]model.Question, error)
		FindListByStreamIDs(ctx context.Context, UUIDs []string) ([]model.Stream, error)
	}
)
