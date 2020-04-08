package service

import (
	"Sharykhin/buffstream-questionnaire/domains/question/application/model"
	"Sharykhin/buffstream-questionnaire/domains/question/repository"
	"context"
)

type (
	// StreamService describes general api that stream domain provides
	QuestionService interface {
		GetOneByUUID(ctx context.Context, UUID string) (*model.Question, error)
	}

	// QuestionHandler is a concrete implementation of StreamService interface
	QuestionHandler struct {
		questionRepo repository.QuestionRepository
	}
)

// NewQuestionService create a new instance of question service
func NewQuestionService(questionRepo repository.QuestionRepository) *QuestionHandler {
	srv := QuestionHandler{
		questionRepo: questionRepo,
	}

	return &srv
}
