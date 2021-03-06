// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"
	models "famous-quote/models"

	mock "github.com/stretchr/testify/mock"
)

// QuotesRepository is an autogenerated mock type for the QuotesRepository type
type QuotesRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, quotes
func (_m *QuotesRepository) Create(ctx context.Context, quotes *models.Quotes) error {
	ret := _m.Called(ctx, quotes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Quotes) error); ok {
		r0 = rf(ctx, quotes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *QuotesRepository) Get(ctx context.Context, id int64) (*models.Quotes, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Quotes
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Quotes); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quotes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRandom provides a mock function with given fields: ctx
func (_m *QuotesRepository) GetRandom(ctx context.Context) (*models.Quotes, error) {
	ret := _m.Called(ctx)

	var r0 *models.Quotes
	if rf, ok := ret.Get(0).(func(context.Context) *models.Quotes); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quotes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Like provides a mock function with given fields: ctx, id, incr
func (_m *QuotesRepository) Like(ctx context.Context, id int64, incr int64) (int64, error) {
	ret := _m.Called(ctx, id, incr)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) int64); ok {
		r0 = rf(ctx, id, incr)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, id, incr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
