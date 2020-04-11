package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"Sharykhin/buffstream-questionnaire/domains/stream/repository/model"
)

func TestNewStreamFromRepository(t *testing.T) {
	assert := assert.New(t)

	repoStream := model.Stream{
		ID:        int64(10),
		UUID:      "UUID",
		Title:     "test title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	actual := NewStreamFromRepository(&repoStream)
	assert.NotNil(actual)
	assert.Equal(repoStream.UUID, actual.UUID)
	assert.Equal(repoStream.Title, actual.Title)
	assert.Equal(repoStream.CreatedAt, actual.CreatedAt)
	assert.Equal(repoStream.UpdatedAt, actual.UpdatedAt)
}
