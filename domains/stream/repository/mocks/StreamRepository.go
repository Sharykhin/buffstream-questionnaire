// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "Sharykhin/buffstream-questionnaire/domains/stream/repository/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// StreamRepository is an autogenerated mock type for the StreamRepository type
type StreamRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: cxt
func (_m *StreamRepository) Count(cxt context.Context) (int64, error) {
	ret := _m.Called(cxt)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(cxt)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(cxt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, UUID, title
func (_m *StreamRepository) Create(ctx context.Context, UUID string, title string) (*model.Stream, error) {
	ret := _m.Called(ctx, UUID, title)

	var r0 *model.Stream
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Stream); ok {
		r0 = rf(ctx, UUID, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Stream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, UUID, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, limit, offset
func (_m *StreamRepository) List(ctx context.Context, limit int64, offset int64) ([]model.Stream, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []model.Stream
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []model.Stream); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Stream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}