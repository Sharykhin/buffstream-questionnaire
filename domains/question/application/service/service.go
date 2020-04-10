package service

import (
	"context"
	"errors"
	"fmt"

	"Sharykhin/buffstream-questionnaire/domains/question/application/model"
	"Sharykhin/buffstream-questionnaire/domains/question/repository"
	appErrors "Sharykhin/buffstream-questionnaire/errors"
)

type (
	// StreamService describes general api that stream domain provides
	QuestionService interface {
		GetOneByID(ctx context.Context, UUID string) (*model.Question, error)
		GetAllByStreamID(ctx context.Context, UUID string) ([]model.Question, error)
		GetAllByStreamIDs(ctx context.Context, UUIDs []string) (model.Streams, error)
	}

	// QuestionHandler is a concrete implementation of StreamService interface
	QuestionHandler struct {
		questionRepo repository.QuestionRepository
	}
)

// GetOneByID returns a question with all related answers
func (s *QuestionHandler) GetOneByID(ctx context.Context, UUID string) (*model.Question, error) {

	repoQuestion, err := s.questionRepo.FindOneByIDWithAnswers(ctx, UUID)

	if err != nil {
		if errors.Is(err, appErrors.ResourceNotFound) {
			return nil, err
		}

		return nil, fmt.Errorf("failed to get question by its uuid: %v", err)
	}

	question := model.NewQuestionFromRepository(repoQuestion)

	return question, nil
}

// GetAllByStreamID find all question for a specific stream by its unique identifier
func (s *QuestionHandler) GetAllByStreamID(ctx context.Context, UUID string) ([]model.Question, error) {
	repoQuestions, err := s.questionRepo.FindListByStreamID(ctx, UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions of a specific stream: %v", err)
	}

	var questions []model.Question
	for _, repoQuestion := range repoQuestions {
		question := model.NewQuestionFromRepository(&repoQuestion)
		questions = append(questions, *question)
	}

	return questions, nil
}

// GetAllByStreamIDs returns question list aggregated by stream identifiers
func (s *QuestionHandler) GetAllByStreamIDs(ctx context.Context, UUIDs []string) (model.Streams, error) {
	repoStreams, err := s.questionRepo.FindListByStreamIDs(ctx, UUIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of questions by streams ids: %v", err)
	}

	streams := make(model.Streams, len(repoStreams))
	for _, repoStream := range repoStreams {
		streams[repoStream.UUID] = make([]model.Question, len(repoStream.Questions))

		for j, repoQuestion := range repoStream.Questions {
			question := model.NewQuestionFromRepository(&repoQuestion)
			streams[repoStream.UUID][j] = *question
		}
	}

	return streams, nil
}

// NewQuestionService create a new instance of question service
func NewQuestionService(questionRepo repository.QuestionRepository) *QuestionHandler {
	srv := QuestionHandler{
		questionRepo: questionRepo,
	}

	return &srv
}
