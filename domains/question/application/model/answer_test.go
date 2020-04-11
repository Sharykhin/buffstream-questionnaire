package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
)

func TestNewAnswerFromRepository(t *testing.T) {
	assert := assert.New(t)

	repoAnswer := model.Answer{
		ID:        10,
		Text:      "test answer",
		IsCorrect: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	actual := NewAnswerFromRepository(&repoAnswer)

	assert.NotNil(actual)
	assert.Equal(repoAnswer.ID, actual.ID)
	assert.Equal(repoAnswer.Text, actual.Text)
	assert.Equal(repoAnswer.IsCorrect, actual.IsCorrect)
	assert.Equal(repoAnswer.CreatedAt, actual.CreatedAt)
	assert.Equal(repoAnswer.UpdatedAt, actual.UpdatedAt)
}
