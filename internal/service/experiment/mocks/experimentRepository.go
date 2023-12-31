// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	experiment "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
)

// ExperimentRepository is an autogenerated mock type for the experimentRepository type
type ExperimentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, segment
func (_m *ExperimentRepository) Create(ctx context.Context, segment *experiment.Experiment) error {
	ret := _m.Called(ctx, segment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *experiment.Experiment) error); ok {
		r0 = rf(ctx, segment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBatch provides a mock function with given fields: ctx, experiments
func (_m *ExperimentRepository) CreateBatch(ctx context.Context, experiments []experiment.Experiment) error {
	ret := _m.Called(ctx, experiments)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []experiment.Experiment) error); ok {
		r0 = rf(ctx, experiments)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, segment
func (_m *ExperimentRepository) Delete(ctx context.Context, segment *experiment.Experiment) (int64, error) {
	ret := _m.Called(ctx, segment)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *experiment.Experiment) (int64, error)); ok {
		return rf(ctx, segment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *experiment.Experiment) int64); ok {
		r0 = rf(ctx, segment)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *experiment.Experiment) error); ok {
		r1 = rf(ctx, segment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAllExpired provides a mock function with given fields: ctx
func (_m *ExperimentRepository) DeleteAllExpired(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBatch provides a mock function with given fields: ctx, experiments
func (_m *ExperimentRepository) DeleteBatch(ctx context.Context, experiments []experiment.Experiment) (int64, error) {
	ret := _m.Called(ctx, experiments)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []experiment.Experiment) (int64, error)); ok {
		return rf(ctx, experiments)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []experiment.Experiment) int64); ok {
		r0 = rf(ctx, experiments)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []experiment.Experiment) error); ok {
		r1 = rf(ctx, experiments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExperimentRepository creates a new instance of ExperimentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExperimentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExperimentRepository {
	mock := &ExperimentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
