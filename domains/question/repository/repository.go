package repository

import (
	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	"context"
)

type (
	QuestionRepository interface {
		FindOneByUUIDWithAnswers(ctx context.Context, UUID string) (*model.Question, error)
	}

	AnswerRepository interface {
	}
)
