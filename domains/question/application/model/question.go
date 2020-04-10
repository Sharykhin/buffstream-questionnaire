package model

import (
	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	"time"
)

type (
	// Question represents a model on application level that can be converted into view model
	// by using json marshaling
	Question struct {
		UUID      string    `json:"UUID"`
		Text      string    `json:"Text"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
		Answers   []Answer  `json:"Answers,omitempty"`
	}
)

// NewQuestionFromRepository creates question model on application layer based on what repository returned
func NewQuestionFromRepository(repoModel *model.Question) *Question {
	question := Question{
		UUID:      repoModel.UUID,
		Text:      repoModel.Text,
		CreatedAt: repoModel.CreatedAt,
		UpdatedAt: repoModel.UpdatedAt,
	}

	var answers []Answer
	for _, repoAnswer := range repoModel.Answers {
		answer := NewAnswerFromRepository(&repoAnswer)
		answers = append(answers, *answer)
	}

	question.Answers = answers

	return &question
}
