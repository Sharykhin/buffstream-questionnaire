package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"

	applicationModel "Sharykhin/buffstream-questionnaire/domains/question/application/model"
	"Sharykhin/buffstream-questionnaire/domains/question/repository/mocks"
	repositoryModel "Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	appErrors "Sharykhin/buffstream-questionnaire/errors"
)

func TestQuestionHandler_GetOneByID(t *testing.T) {
	assert := assert.New(t)

	UUID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	tt := []struct {
		name          string
		expectedModel *repositoryModel.Question
		expectedErr   error
		assertFn      func(actual *applicationModel.Question, err error)
	}{
		{
			name: "Success",
			expectedModel: &repositoryModel.Question{
				ID:        10,
				UUID:      UUID.String(),
				Text:      "test question",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedErr: nil,
			assertFn: func(actual *applicationModel.Question, err error) {
				assert.Nil(err)
				assert.Equal(UUID.String(), actual.UUID)
			},
		},
		{
			name:          "Error Not Found",
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
				On("FindOneByIDWithAnswers", ctx, UUID.String()).
				Once().
				Return(tc.expectedModel, tc.expectedErr)

			srv := NewQuestionService(repoMock)
			actual, err := srv.GetOneByID(ctx, UUID.String())
			repoMock.AssertExpectations(t)
			tc.assertFn(actual, err)
		})
	}
}

func TestQuestionHandler_GetAllByStreamID(t *testing.T) {
	assert := assert.New(t)

	UUID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	tt := []struct {
		name           string
		expectedModels []repositoryModel.Question
		expectedErr    error
		assertFn       func(actual []applicationModel.Question, err error)
	}{
		{
			name: "Success",
			expectedModels: []repositoryModel.Question{{
				ID:        10,
				UUID:      UUID.String(),
				Text:      "test question",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}},
			expectedErr: nil,
			assertFn: func(actual []applicationModel.Question, err error) {
				assert.Nil(err)
				assert.Equal(1, len(actual))
				assert.Equal(UUID.String(), actual[0].UUID)
			},
		},
		{
			name:           "Error Something went wrong",
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
				On("FindListByStreamID", ctx, UUID.String()).
				Once().
				Return(tc.expectedModels, tc.expectedErr)

			srv := NewQuestionService(repoMock)
			actual, err := srv.GetAllByStreamID(ctx, UUID.String())
			repoMock.AssertExpectations(t)
			tc.assertFn(actual, err)
		})
	}
}

func TestQuestionHandler_GetAllByStreamIDs(t *testing.T) {
	assert := assert.New(t)

	UUID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	repoUUID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	tt := []struct {
		name           string
		expectedModels []repositoryModel.Stream
		expectedErr    error
		assertFn       func(actual applicationModel.Streams, err error)
	}{
		{
			name: "Success",
			expectedModels: []repositoryModel.Stream{{
				UUID: UUID.String(),
				Questions: []repositoryModel.Question{
					{
						ID:        10,
						UUID:      repoUUID.String(),
						Text:      "test question",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			}},
			expectedErr: nil,
			assertFn: func(actual applicationModel.Streams, err error) {
				questions, ok := actual[UUID.String()]
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
				On("FindListByStreamIDs", ctx, []string{UUID.String()}).
				Once().
				Return(tc.expectedModels, tc.expectedErr)

			srv := NewQuestionService(repoMock)
			actual, err := srv.GetAllByStreamIDs(ctx, []string{UUID.String()})
			repoMock.AssertExpectations(t)
			tc.assertFn(actual, err)
		})
	}
}
