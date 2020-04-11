package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	applicationModel "Sharykhin/buffstream-questionnaire/domains/stream/application/model"
	repositoryMock "Sharykhin/buffstream-questionnaire/domains/stream/repository/mocks"
	repositoryModel "Sharykhin/buffstream-questionnaire/domains/stream/repository/model"
	uuidMocks "Sharykhin/buffstream-questionnaire/uuid/mocks"
)

func TestStreamHandler_List(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		name             string
		inLimit          int64
		inOffset         int64
		expectedModels   []repositoryModel.Stream
		expectedListErr  error
		expectedCount    int64
		expectedCountErr error
		assertFn         func(actual []applicationModel.Stream, total int64, err error)
	}{
		{
			name:     "Success",
			inLimit:  int64(10),
			inOffset: int64(10),
			expectedModels: []repositoryModel.Stream{
				{
					ID:        int64(10),
					UUID:      "UUID",
					Title:     "test title",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			expectedListErr:  nil,
			expectedCount:    int64(10),
			expectedCountErr: nil,
			assertFn: func(actual []applicationModel.Stream, total int64, err error) {
				assert.Equal(1, len(actual))
				assert.Equal(int64(10), total)
				assert.Nil(err)
			},
		},
		{
			name:             "Error List",
			inLimit:          int64(10),
			inOffset:         int64(10),
			expectedModels:   nil,
			expectedListErr:  errors.New("something went wring"),
			expectedCount:    int64(10),
			expectedCountErr: nil,
			assertFn: func(actual []applicationModel.Stream, total int64, err error) {
				assert.Equal(0, len(actual))
				assert.Equal(int64(0), total)
				assert.NotNil(err)
			},
		},
		{
			name:     "Error Count",
			inLimit:  int64(10),
			inOffset: int64(10),
			expectedModels: []repositoryModel.Stream{
				{
					ID:        int64(10),
					UUID:      "UUID",
					Title:     "test title",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			expectedListErr:  nil,
			expectedCount:    int64(0),
			expectedCountErr: errors.New("something went wrong"),
			assertFn: func(actual []applicationModel.Stream, total int64, err error) {
				assert.Equal(0, len(actual))
				assert.Equal(int64(0), total)
				assert.NotNil(err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := &repositoryMock.StreamRepository{}
			uuidGeneratorMock := &uuidMocks.Generator{}
			ctx := context.Background()

			repoMock.
				On("List", ctx, tc.inLimit, tc.inOffset).
				Once().
				Return(tc.expectedModels, tc.expectedListErr)

			if tc.expectedListErr == nil {
				repoMock.
					On("Count", ctx).
					Once().
					Return(tc.expectedCount, tc.expectedCountErr)
			} else {
				repoMock.AssertNotCalled(t, "Count", mock.Anything)
			}

			srv := NewStreamService(uuidGeneratorMock, repoMock)
			actualModels, actualTotal, err := srv.List(ctx, tc.inLimit, tc.inOffset)
			repoMock.AssertExpectations(t)
			tc.assertFn(actualModels, actualTotal, err)
		})
	}
}

func TestStreamHandler_Create(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		name          string
		inUUID        string
		inTitle       string
		expectedModel *repositoryModel.Stream
		expectedErr   error
		assertFn      func(actual *applicationModel.Stream, err error)
	}{
		{
			name:    "Success",
			inUUID:  "UUID",
			inTitle: "test stream",
			expectedModel: &repositoryModel.Stream{
				ID:        int64(1),
				UUID:      "UUID",
				Title:     "test stream",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedErr: nil,
			assertFn: func(actual *applicationModel.Stream, err error) {
				assert.NotNil(actual)
				assert.Nil(err)
			},
		},
		{
			name:          "Error",
			inUUID:        "UUID",
			inTitle:       "test stream",
			expectedModel: nil,
			expectedErr:   errors.New("something went wring"),
			assertFn: func(actual *applicationModel.Stream, err error) {
				assert.Nil(actual)
				assert.NotNil(err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := &repositoryMock.StreamRepository{}
			uuidGeneratorMock := &uuidMocks.Generator{}
			ctx := context.Background()

			uuidGeneratorMock.
				On("NewV4").
				Once().
				Return(tc.inUUID)

			repoMock.
				On("Create", ctx, tc.inUUID, tc.inTitle).
				Once().
				Return(tc.expectedModel, tc.expectedErr)

			srv := NewStreamService(uuidGeneratorMock, repoMock)
			actual, err := srv.Create(ctx, tc.inTitle)
			repoMock.AssertExpectations(t)
			uuidGeneratorMock.AssertExpectations(t)

			tc.assertFn(actual, err)
		})
	}
}
