package model

import (
	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	"time"
)

type (
	// Answer model represents answer on an application level
	Answer struct {
		ID        int64     `json:"ID"`
		Text      string    `json:"Text"`
		IsCorrect bool      `json:"-"` // not to show correct answer in APi. Can be adjusted according to permissions
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
	}
)

func NewAnswerFromRepository(repoModel model.Answer) *Answer {
	answer := Answer{
		ID:        repoModel.ID,
		Text:      repoModel.Text,
		IsCorrect: repoModel.IsCorrect,
		CreatedAt: repoModel.CreatedAt,
		UpdatedAt: repoModel.UpdatedAt,
	}

	return &answer
}
