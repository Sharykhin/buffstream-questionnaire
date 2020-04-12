package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	applicationModel "Sharykhin/buffstream-questionnaire/domains/question/application/model"
	"Sharykhin/buffstream-questionnaire/domains/question/repository/mocks"
	repositoryModel "Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	appErrors "Sharykhin/buffstream-questionnaire/errors"
)

func TestQuestionHandler_GetOneByID(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		name          string
		inUUID        string
		expectedModel *repositoryModel.Question
		expectedErr   error
		assertFn      func(actual *applicationModel.Question, err error)
	}{
		{
			name:   "Success",
			inUUID: "UUID",
			expectedModel: &repositoryModel.Question{
				ID:        10,
				UUID:      "UUID",
				Text:      "test question",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedErr: nil,
			assertFn: func(actual *applicationModel.Question, err error) {
				assert.Nil(err)
				assert.Equal("UUID", actual.UUID)
			},
		},
		{
			name:          "Error Not Found",
			inUUID:        "UUID",
			expectedModel: nil,
			expectedErr:   fmt.Errorf("no records was found: %w", appErrors.ResourceNotFound),
			assertFn: func(actual *applicationModel.Question, err error) {
				assert.Nil(actual)
				assert.True(errors.Is(err, appErrors.ResourceNotFound))
			},
		},
		{
			name:          "Error Something went wrong",
			expectedModel: nil,
			expectedErr:   errors.New("server error"),
			assertFn: func(actual *applicationModel.Question, err error) {
				assert.Nil(actual)
				assert.False(errors.Is(err, appErrors.ResourceNotFound))
				assert.NotNil(err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := &mocks.QuestionRepository{}
			ctx := context.Background()
			repoMock.
				On("FindOneByIDWithAnswers", ctx, tc.inUUID).
				Once().
				Return(tc.expectedModel, tc.expectedErr)

			srv := NewQuestionService(repoMock)
			actual, err := srv.GetOneByID(ctx, tc.inUUID)
			repoMock.AssertExpectations(t)
			tc.assertFn(actual, err)
		})
	}
}

func TestQuestionHandler_GetAllByStreamID(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		name           string
		inUUID         string
		expectedModels []repositoryModel.Question
		expectedErr    error
		assertFn       func(actual []applicationModel.Question, err error)
	}{
		{
			name:   "Success",
			inUUID: "UUID",
			expectedModels: []repositoryModel.Question{{
				ID:        10,
				UUID:      "UUID",
				Text:      "test question",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}},
			expectedErr: nil,
			assertFn: func(actual []applicationModel.Question, err error) {
				assert.Nil(err)
				assert.Equal(1, len(actual))
				assert.Equal("UUID", actual[0].UUID)
			},
		},
		{
			name:           "Error Something went wrong",
			inUUID:         "UUID",
			expectedModels: nil,
			expectedErr:    errors.New("server error"),
			assertFn: func(actual []applicationModel.Question, err error) {
				assert.Nil(actual)
				assert.NotNil(err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := &mocks.QuestionRepository{}
			ctx := context.Background()
			repoMock.
				On("FindListByStreamID", ctx, tc.inUUID).
				Once().
				Return(tc.expectedModels, tc.expectedErr)

			srv := NewQuestionService(repoMock)
			actual, err := srv.GetAllByStreamID(ctx, tc.inUUID)
			repoMock.AssertExpectations(t)
			tc.assertFn(actual, err)
		})
	}
}

func TestQuestionHandler_GetAllByStreamIDs(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		name           string
		expectedModels []repositoryModel.Stream
		expectedErr    error
		assertFn       func(actual applicationModel.Streams, err error)
	}{
		{
			name: "Success",
			expectedModels: []repositoryModel.Stream{{
				UUID: "UUID",
				Questions: []repositoryModel.Question{
					{
						ID:        10,
						UUID:      "UUID2",
						Text:      "test question",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			}},
			expectedErr: nil,
			assertFn: func(actual applicationModel.Streams, err error) {
				questions, ok := actual["UUID"]
				assert.True(ok)
				assert.Nil(err)
				assert.Equal(1, len(questions))
			},
		},
		{
			name:           "Error Something Went Wrong",
			expectedModels: nil,
			expectedErr:    errors.New("server error"),
			assertFn: func(actual applicationModel.Streams, err error) {
				assert.Nil(actual)
				assert.NotNil(err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := &mocks.QuestionRepository{}
			ctx := context.Background()
			repoMock.
				On("FindListByStreamIDs", ctx, []string{"UUID"}).
				Once().
				Return(tc.expectedModels, tc.expectedErr)

			srv := NewQuestionService(repoMock)
			actual, err := srv.GetAllByStreamIDs(ctx, []string{"UUID"})
			repoMock.AssertExpectations(t)
			tc.assertFn(actual, err)
		})
	}
}
