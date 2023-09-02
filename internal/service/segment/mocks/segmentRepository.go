// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	modelsegment "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
)

// segmentRepository is an autogenerated mock type for the segmentRepository type
type segmentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, slug
func (_m *segmentRepository) Create(ctx context.Context, slug string) error {
	ret := _m.Called(ctx, slug)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBySlug provides a mock function with given fields: ctx, slug
func (_m *segmentRepository) DeleteBySlug(ctx context.Context, slug string) (int64, error) {
	ret := _m.Called(ctx, slug)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySlug provides a mock function with given fields: ctx, slug
func (_m *segmentRepository) GetBySlug(ctx context.Context, slug string) (*modelsegment.Segment, error) {
	ret := _m.Called(ctx, slug)

	var r0 *modelsegment.Segment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*modelsegment.Segment, error)); ok {
		return rf(ctx, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *modelsegment.Segment); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelsegment.Segment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newSegmentRepository creates a new instance of segmentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newSegmentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *segmentRepository {
	mock := &segmentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
