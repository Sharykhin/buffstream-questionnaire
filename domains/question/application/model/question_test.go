package model

import (
	"testing"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"

	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
)

func TestNewQuestionFromRepository(t *testing.T) {
	ID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	createdAt, updatedAt := time.Now(), time.Now()

	repoQuestion := model.Question{
		ID:        10,
		UUID:      ID.String(),
		Text:      "test question",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	question := NewQuestionFromRepository(&repoQuestion)
	assert.Equal(t, repoQuestion.UUID, question.UUID)
	assert.Equal(t, repoQuestion.Text, question.Text)
	assert.Equal(t, repoQuestion.CreatedAt, question.CreatedAt)
	assert.Equal(t, repoQuestion.UpdatedAt, question.UpdatedAt)
	assert.Equal(t, 0, len(question.Answers))

	repoQuestion.Answers = []model.Answer{
		{
			ID:         10,
			QuestionID: repoQuestion.ID,
			Text:       "test answer",
			IsCorrect:  true,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		},
	}

	question = NewQuestionFromRepository(&repoQuestion)
	assert.Equal(t, 1, len(question.Answers))
	assert.Equal(t, repoQuestion.Answers[0].ID, question.Answers[0].ID)
	assert.Equal(t, repoQuestion.Answers[0].Text, question.Answers[0].Text)
	assert.Equal(t, repoQuestion.Answers[0].IsCorrect, question.Answers[0].IsCorrect)
	assert.Equal(t, repoQuestion.Answers[0].CreatedAt, question.Answers[0].CreatedAt)
	assert.Equal(t, repoQuestion.Answers[0].UpdatedAt, question.Answers[0].UpdatedAt)
}
