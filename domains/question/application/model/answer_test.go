package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
)

func TestNewAnswerFromRepository(t *testing.T) {
	createdAt, updatedAt := time.Now(), time.Now()

	repoAnswer := model.Answer{
		ID:        10,
		Text:      "test answer",
		IsCorrect: true,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	actual := NewAnswerFromRepository(&repoAnswer)
	assert.Equal(t, repoAnswer.ID, actual.ID)
	assert.Equal(t, repoAnswer.Text, actual.Text)
	assert.Equal(t, repoAnswer.IsCorrect, actual.IsCorrect)
	assert.Equal(t, repoAnswer.CreatedAt, actual.CreatedAt)
	assert.Equal(t, repoAnswer.UpdatedAt, actual.UpdatedAt)
}
