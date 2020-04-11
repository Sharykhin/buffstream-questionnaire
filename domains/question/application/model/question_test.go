package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
)

func TestNewQuestionFromRepository(t *testing.T) {
	assert := assert.New(t)

	repoQuestion := model.Question{
		ID:        10,
		UUID:      "UUID",
		Text:      "test question",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	actual := NewQuestionFromRepository(&repoQuestion)

	assert.NotNil(actual)
	assert.Equal(repoQuestion.UUID, actual.UUID)
	assert.Equal(repoQuestion.Text, actual.Text)
	assert.Equal(repoQuestion.CreatedAt, actual.CreatedAt)
	assert.Equal(repoQuestion.UpdatedAt, actual.UpdatedAt)
	assert.Equal(0, len(actual.Answers))

	repoQuestion.Answers = []model.Answer{
		{
			ID:         10,
			QuestionID: repoQuestion.ID,
			Text:       "test answer",
			IsCorrect:  true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	actual = NewQuestionFromRepository(&repoQuestion)
	assert.Equal(1, len(actual.Answers))
	assert.Equal(repoQuestion.Answers[0].ID, actual.Answers[0].ID)
	assert.Equal(repoQuestion.Answers[0].Text, actual.Answers[0].Text)
	assert.Equal(repoQuestion.Answers[0].IsCorrect, actual.Answers[0].IsCorrect)
	assert.Equal(repoQuestion.Answers[0].CreatedAt, actual.Answers[0].CreatedAt)
	assert.Equal(repoQuestion.Answers[0].UpdatedAt, actual.Answers[0].UpdatedAt)
}
